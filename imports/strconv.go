// this file was generated by gomacro command: import _b "strconv"
// DO NOT EDIT! Any change will be lost when the file is re-generated

package imports

import (
	. "reflect"
	"strconv"
)

// reflection: allow interpreted code to import "strconv"
func init() {
	Packages["strconv"] = Package{
	Binds: map[string]Value{
		"AppendBool":	ValueOf(strconv.AppendBool),
		"AppendFloat":	ValueOf(strconv.AppendFloat),
		"AppendInt":	ValueOf(strconv.AppendInt),
		"AppendQuote":	ValueOf(strconv.AppendQuote),
		"AppendQuoteRune":	ValueOf(strconv.AppendQuoteRune),
		"AppendQuoteRuneToASCII":	ValueOf(strconv.AppendQuoteRuneToASCII),
		"AppendQuoteRuneToGraphic":	ValueOf(strconv.AppendQuoteRuneToGraphic),
		"AppendQuoteToASCII":	ValueOf(strconv.AppendQuoteToASCII),
		"AppendQuoteToGraphic":	ValueOf(strconv.AppendQuoteToGraphic),
		"AppendUint":	ValueOf(strconv.AppendUint),
		"Atoi":	ValueOf(strconv.Atoi),
		"CanBackquote":	ValueOf(strconv.CanBackquote),
		"ErrRange":	ValueOf(&strconv.ErrRange).Elem(),
		"ErrSyntax":	ValueOf(&strconv.ErrSyntax).Elem(),
		"FormatBool":	ValueOf(strconv.FormatBool),
		"FormatFloat":	ValueOf(strconv.FormatFloat),
		"FormatInt":	ValueOf(strconv.FormatInt),
		"FormatUint":	ValueOf(strconv.FormatUint),
		"IntSize":	ValueOf(strconv.IntSize),
		"IsGraphic":	ValueOf(strconv.IsGraphic),
		"IsPrint":	ValueOf(strconv.IsPrint),
		"Itoa":	ValueOf(strconv.Itoa),
		"ParseBool":	ValueOf(strconv.ParseBool),
		"ParseFloat":	ValueOf(strconv.ParseFloat),
		"ParseInt":	ValueOf(strconv.ParseInt),
		"ParseUint":	ValueOf(strconv.ParseUint),
		"Quote":	ValueOf(strconv.Quote),
		"QuoteRune":	ValueOf(strconv.QuoteRune),
		"QuoteRuneToASCII":	ValueOf(strconv.QuoteRuneToASCII),
		"QuoteRuneToGraphic":	ValueOf(strconv.QuoteRuneToGraphic),
		"QuoteToASCII":	ValueOf(strconv.QuoteToASCII),
		"QuoteToGraphic":	ValueOf(strconv.QuoteToGraphic),
		"Unquote":	ValueOf(strconv.Unquote),
		"UnquoteChar":	ValueOf(strconv.UnquoteChar),
	},
	Types: map[string]Type{
		"NumError":	TypeOf((*strconv.NumError)(nil)).Elem(),
	},
	Proxies: map[string]Type{
	},
	Wrappers: map[string][]string{
	} }
}
