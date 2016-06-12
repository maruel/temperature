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
func toRGBUsingTH(kelvin uint16) (r, b, g uint8) {
	temperature := float64(kelvin) / 100.
	if temperature <= 66. {
		r = 255
	} else {
		red := temperature - 60.
		r = floatToUint8(329.698727446 * math.Pow(red, -0.1332047592))
	}

	if temperature <= 66. {
		g = floatToUint8(99.4708025861*math.Log(temperature) - 161.1195681661)
	} else {
		g = floatToUint8(288.1221695283 * math.Pow(temperature-60., -0.0755148492))
	}

	if temperature >= 66. {
		b = 255
	} else {
		if temperature > 19. {
			b = floatToUint8(138.5177312231*math.Log(temperature-10) - 305.0447927307)
		}
	}
	return
}

func ExampleToRGB() {
	fmt.Printf("Kelvin RRGGBB\n")
	for k := uint16(1000); k <= 9000; k += 500 {
		r, g, b := ToRGB(k)
		fmt.Printf("%-4d   %02X%02X%02X\n", k, r, g, b)
	}
	// Output:
	// Kelvin RRGGBB
	// 1000   FF003B
	// 1500   FF006C
	// 2000   FF008C
	// 2500   FF48A3
	// 3000   FF6DB5
	// 3500   FF8BC4
	// 4000   FFA5D1
	// 4500   FFBADC
	// 5000   FFCEE5
	// 5500   FFE0ED
	// 6000   FFF0F4
	// 6500   FFFFFF
	// 7000   F6FFF4
	// 7500   EBFFEE
	// 8000   E2FFE9
	// 8500   DBFFE5
	// 9000   D6FFE2
}

func ExampleToKelvin() {
	k := ToKelvin(0xFF, 0xFF, 0xFF)
	r, g, b := ToRGB(k)
	fmt.Printf("%d\n", k)
	fmt.Printf("0x%02X0x%02X0x%02X\n", r, g, b)
	// Output:
	// 6473
	// 0xFF0xFF0xFA
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
