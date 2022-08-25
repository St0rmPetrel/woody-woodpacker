package elf

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToLittleEndianness(t *testing.T) {
	tests := []struct {
		name       string
		data       int64
		size       int
		wantResult []byte
	}{
		{
			name:       "2",
			data:       0x0002,
			size:       2,
			wantResult: []byte{0x02, 0x00},
		},
		{
			name:       "2 size 3",
			data:       0x0002,
			size:       3,
			wantResult: []byte{0x02, 0x00, 0x00},
		},
		{
			name:       "int",
			data:       0xFA_02_3B_33,
			size:       4,
			wantResult: []byte{0x33, 0x3B, 0x02, 0xFA},
		},
	}
	for _, test := range tests {
		got := ToLittleEndianness(test.data, test.size)
		want := test.wantResult
		assert.Equal(t, want, got, test.name)
	}
}
