package filewatch

import (
	"context"
	"io"
	"io/fs"
	"log"
	"os"
	"path/filepath"

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
		out: make(chan string),
	}
	if f.IsDir() {
		res.watchDir(ctx, path)
	} else {
		res.watchFile(ctx, path)
	}
	return res
}

type impl struct {
	out chan string
}

func (i *impl) Get() <-chan string {
	return i.out
}

func (i *impl) getWatchFileChan(ctx context.Context, filepath string) chan string {
	res := make(chan string)
	file, err := tail.TailFile(filepath, tail.Config{
		Follow: true,
		ReOpen: true,
		Location: &tail.SeekInfo{
			Offset: 0,
			Whence: io.SeekEnd,
		},
	})
	logx.Must(err)
	go func() {
		defer func(file *tail.Tail) {
			logx.Must(file.Stop())
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
				i.out <- line
			}
		}
	}()
}

func (i *impl) watchFile(ctx context.Context, path string) {
	logx.Infof("watch file: %s", path)
	i.appendChanToOut(ctx, i.getWatchFileChan(ctx, path))
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
			i.watchFile(ctx, path)
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
					i.watchFile(ctx, event.Name)
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
