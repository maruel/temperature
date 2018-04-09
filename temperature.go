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
// intensity for a color temperature in Kelvin.
//
// Color Temperature is the color due to black body radiation at a given
// temperature. The temperature is given in Kelvin. The concept is widely used
// in photography and in tools such as f.lux.
//
// The function here converts a given color temperature into a near equivalent
// in the RGB colorspace. The function is based on a curve fit on standard
// sparse set of Kelvin to RGB mappings with a whitepoint (255, 255, 255) at
// 6500K.
//
// NOTE The approximations used are suitable for photo-manipulation and other
// non-critical uses. They are not suitable for medical or other high accuracy
// use cases.
//
// Accuracy is best between 1000K and 30000K.
package temperature

import "math"

// ToRGB returns an RGB representation of the temperature in Kelvin.
func ToRGB(kelvin uint16) (r, g, b uint8) {
	if kelvin == 6500 {
		// Hard fit at 6500K.
		return 255, 255, 255
	}
	temperature := float64(kelvin) * 0.01
	if kelvin < 6500 {
		r = 255
		// a + b x + c Log[x] /.
		// {a -> -155.25485562709179`,
		// b -> -0.44596950469579133`,
		// c -> 104.49216199393888`,
		// x -> (kelvin/100) - 2}
		green := temperature - 2
		g = floatToUint8(-155.25485562709179 - 0.44596950469579133*green + 104.49216199393888*math.Log(green))
		if kelvin > 2000 {
			// a + b x + c Log[x] /.
			// {a -> -254.76935184120902`,
			// b -> 0.8274096064007395`,
			// c -> 115.67994401066147`,
			// x -> kelvin/100 - 10}
			blue := temperature - 10
			b = floatToUint8(-254.76935184120902 + 0.8274096064007395*blue + 115.67994401066147*math.Log(blue))
		}
		return
	}
	b = 255
	// a + b x + c Log[x] /.
	// {a -> 351.97690566805693`,
	// b -> 0.114206453784165`,
	// c -> -40.25366309332127
	//x -> (kelvin/100) - 55}
	red := temperature - 55.
	r = floatToUint8(351.97690566805693 + 0.114206453784165*red - 40.25366309332127*math.Log(red))
	// a + b x + c Log[x] /.
	// {a -> 325.4494125711974`,
	// b -> 0.07943456536662342`,
	// c -> -28.0852963507957`,
	// x -> (kelvin/100) - 50}
	green := temperature - 50.
	g = floatToUint8(325.4494125711974 + 0.07943456536662342*green - 28.0852963507957*math.Log(green))
	return
}

// Leveraging http://www.vendian.org/mncharity/dir3/blackbody/ which is using
// D65.
// TODO(maruel): Convert to 256 steps with a base at 1024K, this would help
// optimize toRGBFast.
const step = 200

const lookUpRedStart = 6400

var lookUpRed = []uint8{
	0xFF, //  6400K
	0xFE, //  6600K
	0xF9, //  6800K
	0xF5, //  7000K
	0xF0, //  7200K
	0xED, //  7400K
	0xE9, //  7600K
	0xE6, //  7800K
	0xE3, //  8000K
	0xE0, //  8200K
	0xDD, //  8400K
	0xDA, //  8600K
	0xD8, //  8800K
	0xD6, //  9000K
	0xD3, //  9200K
	0xD1, //  9400K
	0xCF, //  9600K
	0xCE, //  9800K
	0xCC, // 10000K
	0xCA, // 10200K
	0xC9, // 10400K
	0xC7, // 10600K
	0xC6, // 10800K
	0xC4, // 11000K
	0xC3, // 11200K
	0xC2, // 11400K
	0xC1, // 11600K
	0xC0, // 11800K
	0xBF, // 12000K
	0xBE, // 12200K
	0xBD, // 12400K
	0xBC, // 12600K
	0xBB, // 12800K
	0xBA, // 13000K
	0xB9, // 13200K
	0xB8, // 13400K
	0xB7, // 13600K
	0xB7, // 13800K
	0xB6, // 14000K
	0xB5, // 14200K
	0xB5, // 14400K
	0xB4, // 14600K
	0xB3, // 14800K
	0xB3, // 15000K
	0xB2, // 15200K
	0xB2, // 15400K
	0xB1, // 15600K
	0xB1, // 15800K
	0xB0, // 16000K
	0xAF, // 16200K
	0xAF, // 16400K
	0xAF, // 16600K
	0xAE, // 16800K
	0xAE, // 17000K
	0xAD, // 17200K
	0xAD, // 17400K
	0xAC, // 17600K
	0xAC, // 17800K
	0xAC, // 18000K
	0xAB, // 18200K
	0xAB, // 18400K
	0xAA, // 18600K
	0xAA, // 18800K
	0xAA, // 19000K
	0xA9, // 19200K
	0xA9, // 19400K
	0xA9, // 19600K
	0xA9, // 19800K
	0xA8, // 20000K
	0xA8, // 20200K
	0xA8, // 20400K
	0xA7, // 20600K
	0xA7, // 20800K
	0xA7, // 21000K
	0xA7, // 21200K
	0xA6, // 21400K
	0xA6, // 21600K
	0xA6, // 21800K
	0xA6, // 22000K
	0xA5, // 22200K
	0xA5, // 22400K
	0xA5, // 22600K
	0xA5, // 22800K
	0xA4, // 23000K
	0xA4, // 23200K
	0xA4, // 23400K
	0xA4, // 23600K
	0xA4, // 23800K
	0xA3, // 24000K
	0xA3, // 24200K
	0xA3, // 24400K
	0xA3, // 24600K
	0xA3, // 24800K
	0xA3, // 25000K
	0xA2, // 25200K
	0xA2, // 25400K
	0xA2, // 25600K
	0xA2, // 25800K
	0xA2, // 26000K
	0xA2, // 26200K
	0xA1, // 26400K
	0xA1, // 26600K
	0xA1, // 26800K
	0xA1, // 27000K
	0xA1, // 27200K
	0xA1, // 27400K
	0xA1, // 27600K
	0xA0, // 27800K
	0xA0, // 28000K
	0xA0, // 28200K
	0xA0, // 28400K
	0xA0, // 28600K
	0xA0, // 28800K
	0xA0, // 29000K
	0xA0, // 29200K
	0x9F, // 29400K
	0x9F, // 29600K
	0x9F, // 29800K
	0x9F, // 30000K
}

const lookUpGreenStart = 1000

var lookUpGreen = []uint8{
	0x38, //  1000K
	0x53, //  1200K
	0x65, //  1400K
	0x73, //  1600K
	0x7E, //  1800K
	0x89, //  2000K
	0x93, //  2200K
	0x9D, //  2400K
	0xA5, //  2600K
	0xAD, //  2800K
	0xB4, //  3000K
	0xBB, //  3200K
	0xC1, //  3400K
	0xC7, //  3600K
	0xCC, //  3800K
	0xD1, //  4000K
	0xD5, //  4200K
	0xD9, //  4400K
	0xDD, //  4600K
	0xE1, //  4800K
	0xE4, //  5000K
	0xE8, //  5200K
	0xEB, //  5400K
	0xEE, //  5600K
	0xF0, //  5800K
	0xF3, //  6000K
	0xF5, //  6200K
	0xF8, //  6400K
	0xF9, //  6600K
	0xF6, //  6800K
	0xF3, //  7000K
	0xF1, //  7200K
	0xEF, //  7400K
	0xED, //  7600K
	0xEB, //  7800K
	0xE9, //  8000K
	0xE7, //  8200K
	0xE6, //  8400K
	0xE4, //  8600K
	0xE3, //  8800K
	0xE1, //  9000K
	0xE0, //  9200K
	0xDF, //  9400K
	0xDD, //  9600K
	0xDC, //  9800K
	0xDB, // 10000K
	0xDA, // 10200K
	0xD9, // 10400K
	0xD8, // 10600K
	0xD8, // 10800K
	0xD7, // 11000K
	0xD6, // 11200K
	0xD5, // 11400K
	0xD4, // 11600K
	0xD4, // 11800K
	0xD3, // 12000K
	0xD2, // 12200K
	0xD2, // 12400K
	0xD1, // 12600K
	0xD1, // 12800K
	0xD0, // 13000K
	0xD0, // 13200K
	0xCF, // 13400K
	0xCF, // 13600K
	0xCE, // 13800K
	0xCE, // 14000K
	0xCD, // 14200K
	0xCD, // 14400K
	0xCC, // 14600K
	0xCC, // 14800K
	0xCC, // 15000K
	0xCB, // 15200K
	0xCB, // 15400K
	0xCA, // 15600K
	0xCA, // 15800K
	0xCA, // 16000K
	0xC9, // 16200K
	0xC9, // 16400K
	0xC9, // 16600K
	0xC9, // 16800K
	0xC8, // 17000K
	0xC8, // 17200K
	0xC8, // 17400K
	0xC7, // 17600K
	0xC7, // 17800K
	0xC7, // 18000K
	0xC7, // 18200K
	0xC6, // 18400K
	0xC6, // 18600K
	0xC6, // 18800K
	0xC6, // 19000K
	0xC6, // 19200K
	0xC5, // 19400K
	0xC5, // 19600K
	0xC5, // 19800K
	0xC5, // 20000K
	0xC5, // 20200K
	0xC4, // 20400K
	0xC4, // 20600K
	0xC4, // 20800K
	0xC4, // 21000K
	0xC4, // 21200K
	0xC3, // 21400K
	0xC3, // 21600K
	0xC3, // 21800K
	0xC3, // 22000K
	0xC3, // 22200K
	0xC3, // 22400K
	0xC3, // 22600K
	0xC2, // 22800K
	0xC2, // 23000K
	0xC2, // 23200K
	0xC2, // 23400K
	0xC2, // 23600K
	0xC2, // 23800K
	0xC2, // 24000K
	0xC1, // 24200K
	0xC1, // 24400K
	0xC1, // 24600K
	0xC1, // 24800K
	0xC1, // 25000K
	0xC1, // 25200K
	0xC1, // 25400K
	0xC1, // 25600K
	0xC1, // 25800K
	0xC0, // 26000K
	0xC0, // 26200K
	0xC0, // 26400K
	0xC0, // 26600K
	0xC0, // 26800K
	0xC0, // 27000K
	0xC0, // 27200K
	0xC0, // 27400K
	0xC0, // 27600K
	0xC0, // 27800K
	0xBF, // 28000K
	0xBF, // 28200K
	0xBF, // 28400K
	0xBF, // 28600K
	0xBF, // 28800K
	0xBF, // 29000K
	0xBF, // 29200K
	0xBF, // 29400K
	0xBF, // 29600K
	0xBF, // 29800K
	0xBF, // 30000K
}

const lookUpBlueStart = 1000

var lookUpBlue = []uint8{
	0x00, //  1000K
	0x00, //  1200K
	0x00, //  1400K
	0x00, //  1600K
	0x00, //  1800K
	0x12, //  2000K
	0x2C, //  2200K
	0x3F, //  2400K
	0x4F, //  2600K
	0x5E, //  2800K
	0x6B, //  3000K
	0x78, //  3200K
	0x84, //  3400K
	0x8F, //  3600K
	0x99, //  3800K
	0xA3, //  4000K
	0xAD, //  4200K
	0xB6, //  4400K
	0xBE, //  4600K
	0xC6, //  4800K
	0xCE, //  5000K
	0xD5, //  5200K
	0xDC, //  5400K
	0xE3, //  5600K
	0xE9, //  5800K
	0xEF, //  6000K
	0xF5, //  6200K
	0xFB, //  6400K
	0xFF, //  6600K
}

// toRGBFast returns an RGB representation of the temperature in Kelvin using
// internal lookup tables, linear interpolation and no floating point
// calculation.
func toRGBFast(kelvin uint16) (r, g, b uint8) {
	if kelvin == 6500 {
		// Hard fit at 6500K.
		return 255, 255, 255
	}
	if kelvin < 1000 {
		kelvin = 1000
	}
	if kelvin >= 30000 {
		kelvin = 29999
	}
	d := kelvin - lookUpGreenStart
	i := d / step
	ratio := uint32((d % step) * 255 / step)
	g = uint8((ratio*uint32(lookUpGreen[i]) + (255-ratio)*uint32(lookUpGreen[i+1])) / 255)
	if kelvin < 6500 {
		r = 255
		d = kelvin - lookUpBlueStart
		i = d / step
		ratio = uint32((d % step) * 255 / step)
		b = uint8((ratio*uint32(lookUpBlue[i]) + (255-ratio)*uint32(lookUpBlue[i+1])) / 255)
		return
	}
	d = kelvin - lookUpRedStart
	i = d / step
	ratio = uint32((d % step) * 255 / step)
	r = uint8((ratio*uint32(lookUpRed[i]) + (255-ratio)*uint32(lookUpRed[i+1])) / 255)
	b = 255
	return
}

// ToKelvin converts a RGB color into to the closest Kelvin color temperature.
func ToKelvin(r, g, b uint8) uint16 {
	const epsilon = 0.4
	temperature := 0.
	minTemperature := 1000.
	maxTemperature := 40000.
	for maxTemperature-minTemperature > epsilon {
		temperature = (maxTemperature + minTemperature) * 0.5
		tR, _, tB := ToRGB(floatToUint16(temperature))
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
