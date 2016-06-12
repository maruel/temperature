// Copyright 2016 Marc-Antoine Ruel. All rights reserved.
// Use of this source code is governed under the MIT License that can be found
// in the LICENSE file.

package temperature

import (
	"math"
	"testing"
)

// toRGBUsingTH implements Tanner Helland's original algorithm.
//
// It is not exported because it is both 2x slower than ToRGB and it is less
// accurate. It is kept here to implement the comparison benchmark.
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

func BenchmarkToRGBUsingTH(b *testing.B) {
	// Go over 1000K~11000K range to exercise the various code paths.
	for i := 0; i < b.N; i++ {
		for k := uint16(1000); k < 11000; k++ {
			toRGBUsingTH(k)
		}
	}
}

func BenchmarkToRGB(b *testing.B) {
	// Go over 1000K~11000K range to exercise the various code paths.
	for i := 0; i < b.N; i++ {
		for k := uint16(1000); k < 11000; k++ {
			ToRGB(k)
		}
	}
}
