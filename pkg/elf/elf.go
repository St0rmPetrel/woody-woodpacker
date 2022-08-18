package elf

import (
	"errors"
	"fmt"
)

// Basic ELF types
type (
	HalfWord  uint16
	Word      uint32
	Address32 uint32
	Address64 uint64
	Offset32  uint32
	Offset64  uint64
)

const (
	identN     = 16
	mag0       = 0x0
	mag1       = 0x1
	mag2       = 0x2
	mag3       = 0x3
	class      = 0x4
	data       = 0x5
	version    = 0x6
	osabi      = 0x7
	abiversion = 0x8
	pad        = 0x9

	Class32 = 1
	Class64 = 2

	LittleEndianness = 1
	BigEndianness    = 2

	ElfVersion = 1

	SystemV                    = 0x00
	HPUX                       = 0x01
	NetBSD                     = 0x02
	Linux                      = 0x03
	GNUHurd                    = 0x04
	Solaris                    = 0x06
	AIX                        = 0x07
	IRIX                       = 0x08
	FreeBSD                    = 0x09
	Tru64                      = 0x0a
	NovellModesto              = 0x0b
	OpenBSD                    = 0x0c
	OpenVMS                    = 0x0d
	NonStopKernel              = 0x0e
	AROS                       = 0x0f
	FenixOS                    = 0x10
	NuxiCloudABI               = 0x11
	StratusTechnologiesOpenVOS = 0x12

	sizeElfHeader32 = 0x34
	sizeElfHeader64 = 0x40
)

var (
	MagicValidationError   = errors.New("invalid magic number")
	ClassValidationError   = errors.New("invalid class format")
	DataValidationError    = errors.New("invalid data endianness")
	VersionValidationError = errors.New("unsupported ELF version")
	OSABIValidationError   = errors.New("unsupported operation system ABI")
)

type Elf64 struct {
	data      []byte
	elfHeader Elf64Header
}

func NewElf64(file []byte) *Elf64 {
	if len(file) < sizeElfHeader64 {
		return nil
	}
	f := Elf64{
		data: file,
	}
	//if !f.IsElfMagic() {
	//	return nil
	//}
	if f.GetClass() != Class64 {
		return nil
	}
	return &f
}

func (f *Elf64) ValidateMagic() error {
	if f.data[mag0] != 0x7f {
		return MagicValidationError
	}
	if f.data[mag1] != 'E' {
		return MagicValidationError
	}
	if f.data[mag2] != 'L' {
		return MagicValidationError
	}
	if f.data[mag3] != 'F' {
		return MagicValidationError
	}
	return nil
}

func (f *Elf64) ValidateClass() error {
	switch f.data[class] {
	case Class64:
		return nil
	case Class32:
		return fmt.Errorf("32 bits are not supported: %w", ClassValidationError)
	default:
		return ClassValidationError
	}
}

func (f *Elf64) ValidateData() error {
	switch f.data[data] {
	case LittleEndianness:
		return nil
	case BigEndianness:
		return nil
	default:
		return DataValidationError
	}
}

func (f *Elf64) ValidateVersion() error {
	if f.data[version] != ElfVersion {
		return VersionValidationError
	}
	return nil
}

func (f *Elf64) ValidateOSABI() error {
	if !(f.data[osabi] <= 0x12) {
		return OSABIValidationError
	}
	return nil
}

func (f *Elf64) GetClass() byte {
	return f.data[class]
}

func (f *Elf64) GetData() byte {
	return f.data[data]
}

func (f *Elf64) GetVersion() byte {
	return f.data[version]
}

func (f *Elf64) GetOSABI() byte {
	return f.data[osabi]
}

func (f *Elf64) GetABIVersion() byte {
	return f.data[abiversion]
}

func (f *Elf64) ParseElfHeader() error {
	for i := range f.data {
		switch {
		case i < 0x04:
			f.elfHeader.ident[i] = f.data[i]
		case i == 0x04:
			if err := f.ValidateMagic(); err != nil {
				return fmt.Errorf("parse elf header: %w", err)
			}
			f.elfHeader.ident[i] = f.data[i]
		case i == 0x05:
			if err := f.ValidateClass(); err != nil {
				return fmt.Errorf("parse elf header: %w", err)
			}
			f.elfHeader.ident[i] = f.data[i]
		case i == 0x06:
			if err := f.ValidateData(); err != nil {
				return fmt.Errorf("parse elf header: %w", err)
			}
			f.elfHeader.ident[i] = f.data[i]
		case i == 0x07:
			if err := f.ValidateVersion(); err != nil {
				return fmt.Errorf("parse elf header: %w", err)
			}
			f.elfHeader.ident[i] = f.data[i]
		case i == 0x08:
			if err := f.ValidateOSABI(); err != nil {
				return fmt.Errorf("parse elf header: %w", err)
			}
			f.elfHeader.ident[i] = f.data[i]
		case i < 0x10:
			// padding
			f.elfHeader.ident[i] = f.data[i]
		}
	}
	return nil
}

type Elf64Header struct {
	ident     [identN]byte
	machine   HalfWord
	version   Word
	entry     Address64
	phoff     Offset64
	shoff     Offset64
	flags     Word
	ehsize    HalfWord
	phentsize HalfWord
	phnum     HalfWord
	shentsize HalfWord
	shnum     HalfWord
	shstrndx  HalfWord
}
