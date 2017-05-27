// this file was generated by gomacro command: import _b "net/smtp"
// DO NOT EDIT! Any change will be lost when the file is re-generated

package imports

import (
	. "reflect"
	"net/smtp"
)

// reflection: allow interpreted code to import "net/smtp"
func init() {
	Packages["net/smtp"] = Package{
	Binds: map[string]Value{
		"CRAMMD5Auth":	ValueOf(smtp.CRAMMD5Auth),
		"Dial":	ValueOf(smtp.Dial),
		"NewClient":	ValueOf(smtp.NewClient),
		"PlainAuth":	ValueOf(smtp.PlainAuth),
		"SendMail":	ValueOf(smtp.SendMail),
	},
	Types: map[string]Type{
		"Auth":	TypeOf((*smtp.Auth)(nil)).Elem(),
		"Client":	TypeOf((*smtp.Client)(nil)).Elem(),
		"ServerInfo":	TypeOf((*smtp.ServerInfo)(nil)).Elem(),
	},
	Proxies: map[string]Type{
		"Auth":	TypeOf((*Auth_net_smtp)(nil)).Elem(),
	},
	Wrappers: map[string][]string{
	} }
}

// --------------- proxy for net/smtp.Auth ---------------
type Auth_net_smtp struct {
	Object	interface{}
	Next_	func(fromServer []byte, more bool) (toServer []byte, err error)
	Start_	func(server *smtp.ServerInfo) (proto string, toServer []byte, err error)
}
func (Proxy *Auth_net_smtp) Next(fromServer []byte, more bool) (toServer []byte, err error) {
	return Proxy.Next_(fromServer, more)
}
func (Proxy *Auth_net_smtp) Start(server *smtp.ServerInfo) (proto string, toServer []byte, err error) {
	return Proxy.Start_(server)
}
