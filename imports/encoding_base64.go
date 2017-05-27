// this file was generated by gomacro command: import _b "encoding/base64"
// DO NOT EDIT! Any change will be lost when the file is re-generated

package imports

import (
	. "reflect"
	"encoding/base64"
)

// reflection: allow interpreted code to import "encoding/base64"
func init() {
	Packages["encoding/base64"] = Package{
	Binds: map[string]Value{
		"NewDecoder":	ValueOf(base64.NewDecoder),
		"NewEncoder":	ValueOf(base64.NewEncoder),
		"NewEncoding":	ValueOf(base64.NewEncoding),
		"NoPadding":	ValueOf(base64.NoPadding),
		"RawStdEncoding":	ValueOf(&base64.RawStdEncoding).Elem(),
		"RawURLEncoding":	ValueOf(&base64.RawURLEncoding).Elem(),
		"StdEncoding":	ValueOf(&base64.StdEncoding).Elem(),
		"StdPadding":	ValueOf(base64.StdPadding),
		"URLEncoding":	ValueOf(&base64.URLEncoding).Elem(),
	},
	Types: map[string]Type{
		"CorruptInputError":	TypeOf((*base64.CorruptInputError)(nil)).Elem(),
		"Encoding":	TypeOf((*base64.Encoding)(nil)).Elem(),
	},
	Proxies: map[string]Type{
	},
	Wrappers: map[string][]string{
	} }
}
