package ay8912

// #cgo CFLAGS: -g -Wall -I ../include
// #include <stdlib.h>
// #include "../include/ayemu.h"
import "C"
import (
	// "fmt"
	"unsafe"
	// "encoding/hex"

	// "github.com/hajimehoshi/oto"
)

type AY8912 interface {
	WriteReg(r byte, v byte)
	ReadReg(r byte) byte
	EndFrame(audioBufSize int) []byte
	SetSoundFormat(freq,chans,bits int)
	SetFrequency(freq int)
}

type ay8912 struct {
	regs []byte
	ptr  *C.ayemu_ay_t
}

func NewAY8912() AY8912 {
	ay := &ay8912{
		regs: make([]byte, 16),
	}

	ay.ptr = (*C.ayemu_ay_t)(C.malloc(C.sizeof_ayemu_ay_t))
	C.ayemu_init(ay.ptr)
	return ay
}

func (ay *ay8912) SetSoundFormat(freq,chans,bits int) {
	C.ayemu_set_sound_format(ay.ptr,C.int(freq),C.int(chans),C.int(bits>>3))
}

func (ay *ay8912) SetFrequency(freq int) {
	C.ayemu_set_chip_freq(ay.ptr,C.int(freq))
}

func (ay *ay8912) EndFrame(audioBufSize int) []byte {
	audioBuf := C.malloc(C.sizeof_char * C.ulong(audioBufSize))
	defer C.free(unsafe.Pointer(audioBuf))

	// fmt.Print(hex.Dump(ay.regs))
	C.ayemu_set_regs(ay.ptr, (*C.uchar)(&ay.regs[0]))
	C.ayemu_gen_sound(ay.ptr, audioBuf,  C.ulong(audioBufSize))

	data := C.GoBytes(audioBuf, C.int(audioBufSize))
	// fmt.Println("[EndFrame] ay.regs",hex.Dump(ay.regs))
	// fmt.Println("[EndFrame] audioBuf",hex.Dump(C.GoBytes(audioBuf,0x20)))

	return data
}

func (ay *ay8912) WriteReg(r byte, v byte) {
	// fmt.Printf("r:%d v:%d\n", r, v)
	ay.regs[r] = v
}

func (ay *ay8912) ReadReg(r byte) byte {
	return ay.regs[r]
}
