package elf

type File struct {
	data []byte
}

func (f *File) IsElf64Exec() bool {
	if f.data[mag0] != 0x7f {
		return false
	}
	if f.data[mag1] != 'E' {
		return false
	}
	if f.data[mag2] != 'L' {
		return false
	}
	if f.data[mag3] != 'F' {
		return false
	}
	if f.data[class] != Class64 {
		return false
	}
	if f.data[version] != ElfVersion {
		return false
	}
	// f.data[e_type] == ET_EXEC
	if !(f.data[0x10] == 0x02 && f.data[0x11] == 0x00) {
		return false
	}
	return true
}

func ToLittleEndianness(data int64, size int) []byte {
	ret := make([]byte, size)
	for i := range ret {
		ret[i] = byte(data)
		data = data >> 8
	}
	return ret
}

func (f *File) ChangeEntrtPoint() error {
	return nil
}
