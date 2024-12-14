package initialize

import (
	"context"
	"path/filepath"

	"github.com/zkep/mygeektime/internal/global"
	"github.com/zkep/mygeektime/lib/storage"
)

func Storage(_ context.Context) error {
	directory, err := filepath.Abs(global.CONF.Storage.Directory)
	if err != nil {
		return err
	}
	global.CONF.Storage.Directory = directory
	s, err := storage.NewLocalStorage(
		global.CONF.Storage.Host,
		global.CONF.Storage.Bucket,
		global.CONF.Storage.Directory,
	)
	if err != nil {
		return err
	}
	global.Storage = s
	return nil
}
