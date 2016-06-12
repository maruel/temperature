// Copyright 2016 Marc-Antoine Ruel. All rights reserved.
// Use of this source code is governed under the MIT License that can be found
// in the LICENSE file.

package temperature

import (
	"fmt"
	"math"
	"testing"
)

// http://www.vendian.org/mncharity/dir3/blackbody/
var blackBody = []uint32{
	0xFF3800, //  1000K
	0xFF5300, //  1200K
	0xFF6500, //  1400K
	0xFF7300, //  1600K
	0xFF7E00, //  1800K
	0xFF8912, //  2000K
	0xFF932C, //  2200K
	0xFF9D3F, //  2400K
	0xFFA54F, //  2600K
	0xFFAD5E, //  2800K
	0xFFB46B, //  3000K
	0xFFBB78, //  3200K
	0xFFC184, //  3400K
	0xFFC78F, //  3600K
	0xFFCC99, //  3800K
	0xFFD1A3, //  4000K
	0xFFD5AD, //  4200K
	0xFFD9B6, //  4400K
	0xFFDDBE, //  4600K
	0xFFE1C6, //  4800K
	0xFFE4CE, //  5000K
	0xFFE8D5, //  5200K
	0xFFEBDC, //  5400K
	0xFFEEE3, //  5600K
	0xFFF0E9, //  5800K
	0xFFF3EF, //  6000K
	0xFFF5F5, //  6200K
	0xFFF8FB, //  6400K
	0xFEF9FF, //  6600K
	0xF9F6FF, //  6800K
	0xF5F3FF, //  7000K
	0xF0F1FF, //  7200K
	0xEDEFFF, //  7400K
	0xE9EDFF, //  7600K
	0xE6EBFF, //  7800K
	0xE3E9FF, //  8000K
	0xE0E7FF, //  8200K
	0xDDE6FF, //  8400K
	0xDAE4FF, //  8600K
	0xD8E3FF, //  8800K
	0xD6E1FF, //  9000K
	0xD3E0FF, //  9200K
	0xD1DFFF, //  9400K
	0xCFDDFF, //  9600K
	0xCEDCFF, //  9800K
	0xCCDBFF, // 10000K
	0xCADAFF, // 10200K
	0xC9D9FF, // 10400K
	0xC7D8FF, // 10600K
	0xC6D8FF, // 10800K
	0xC4D7FF, // 11000K
	0xC3D6FF, // 11200K
	0xC2D5FF, // 11400K
	0xC1D4FF, // 11600K
	0xC0D4FF, // 11800K
	0xBFD3FF, // 12000K
	0xBED2FF, // 12200K
	0xBDD2FF, // 12400K
	0xBCD1FF, // 12600K
	0xBBD1FF, // 12800K
	0xBAD0FF, // 13000K
	0xB9D0FF, // 13200K
	0xB8CFFF, // 13400K
	0xB7CFFF, // 13600K
	0xB7CEFF, // 13800K
	0xB6CEFF, // 14000K
	0xB5CDFF, // 14200K
	0xB5CDFF, // 14400K
	0xB4CCFF, // 14600K
	0xB3CCFF, // 14800K
	0xB3CCFF, // 15000K
	0xB2CBFF, // 15200K
	0xB2CBFF, // 15400K
	0xB1CAFF, // 15600K
	0xB1CAFF, // 15800K
	0xB0CAFF, // 16000K
	0xAFC9FF, // 16200K
	0xAFC9FF, // 16400K
	0xAFC9FF, // 16600K
	0xAEC9FF, // 16800K
	0xAEC8FF, // 17000K
	0xADC8FF, // 17200K
	0xADC8FF, // 17400K
	0xACC7FF, // 17600K
	0xACC7FF, // 17800K
	0xACC7FF, // 18000K
	0xABC7FF, // 18200K
	0xABC6FF, // 18400K
	0xAAC6FF, // 18600K
	0xAAC6FF, // 18800K
	0xAAC6FF, // 19000K
	0xA9C6FF, // 19200K
	0xA9C5FF, // 19400K
	0xA9C5FF, // 19600K
	0xA9C5FF, // 19800K
	0xA8C5FF, // 20000K
	0xA8C5FF, // 20200K
	0xA8C4FF, // 20400K
	0xA7C4FF, // 20600K
	0xA7C4FF, // 20800K
	0xA7C4FF, // 21000K
	0xA7C4FF, // 21200K
	0xA6C3FF, // 21400K
	0xA6C3FF, // 21600K
	0xA6C3FF, // 21800K
	0xA6C3FF, // 22000K
	0xA5C3FF, // 22200K
	0xA5C3FF, // 22400K
	0xA5C3FF, // 22600K
	0xA5C2FF, // 22800K
	0xA4C2FF, // 23000K
	0xA4C2FF, // 23200K
	0xA4C2FF, // 23400K
	0xA4C2FF, // 23600K
	0xA4C2FF, // 23800K
	0xA3C2FF, // 24000K
	0xA3C1FF, // 24200K
	0xA3C1FF, // 24400K
	0xA3C1FF, // 24600K
	0xA3C1FF, // 24800K
	0xA3C1FF, // 25000K
	0xA2C1FF, // 25200K
	0xA2C1FF, // 25400K
	0xA2C1FF, // 25600K
	0xA2C1FF, // 25800K
	0xA2C0FF, // 26000K
	0xA2C0FF, // 26200K
	0xA1C0FF, // 26400K
	0xA1C0FF, // 26600K
	0xA1C0FF, // 26800K
	0xA1C0FF, // 27000K
	0xA1C0FF, // 27200K
	0xA1C0FF, // 27400K
	0xA1C0FF, // 27600K
	0xA0C0FF, // 27800K
	0xA0BFFF, // 28000K
	0xA0BFFF, // 28200K
	0xA0BFFF, // 28400K
	0xA0BFFF, // 28600K
	0xA0BFFF, // 28800K
	0xA0BFFF, // 29000K
	0xA0BFFF, // 29200K
	0x9FBFFF, // 29400K
	0x9FBFFF, // 29600K
	0x9FBFFF, // 29800K
}

// toRGBUsingTH implements Tanner Helland's original algorithm ported to Go.
//
// http://www.tannerhelland.com/4435/convert-temperature-rgb-algorithm-code/
//
// It is not exported because it is both 2x slower than ToRGB and it is less
// accurate. It is kept here to implement the comparison benchmark with Neil
// Bartlett's version.
func toRGBUsingTH(kelvin uint16) (r, g, b uint8) {
	if kelvin == 6500 {
		return 255, 255, 255
	}
	temperature := float64(kelvin) * 0.01
	if kelvin < 6500 {
		r = 255
		g = floatToUint8(99.4708025861*math.Log(temperature) - 161.1195681661)
		if kelvin > 1900 {
			b = floatToUint8(138.5177312231*math.Log(temperature-10) - 305.0447927307)
		}
		return
	}
	red := temperature - 60.
	r = floatToUint8(329.698727446 * math.Pow(red, -0.1332047592))
	g = floatToUint8(288.1221695283 * math.Pow(temperature-60., -0.0755148492))
	b = 255
	return
}

func ExampleToRGB() {
	fmt.Printf("Kelvin RRGGBB\n")
	for k := uint16(1000); k <= 9000; k += 500 {
		if k == 6500 {
			for j := uint16(6499); j <= 6501; j++ {
				r, g, b := ToRGB(j)
				fmt.Printf("%-4d   %02X%02X%02X\n", j, r, g, b)
			}
		} else {
			r, g, b := ToRGB(k)
			fmt.Printf("%-4d   %02X%02X%02X\n", k, r, g, b)
		}
	}
	// Output:
	// Kelvin RRGGBB
	// 1000   FF3B00
	// 1500   FF6C00
	// 2000   FF8C00
	// 2500   FFA348
	// 3000   FFB56D
	// 3500   FFC48B
	// 4000   FFD1A5
	// 4500   FFDCBA
	// 5000   FFE5CE
	// 5500   FFEDE0
	// 6000   FFF4F0
	// 6499   FFFBFF
	// 6500   FFFFFF
	// 6501   FFFCFF
	// 7000   F6F4FF
	// 7500   EBEEFF
	// 8000   E2E9FF
	// 8500   DBE5FF
	// 9000   D6E2FF
}

func ExampleToRGBFast() {
	fmt.Printf("Kelvin RRGGBB\n")
	for k := uint16(1000); k <= 9000; k += 500 {
		if k == 6500 {
			for j := uint16(6499); j <= 6501; j++ {
				r, g, b := toRGBFast(j)
				fmt.Printf("%-4d   %02X%02X%02X\n", j, r, g, b)
			}
		} else {
			r, g, b := toRGBFast(k)
			fmt.Printf("%-4d   %02X%02X%02X\n", k, r, g, b)
		}
	}
	// Output:
	// Kelvin RRGGBB
	// 1000   FF5300
	// 1500   FF6C00
	// 2000   FF932C
	// 2500   FFA147
	// 3000   FFBB78
	// 3500   FFC489
	// 4000   FFD5AD
	// 4500   FFDBBA
	// 5000   FFE8D5
	// 5500   FFECDF
	// 6000   FFF5F5
	// 6499   FFF8FD
	// 6500   FFFFFF
	// 6501   FEF8FF
	// 7000   F0F1FF
	// 7500   EAEDFF
	// 8000   E0E7FF
	// 8500   DBE4FF
	// 9000   D3E0FF
}

func TestKelvin(t *testing.T) {
	k := ToKelvin(0xFF, 0xFF, 0xFF)
	// r, g, b :=
	ToRGB(k)
	// Output is random.
	//fmt.Printf("%d %02X%02X%02X\n", k, r, g, b)
	// // Output:
	// // 7200 F1F1FF
}

func BenchmarkToRGBUsingTH1500(b *testing.B) {
	for i := 0; i < b.N; i++ {
		toRGBUsingTH(1500)
	}
}

func BenchmarkToRGBUsingTH2500(b *testing.B) {
	for i := 0; i < b.N; i++ {
		toRGBUsingTH(2500)
	}
}

func BenchmarkToRGBUsingTH6500(b *testing.B) {
	for i := 0; i < b.N; i++ {
		toRGBUsingTH(6500)
	}
}

func BenchmarkToRGBUsingTH7000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		toRGBUsingTH(7000)
	}
}

func BenchmarkToRGB1500(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ToRGB(1500)
	}
}

func BenchmarkToRGB2500(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ToRGB(2500)
	}
}

func BenchmarkToRGB6500(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ToRGB(6500)
	}
}

func BenchmarkToRGB7000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ToRGB(7000)
	}
}

func BenchmarkToRGBFast1500(b *testing.B) {
	for i := 0; i < b.N; i++ {
		toRGBFast(1500)
	}
}

func BenchmarkToRGBFast2500(b *testing.B) {
	for i := 0; i < b.N; i++ {
		toRGBFast(2500)
	}
}

func BenchmarkToRGBFast6500(b *testing.B) {
	for i := 0; i < b.N; i++ {
		toRGBFast(6500)
	}
}

func BenchmarkToRGBFast7000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		toRGBFast(7000)
	}
}

func TestToRGBConformance(t *testing.T) {
	const acceptable = 18
	for k := uint16(1000); k < 30000; k += 200 {
		r1, g1, b1 := ToRGB(k)
		v := blackBody[(k-1000)/200]
		r2 := uint8(v >> 16)
		g2 := uint8(v >> 8)
		b2 := uint8(v)
		if delta(r1, r2) > acceptable || delta(g1, g2) > acceptable || delta(b1, b2) > acceptable {
			t.Fatalf("Delta (%d,%d,%d) vs (%d,%d,%d) too high for %dK", r1, g1, b1, r2, g2, b2, k)
		}
	}
}

func TestToRGBTHConformance(t *testing.T) {
	const acceptable = 13
	for k := uint16(1000); k < 30000; k += 200 {
		r1, g1, b1 := toRGBUsingTH(k)
		v := blackBody[(k-1000)/200]
		r2 := uint8(v >> 16)
		g2 := uint8(v >> 8)
		b2 := uint8(v)
		if delta(r1, r2) > acceptable || delta(g1, g2) > acceptable || delta(b1, b2) > acceptable {
			t.Fatalf("Delta (%d,%d,%d) vs (%d,%d,%d) too high for %dK", r1, g1, b1, r2, g2, b2, k)
		}
	}
}

func TestToRGBFastConformance(t *testing.T) {
	const acceptable = 27
	for k := uint16(1000); k < 30000; k += 200 {
		r1, g1, b1 := toRGBFast(k)
		v := blackBody[(k-1000)/200]
		r2 := uint8(v >> 16)
		g2 := uint8(v >> 8)
		b2 := uint8(v)
		if delta(r1, r2) > acceptable || delta(g1, g2) > acceptable || delta(b1, b2) > acceptable {
			t.Fatalf("Delta (%d,%d,%d) vs (%d,%d,%d) too high for %dK", r1, g1, b1, r2, g2, b2, k)
		}
	}
}

func delta(a, b uint8) uint8 {
	if a > b {
		return a - b
	}
	return b - a
}
