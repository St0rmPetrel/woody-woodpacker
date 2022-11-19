package elfeditor

import (
	"bytes"
	"debug/elf"
	"io/fs"
	"io/ioutil"
)

type File struct {
	data []byte
	elf  *elf.File
}

func NewFileCopy(name string) (*File, error) {
	data, err := ioutil.ReadFile(name)
	if err != nil {
		return nil, err
	}
	elfFile, err := elf.NewFile(bytes.NewReader(data))
	if err != nil {
		return nil, err
	}
	if err = elfFile.Close(); err != nil {
		return nil, err
	}
	return &File{
		data: data,
		elf:  elfFile,
	}, nil
}

func (f *File) SaveAs(name string, perm fs.FileMode) error {
	return ioutil.WriteFile(name, f.data, perm)
}
