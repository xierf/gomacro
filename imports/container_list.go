// this file was generated by gomacro command: import _b "container/list"
// DO NOT EDIT! Any change will be lost when the file is re-generated

package imports

import (
	. "reflect"
	"container/list"
)

// reflection: allow interpreted code to import "container/list"
func init() {
	Packages["container/list"] = Package{
	Binds: map[string]Value{
		"New":	ValueOf(list.New),
	},
	Types: map[string]Type{
		"Element":	TypeOf((*list.Element)(nil)).Elem(),
		"List":	TypeOf((*list.List)(nil)).Elem(),
	},
	Proxies: map[string]Type{
	},
	Wrappers: map[string][]string{
	} }
}
