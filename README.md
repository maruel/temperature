## temperature

package temperature implements an algorithm to calculate the channel intensity
for a color temperature.

[![GoDoc](https://godoc.org/github.com/maruel/temperature?status.svg)](https://godoc.org/github.com/maruel/temperature)

It is a reimplementation of http://github.com/neilbartlett/color-temperature
in Go. License is MIT.

Color Temperature is the color due to black body radiation at a given
temperature. The temperature is given in Kelvin. The concept is widely used in
photography and in tools such as f.lux.

The function here converts a given color temperature into a near equivalent in
the RGB colorspace. The function is based on a curve fit on standard sparse set
of Kelvin to RGB mappings.

NOTE The approximations used are suitable for photo-mainpulation and other
non-critical uses. They are not suitable for medical or other high accuracy use
cases.

Accuracy is best between 1000K and 40000K.

