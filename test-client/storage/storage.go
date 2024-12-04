package storage

import (
	"errors"
	"flag"
	"time"

	"github.com/boltdb/bolt"
	"github.com/zeromicro/go-zero/core/logx"
)

var storageFile = flag.String("s", "/etc/guardeye-agent/client.storage", "the storage file")

type IF interface {
	FetchOrSet(key string, f func() []byte) ([]byte, error)
	Close() error
}

var (
	MainBucket  = []byte("main")
	ErrNotFound = errors.New("not found")
)

func New() IF {
	db, err := bolt.Open(*storageFile, 0600, &bolt.Options{Timeout: 1 * time.Second})
	logx.Must(err)
	err = db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(MainBucket)
		if b != nil {
			return nil
		}
		b, err = tx.CreateBucket(MainBucket)
		if err != nil {
			return err
		}
		return nil
	})
	return &impl{
		db: db,
	}
}

type impl struct {
	db *bolt.DB
}

func (i *impl) Close() error {
	return i.db.Close()
}

func (i *impl) FetchOrSet(key string, f func() []byte) ([]byte, error) {
	res, err := i.fetch([]byte(key))
	if err == nil {
		return res, nil
	}
	v := f()
	err = i.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(MainBucket)
		return b.Put([]byte(key), v)
	})
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (i *impl) fetch(key []byte) (res []byte, err error) {
	err = i.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(MainBucket)
		res = b.Get(key)
		if res == nil {
			return ErrNotFound
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}
