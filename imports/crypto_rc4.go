// this file was generated by gomacro command: import _b "crypto/rc4"
// DO NOT EDIT! Any change will be lost when the file is re-generated

package imports

import (
	. "reflect"
	"crypto/rc4"
)

// reflection: allow interpreted code to import "crypto/rc4"
func init() {
	Packages["crypto/rc4"] = Package{
	Binds: map[string]Value{
		"NewCipher":	ValueOf(rc4.NewCipher),
	},
	Types: map[string]Type{
		"Cipher":	TypeOf((*rc4.Cipher)(nil)).Elem(),
		"KeySizeError":	TypeOf((*rc4.KeySizeError)(nil)).Elem(),
	},
	Proxies: map[string]Type{
	},
	Wrappers: map[string][]string{
	} }
}
