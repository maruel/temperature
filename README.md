## temperature

package temperature implements an algorithm to calculate the channel intensity
for a color temperature.

[![GoDoc](https://godoc.org/github.com/maruel/temperature?status.svg)](https://godoc.org/github.com/maruel/temperature)

- It is a modified reimplementation of
  [color-temperature.js](http://github.com/neilbartlett/color-temperature) by
  Neil Bartlett ([blog](http://www.zombieprototypes.com/?p=210)) in Go.
  - It is optimized to reduce the number of conditions and do the conditions in
    integers instead of on floats.
  - Add the 6500 whitepoint, which was missing (!)
- ... which itself is a modified reimplementation of [Tanner Helland's
  implementation in Visual
  Basic](http://www.tannerhelland.com/4435/convert-temperature-rgb-algorithm-code/).
- ... which itself is based on Mitchell Charity’s [raw black body color
  table](http://www.vendian.org/mncharity/dir3/blackbody/UnstableURLs/bbr_color.html).
  - this table uses D65. [sRGB specification uses D65 but assumes D50 external
    environment](https://en.wikipedia.org/wiki/SRGB#Viewing_environment).
- 15 years later, [Mitchell Charity changed his mind and now
  recommends](http://www.vendian.org/mncharity/dir3/starcolor/) to use D58
  instead of D65.
  - He created [a new D58
    table](http://www.vendian.org/mncharity/dir3/starcolor/UnstableURLs/starcolorsD58.html).
  - [Original blog post](http://www.vendian.org/mncharity/dir3/blackbody/).


## Explanation

Color Temperature is the color due to black body radiation at a given
temperature. The temperature is given in Kelvin. The concept is widely used in
photography and in tools such as [f.lux](https://justgetflux.com).

The function here converts a given color temperature into a near equivalent in
the RGB colorspace. The function is based on a curve fit on standard sparse set
of Kelvin to RGB mappings with a whitepoint (255, 255, 255) at 6500K.

NOTE The approximations used are suitable for photo-manipulation and other
non-critical uses. They are not suitable for medical or other high accuracy use
cases.

Accuracy is best between 1000K and 40000K.


# History

This project has been implemented specifically for the use case of APA-102 LEDs,
[driven by a Raspberry Pi](https:/github.com/maruel/dlibox-go),
which are very cold by default. Then he reimplemented it in C++ with integer
only arithmetic to [run on a ESP8266](https://github.com/maruel/dlibox-esp).


## Performance

Benchmark run on a Intel i7-5600U:

    BenchmarkToRGBUsingTH1500-4     30000000                42.8 ns/op
    BenchmarkToRGBUsingTH2500-4     30000000                56.8 ns/op
    BenchmarkToRGBUsingTH6500-4     1000000000               2.35 ns/op
    BenchmarkToRGBUsingTH7000-4     10000000               195 ns/op
    BenchmarkToRGB1500-4            30000000                43.8 ns/op
    BenchmarkToRGB2500-4            30000000                57.5 ns/op
    BenchmarkToRGB6500-4            1000000000               2.34 ns/op
    BenchmarkToRGB7000-4            30000000                57.8 ns/op


## License

License is MIT.
