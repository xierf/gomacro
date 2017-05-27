// this file was generated by gomacro command: import _b "image"
// DO NOT EDIT! Any change will be lost when the file is re-generated

package imports

import (
	. "reflect"
	"image"
	"image/color"
)

// reflection: allow interpreted code to import "image"
func init() {
	Packages["image"] = Package{
	Binds: map[string]Value{
		"Black":	ValueOf(&image.Black).Elem(),
		"Decode":	ValueOf(image.Decode),
		"DecodeConfig":	ValueOf(image.DecodeConfig),
		"ErrFormat":	ValueOf(&image.ErrFormat).Elem(),
		"NewAlpha":	ValueOf(image.NewAlpha),
		"NewAlpha16":	ValueOf(image.NewAlpha16),
		"NewCMYK":	ValueOf(image.NewCMYK),
		"NewGray":	ValueOf(image.NewGray),
		"NewGray16":	ValueOf(image.NewGray16),
		"NewNRGBA":	ValueOf(image.NewNRGBA),
		"NewNRGBA64":	ValueOf(image.NewNRGBA64),
		"NewNYCbCrA":	ValueOf(image.NewNYCbCrA),
		"NewPaletted":	ValueOf(image.NewPaletted),
		"NewRGBA":	ValueOf(image.NewRGBA),
		"NewRGBA64":	ValueOf(image.NewRGBA64),
		"NewUniform":	ValueOf(image.NewUniform),
		"NewYCbCr":	ValueOf(image.NewYCbCr),
		"Opaque":	ValueOf(&image.Opaque).Elem(),
		"Pt":	ValueOf(image.Pt),
		"Rect":	ValueOf(image.Rect),
		"RegisterFormat":	ValueOf(image.RegisterFormat),
		"Transparent":	ValueOf(&image.Transparent).Elem(),
		"White":	ValueOf(&image.White).Elem(),
		"YCbCrSubsampleRatio410":	ValueOf(image.YCbCrSubsampleRatio410),
		"YCbCrSubsampleRatio411":	ValueOf(image.YCbCrSubsampleRatio411),
		"YCbCrSubsampleRatio420":	ValueOf(image.YCbCrSubsampleRatio420),
		"YCbCrSubsampleRatio422":	ValueOf(image.YCbCrSubsampleRatio422),
		"YCbCrSubsampleRatio440":	ValueOf(image.YCbCrSubsampleRatio440),
		"YCbCrSubsampleRatio444":	ValueOf(image.YCbCrSubsampleRatio444),
		"ZP":	ValueOf(&image.ZP).Elem(),
		"ZR":	ValueOf(&image.ZR).Elem(),
	},
	Types: map[string]Type{
		"Alpha":	TypeOf((*image.Alpha)(nil)).Elem(),
		"Alpha16":	TypeOf((*image.Alpha16)(nil)).Elem(),
		"CMYK":	TypeOf((*image.CMYK)(nil)).Elem(),
		"Config":	TypeOf((*image.Config)(nil)).Elem(),
		"Gray":	TypeOf((*image.Gray)(nil)).Elem(),
		"Gray16":	TypeOf((*image.Gray16)(nil)).Elem(),
		"Image":	TypeOf((*image.Image)(nil)).Elem(),
		"NRGBA":	TypeOf((*image.NRGBA)(nil)).Elem(),
		"NRGBA64":	TypeOf((*image.NRGBA64)(nil)).Elem(),
		"NYCbCrA":	TypeOf((*image.NYCbCrA)(nil)).Elem(),
		"Paletted":	TypeOf((*image.Paletted)(nil)).Elem(),
		"PalettedImage":	TypeOf((*image.PalettedImage)(nil)).Elem(),
		"Point":	TypeOf((*image.Point)(nil)).Elem(),
		"RGBA":	TypeOf((*image.RGBA)(nil)).Elem(),
		"RGBA64":	TypeOf((*image.RGBA64)(nil)).Elem(),
		"Rectangle":	TypeOf((*image.Rectangle)(nil)).Elem(),
		"Uniform":	TypeOf((*image.Uniform)(nil)).Elem(),
		"YCbCr":	TypeOf((*image.YCbCr)(nil)).Elem(),
		"YCbCrSubsampleRatio":	TypeOf((*image.YCbCrSubsampleRatio)(nil)).Elem(),
	},
	Proxies: map[string]Type{
		"Image":	TypeOf((*Image_image)(nil)).Elem(),
		"PalettedImage":	TypeOf((*PalettedImage_image)(nil)).Elem(),
	},
	Wrappers: map[string][]string{
		"NYCbCrA":	[]string{"Bounds","COffset","YCbCrAt","YOffset",},
	} }
}

// --------------- proxy for image.Image ---------------
type Image_image struct {
	Object	interface{}
	At_	func(x int, y int) color.Color
	Bounds_	func() image.Rectangle
	ColorModel_	func() color.Model
}
func (Proxy *Image_image) At(x int, y int) color.Color {
	return Proxy.At_(x, y)
}
func (Proxy *Image_image) Bounds() image.Rectangle {
	return Proxy.Bounds_()
}
func (Proxy *Image_image) ColorModel() color.Model {
	return Proxy.ColorModel_()
}

// --------------- proxy for image.PalettedImage ---------------
type PalettedImage_image struct {
	Object	interface{}
	At_	func(x int, y int) color.Color
	Bounds_	func() image.Rectangle
	ColorIndexAt_	func(x int, y int) uint8
	ColorModel_	func() color.Model
}
func (Proxy *PalettedImage_image) At(x int, y int) color.Color {
	return Proxy.At_(x, y)
}
func (Proxy *PalettedImage_image) Bounds() image.Rectangle {
	return Proxy.Bounds_()
}
func (Proxy *PalettedImage_image) ColorIndexAt(x int, y int) uint8 {
	return Proxy.ColorIndexAt_(x, y)
}
func (Proxy *PalettedImage_image) ColorModel() color.Model {
	return Proxy.ColorModel_()
}
