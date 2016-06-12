// Copyright 2016 Marc-Antoine Ruel. All rights reserved.
// Use of this source code is governed under the MIT License that can be found
// in the LICENSE file.
//
//  Neil Bartlett
//  neilbartlett.com
//  2015-01-22
//
//  Copyright [2015] [Neil Bartlett] *

// package temperature implements an algorithm to calculate the channel
// intensity for a color temperature.
//
// It is a reimplementation of http://github.com/neilbartlett/color-temperature
// in Go. License is MIT.
//
// Color Temperature is the color due to black body radiation at a given
// temperature. The temperature is given in Kelvin. The concept is widely used
// in photography and in tools such as f.lux.
//
// The function here converts a given color temperature into a near equivalent
// in the RGB colorspace. The function is based on a curve fit on standard
// sparse set of Kelvin to RGB mappings.
//
// NOTE The approximations used are suitable for photo-mainpulation and other
// non-critical uses. They are not suitable for medical or other high accuracy
// use cases.
//
// Accuracy is best between 1000K and 40000K.
package temperature

import "math"

// ToRGB returns an RGB representation of the temperature in Kelvin.
func ToRGB(kelvin uint16) (r, b, g uint8) {
	temperature := float64(kelvin) / 100.
	if temperature < 66. {
		r = 255
	} else {
		// a + b x + c Log[x] /.
		// {a -> 351.97690566805693`,
		// b -> 0.114206453784165`,
		// c -> -40.25366309332127
		//x -> (kelvin/100) - 55}
		red := temperature - 55.
		r = floatToUint8(351.97690566805693 + 0.114206453784165*red - 40.25366309332127*math.Log(red))
	}

	// Calculate green
	if temperature < 66. {
		// a + b x + c Log[x] /.
		// {a -> -155.25485562709179`,
		// b -> -0.44596950469579133`,
		// c -> 104.49216199393888`,
		// x -> (kelvin/100) - 2}
		green := temperature - 2
		g = floatToUint8(-155.25485562709179 - 0.44596950469579133*green + 104.49216199393888*math.Log(green))
	} else {
		// a + b x + c Log[x] /.
		// {a -> 325.4494125711974`,
		// b -> 0.07943456536662342`,
		// c -> -28.0852963507957`,
		// x -> (kelvin/100) - 50}
		green := temperature - 50.
		g = floatToUint8(325.4494125711974 + 0.07943456536662342*green - 28.0852963507957*math.Log(green))
	}

	// Calculate blue
	if temperature >= 66. {
		b = 255
	} else {
		if temperature > 20. {
			// a + b x + c Log[x] /.
			// {a -> -254.76935184120902`,
			// b -> 0.8274096064007395`,
			// c -> 115.67994401066147`,
			// x -> kelvin/100 - 10}
			blue := temperature - 10
			b = floatToUint8(-254.76935184120902 + 0.8274096064007395*blue + 115.67994401066147*math.Log(blue))
		}
	}
	return
}

// ToKelvin converts a RGB color into to the closest Kelvin color temperature.
func ToKelvin(r, g, b uint8) uint16 {
	var temperature float64
	const epsilon = 0.4
	minTemperature := 1000.
	maxTemperature := 40000.
	for maxTemperature-minTemperature > epsilon {
		temperature = (maxTemperature + minTemperature) / 2.
		tR, tB, _ := ToRGB(floatToUint16(temperature))
		if float32(tB)/float32(tR) >= float32(b)/float32(r) {
			maxTemperature = temperature
		} else {
			minTemperature = temperature
		}
	}
	return floatToUint16(temperature)
}

func floatToUint8(x float64) uint8 {
	if x >= 254.4 {
		return 255
	}
	if x <= 0. {
		return 0
	}
	return uint8(math.Ceil(x + 0.5))
}

func floatToUint16(x float64) uint16 {
	if x >= 65534.4 {
		return 65535
	}
	if x <= 0. {
		return 0
	}
	return uint16(math.Ceil(x + 0.5))
}
