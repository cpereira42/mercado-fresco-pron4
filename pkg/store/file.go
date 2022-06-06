package store

import (
	"encoding/json"
	"os"
)

const (
	FileType Type = "file"
)

type Store interface {
	Read(data interface{}) error
	Write(data interface{}) error
}

type Type string

type FileStore struct {
	FileName string
}

func New(store Type, fileName string) Store {
	switch store {
	case FileType:
		return &FileStore{fileName}
	}
	return nil
}

func (fs *FileStore) Write(data interface{}) error {
	fileData, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		return err
	}
	return os.WriteFile(fs.FileName, fileData, 0644)
}

func (fs *FileStore) Read(data interface{}) error {
	file, err := os.ReadFile(fs.FileName)
	if err != nil {
		return err
	}
	return json.Unmarshal(file, &data)
}
