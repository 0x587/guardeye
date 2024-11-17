package filewatch

import (
	"context"
	"io"
	"io/fs"
	"log"
	"os"
	"path/filepath"

	"github.com/0x587/guardeye/report/report"
	"github.com/0x587/guardeye/report/reportclient"
	"github.com/0x587/guardeye/test-client/provider"
	"github.com/fsnotify/fsnotify"
	"github.com/hpcloud/tail"
	"github.com/zeromicro/go-zero/core/logx"
)

func New(ctx context.Context, path string) provider.IF {
	f, err := os.Stat(path)
	if os.IsNotExist(err) {
		log.Fatal("path not exist")
	}
	res := &impl{
		path: path,
		out:  make(chan string),
	}
	if f.IsDir() {
		res.watchDir(ctx, path)
	} else {
		res.watchFile(ctx, path, true)
	}
	return res
}

type impl struct {
	path string
	out  chan string
}

func (i *impl) getProvider() reportclient.Provider {
	return reportclient.Provider{
		Type: ProviderType,
		Args: []string{i.path},
	}
}

const (
	ProviderType = "FileWatcher"
)

func (i *impl) Get() <-chan *provider.Msg {
	res := make(chan *provider.Msg)
	go func() {
		for msg := range i.out {
			res <- &provider.Msg{
				Message:  msg,
				Type:     report.LogType_TEXT,
				Provider: i.getProvider(),
			}
		}
	}()
	return res
}

func (i *impl) getWatchFileChan(ctx context.Context, filepath string, alreadyExist bool) chan string {
	res := make(chan string)
	tconf := tail.Config{
		Follow: true,
		ReOpen: true,
	}
	if alreadyExist {
		tconf.Location = &tail.SeekInfo{
			Offset: 0,
			Whence: io.SeekEnd,
		}
	}
	file, err := tail.TailFile(filepath, tconf)
	logx.Must(err)
	go func() {
		defer func(file *tail.Tail) {
			_ = file.Stop()
			close(res)
		}(file)
		for {
			select {
			case <-ctx.Done():
				return
			case line, ok := <-file.Lines:
				if !ok {
					return
				}
				res <- line.Text
			}
		}
	}()
	return res
}

func (i *impl) appendChanToOut(ctx context.Context, c chan string) {
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case line := <-c:
				if line == "" {
					continue
				}
				i.out <- line
			}
		}
	}()
}

func (i *impl) watchFile(ctx context.Context, path string, alreadyExist bool) {
	logx.Infof("watch file: %s", path)
	i.appendChanToOut(ctx, i.getWatchFileChan(ctx, path, alreadyExist))
}

func (i *impl) watchDir(ctx context.Context, dirpath string) {
	watcher, err := fsnotify.NewWatcher()
	logx.Must(err)
	logx.Must(filepath.Walk(dirpath, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			logx.Must(watcher.Add(path))
		} else {
			i.watchFile(ctx, path, true)
		}
		return nil
	}))
	go func() {
		defer func(watcher *fsnotify.Watcher) {
			logx.Must(watcher.Close())
		}(watcher)
		for {
			select {
			case <-ctx.Done():
				return
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				if event.Op&fsnotify.Create == fsnotify.Create {
					stat, err := os.Stat(event.Name)
					if err != nil {
						continue
					}
					if stat.IsDir() {
						logx.Must(watcher.Add(event.Name))
						continue
					}
					i.watchFile(ctx, event.Name, false)
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				logx.Error(err)
			}
		}
	}()
}
