// this file was generated by gomacro command: import _b "image/png"
// DO NOT EDIT! Any change will be lost when the file is re-generated

package imports

import (
	. "reflect"
	"image/png"
)

// reflection: allow interpreted code to import "image/png"
func init() {
	Packages["image/png"] = Package{
	Binds: map[string]Value{
		"BestCompression":	ValueOf(png.BestCompression),
		"BestSpeed":	ValueOf(png.BestSpeed),
		"Decode":	ValueOf(png.Decode),
		"DecodeConfig":	ValueOf(png.DecodeConfig),
		"DefaultCompression":	ValueOf(png.DefaultCompression),
		"Encode":	ValueOf(png.Encode),
		"NoCompression":	ValueOf(png.NoCompression),
	},Types: map[string]Type{
		"CompressionLevel":	TypeOf((*png.CompressionLevel)(nil)).Elem(),
		"Encoder":	TypeOf((*png.Encoder)(nil)).Elem(),
		"FormatError":	TypeOf((*png.FormatError)(nil)).Elem(),
		"UnsupportedError":	TypeOf((*png.UnsupportedError)(nil)).Elem(),
	},
	}
}