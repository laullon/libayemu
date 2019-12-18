package ay8912

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

const (
	MUTE    = 0
	TONE_A  = 1
	TONE_B  = 2
	TONE_C  = 4
	NOISE_A = 8
	NOISE_B = 16
	NOISE_C = 32
)

const (
	sampleRate      = 44100
	channelNum      = 2
	bitDepthInBytes = 2
)

type test struct {
	name                                                                     string
	tonea, toneb, tonec, noise, control, vola, volb, volc, envfreq, envstyle int
}

var testcases = []*test{
	&test{"Mute: tones 400, volumes 15 noise 15", 400, 400, 400, 15 /*Noise*/, MUTE /* Ctrl */, 15, 15, 15 /* env freq,style */, 4000, 4},
	&test{"Mute: tones 400, noise 25, volumes 31 (use env)", 400, 400, 400, 25, MUTE, 31, 31, 31 /* env freq,style */, 4000, 4},

	&test{"Channel A: tone 400, volume 0", 400, 0, 0, 0, TONE_A, 0, 0, 0 /* env freq,style */, 0, 0},
	&test{"Channel A: tone 400, volume 5", 400, 0, 0, 0, TONE_A, 5, 0, 0, 0, 0},
	&test{"Channel A: tone 400, volume 10", 400, 0, 0, 0, TONE_A, 10, 0, 0, 0, 0},
	&test{"Channel A: tone 400, volume 15", 400, 0, 0, 0, TONE_A, 15, 0, 0, 0, 0},

	&test{"Channel B: tone 400, volume 0", 0, 400, 0, 0 /*Noise*/, TONE_B /* Ctrl */, 0, 0, 0 /* env freq,style */, 0, 0},
	&test{"Channel B: tone 400, volume 5", 0, 400, 0, 0 /*Noise*/, TONE_B /* Ctrl */, 0, 5, 0 /* env freq,style */, 0, 0},
	&test{"Channel B: tone 400, volume 10", 0, 400, 0, 0 /*Noise*/, TONE_B /* Ctrl */, 0, 10, 0 /* env freq,style */, 0, 0},
	&test{"Channel B: tone 400, volume 15", 0, 400, 0, 0 /*Noise*/, TONE_B /* Ctrl */, 0, 15, 0 /* env freq,style */, 0, 0},

	&test{"Channel C: tone 400, volume 0", 0, 0, 400, 0 /*Noise*/, TONE_C /* Ctrl */, 0, 0, 0 /* env freq,style */, 0, 0},
	&test{"Channel C: tone 400, volume 5", 0, 0, 400, 0 /*Noise*/, TONE_C /* Ctrl */, 0, 0, 5 /* env freq,style */, 0, 0},
	&test{"Channel C: tone 400, volume 10", 0, 0, 400, 0 /*Noise*/, TONE_C /* Ctrl */, 0, 0, 10 /* env freq,style */, 0, 0},
	&test{"Channel C: tone 400, volume 15", 0, 0, 400, 0 /*Noise*/, TONE_C /* Ctrl */, 0, 0, 15 /* env freq,style */, 0, 0},

	&test{"Channel B: noise period = 0, volume = 15", 0, 3000, 0, 0, NOISE_B, 0, 15, 0, 0, 0},
	&test{"Channel B: noise period = 5, volume = 15", 0, 3000, 0, 5, NOISE_B, 0, 15, 0, 0, 0},
	&test{"Channel B: noise period = 10, volume = 15", 0, 3000, 0, 10, NOISE_B, 0, 15, 0, 0, 0},
	&test{"Channel B: noise period = 15, volume = 15", 0, 3000, 0, 15, NOISE_B, 0, 15, 0, 0, 0},
	&test{"Channel B: noise period = 20, volume = 15", 0, 3000, 0, 20, NOISE_B, 0, 15, 0, 0, 0},
	&test{"Channel B: noise period = 25, volume = 15", 0, 3000, 0, 25, NOISE_B, 0, 15, 0, 0, 0},
	&test{"Channel B: noise period = 31, volume = 15", 0, 3000, 0, 31, NOISE_B, 0, 15, 0, 0, 0},

	&test{"Channel A: tone 400, volume = 15, envelop 0 freq 4000", 400, 0, 0, 0, TONE_A, 15 | 0x10, 0, 0, 4000, 0},
	&test{"Channel A: tone 400, volume = 15, envelop 1 freq 4000", 400, 0, 0, 0, TONE_A, 15 | 0x10, 0, 0, 4000, 1},
	&test{"Channel A: tone 400, volume = 15, envelop 2 freq 4000", 400, 0, 0, 0, TONE_A, 15 | 0x10, 0, 0, 4000, 2},
	&test{"Channel A: tone 400, volume = 15, envelop 3 freq 4000", 400, 0, 0, 0, TONE_A, 15 | 0x10, 0, 0, 4000, 3},
	&test{"Channel A: tone 400, volume = 15, envelop 4 freq 4000", 400, 0, 0, 0, TONE_A, 15 | 0x10, 0, 0, 4000, 4},
	&test{"Channel A: tone 400, volume = 15, envelop 5 freq 4000", 400, 0, 0, 0, TONE_A, 15 | 0x10, 0, 0, 4000, 5},
	&test{"Channel A: tone 400, volume = 15, envelop 6 freq 4000", 400, 0, 0, 0, TONE_A, 15 | 0x10, 0, 0, 4000, 6},
	&test{"Channel A: tone 400, volume = 15, envelop 7 freq 4000", 400, 0, 0, 0, TONE_A, 15 | 0x10, 0, 0, 4000, 7},
	&test{"Channel A: tone 400, volume = 15, envelop 8 freq 4000", 400, 0, 0, 0, TONE_A, 15 | 0x10, 0, 0, 4000, 8},
	&test{"Channel A: tone 400, volume = 15, envelop 9 freq 4000", 400, 0, 0, 0, TONE_A, 15 | 0x10, 0, 0, 4000, 9},
	&test{"Channel A: tone 400, volume = 15, envelop 10 freq 4000", 400, 0, 0, 0, TONE_A, 15 | 0x10, 0, 0, 4000, 10},
	&test{"Channel A: tone 400, volume = 15, envelop 11 freq 4000", 400, 0, 0, 0, TONE_A, 15 | 0x10, 0, 0, 4000, 11},
	&test{"Channel A: tone 400, volume = 15, envelop 12 freq 4000", 400, 0, 0, 0, TONE_A, 15 | 0x10, 0, 0, 4000, 12},
	&test{"Channel A: tone 400, volume = 15, envelop 13 freq 4000", 400, 0, 0, 0, TONE_A, 15 | 0x10, 0, 0, 4000, 13},
	&test{"Channel A: tone 400, volume = 15, envelop 14 freq 4000", 400, 0, 0, 0, TONE_A, 15 | 0x10, 0, 0, 400, 14},
	&test{"Channel A: tone 400, volume = 15, envelop 15 freq 4000", 400, 0, 0, 0, TONE_A, 15 | 0x10, 0, 0, 4000, 15},
}

var audioData []byte

func TestST(t *testing.T) {
	chip := NewAY8912()
	for test := 0; test < len(testcases); test++ {
		fmt.Printf("Test %d: %s\n", test, testcases[test].name)
		gen_sound(chip, testcases[test].tonea, testcases[test].toneb,
			testcases[test].tonec, testcases[test].noise,
			testcases[test].control, testcases[test].vola,
			testcases[test].volb, testcases[test].volc,
			testcases[test].envfreq, testcases[test].envstyle)
	}
	time.Sleep(30 * time.Second)
	assert.FailNow(t, "--")
}

func gen_sound(chip AY8912, tonea, toneb, tonec, noise, control, vola, volb, volc, envfreq, envstyle int) []byte {
	//   int n, len;
	regs := make([]byte, 14)

	/* setup regs */
	regs[0] = byte(tonea & 0xff)
	regs[1] = byte(tonea >> 8)
	regs[2] = byte(toneb & 0xff)
	regs[3] = byte(toneb >> 8)
	regs[4] = byte(tonec & 0xff)
	regs[5] = byte(tonec >> 8)
	regs[6] = byte(noise)
	regs[7] = byte(control) & 0x3f /* invert bits 0-5 */
	regs[8] = byte(vola)           /* included bit 4 */
	regs[9] = byte(volb)
	regs[10] = byte(volc)
	regs[11] = byte(envfreq & 0xff)
	regs[12] = byte(envfreq >> 8)
	regs[13] = byte(envstyle)

	/* test setreg function: set from array and dump internal regs data */
	for i := byte(0); i < 14; i++ {
		chip.WriteReg(i, regs[i])
	}
	// fmt.Printf("\tRegs: A=%d B=%d C=%d N=%02d R7=[%d%d%d%d%d%d] vols: A=%d B=%d C=%d EnvFreq=%d style %d\n",
	// 	ay.regs.tone_a, ay.regs.tone_b, ay.regs.tone_c, ay.regs.noise,
	// 	ay.regs.R7_tone_a, ay.regs.R7_tone_b, ay.regs.R7_tone_c,
	// 	ay.regs.R7_noise_a, ay.regs.R7_noise_b, ay.regs.R7_noise_c,
	// 	ay.regs.vol_a, ay.regs.vol_b, ay.regs.vol_c,
	// 	ay.regs.env_freq, ay.regs.env_style)

	return chip.EndFrame()
}
