package elf

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
	identN = 16
)

type Elf32Header struct {
	ident     [identN]byte
	machine   HalfWord
	version   Word
	entry     Address32
	phoff     Offset32
	shoff     Offset32
	flags     Word
	ehsize    HalfWord
	phentsize HalfWord
	phnum     HalfWord
	shentsize HalfWord
	shnum     HalfWord
	shstrndx  HalfWord
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
