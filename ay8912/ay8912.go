package ay8912

type AY8912 interface {
	PortManager
	SetReg(r byte)
	WriteReg(v byte)
	ReadReg() byte
}

type ay8912 struct {
	regs []byte
	reg  byte
}

func NewAY8912() AY8912 {
	return &ay8912{
		regs: make([]byte, 16),
	}
}

func (ay *ay8912) SetReg(r byte) {
	ay.reg = r
}

func (ay *ay8912) WriteReg(v byte) {
	ay.regs[ay.reg] = v
}

func (ay *ay8912) ReadReg() byte {
	return ay.regs[ay.reg]
}
