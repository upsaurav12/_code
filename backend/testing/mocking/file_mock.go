package main

import "io"

type FileSystem interface {
	Open(filename string) (io.ReadCloser, error)
}

type mockFileSystem struct {
	openResult io.ReadCloser
	openError  error
}

func (m mockFileSystem) Open(filename string) (io.ReadCloser, error) {
	return m.openResult, m.openError
}
