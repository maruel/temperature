// Copyright 2016 Marc-Antoine Ruel. All rights reserved.
// Use of this source code is governed under the MIT License that can be found
// in the LICENSE file.

package temperature

import (
	"fmt"
	"math"
	"testing"
)

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

func ExampleToKelvin() {
	k := ToKelvin(0xFF, 0xFF, 0xFF)
	r, g, b := ToRGB(k)
	fmt.Printf("%d %02X%02X%02X\n", k, r, g, b)
	// Output:
	// 7087 F4F3FF
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
