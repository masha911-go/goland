// **Интерфейс с контекстом**
//
//	Создайте интерфейс `Worker` с методом `Do(ctx context.Context) error`. Реализуйте его для `FileProcessor` и `NetworkFetcher`.
package main

import (
	"context"
	"errors"
)

type Worker interface {
	Do(ctx context.Context) error
}

type FileProcessor struct {
	FilePath string
}

func (fp *FileProcessor) Do(ctx context.Context) error {
	if fp.FilePath == "/path/to/file" {
		return nil
	} else {
		return errors.New("file not found")
	}
}

type NetworkFetcher struct {
	URL string
}

func (nf *NetworkFetcher) Do(ctx context.Context) error {
	if nf.URL == "https/ok" {
		return nil
	} else {
		return errors.New("not ok")
	}
}
