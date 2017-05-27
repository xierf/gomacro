// this file was generated by gomacro command: import _b "io"
// DO NOT EDIT! Any change will be lost when the file is re-generated

package imports

import (
	. "reflect"
	"io"
)

// reflection: allow interpreted code to import "io"
func init() {
	Packages["io"] = Package{
	Binds: map[string]Value{
		"Copy":	ValueOf(io.Copy),
		"CopyBuffer":	ValueOf(io.CopyBuffer),
		"CopyN":	ValueOf(io.CopyN),
		"EOF":	ValueOf(&io.EOF).Elem(),
		"ErrClosedPipe":	ValueOf(&io.ErrClosedPipe).Elem(),
		"ErrNoProgress":	ValueOf(&io.ErrNoProgress).Elem(),
		"ErrShortBuffer":	ValueOf(&io.ErrShortBuffer).Elem(),
		"ErrShortWrite":	ValueOf(&io.ErrShortWrite).Elem(),
		"ErrUnexpectedEOF":	ValueOf(&io.ErrUnexpectedEOF).Elem(),
		"LimitReader":	ValueOf(io.LimitReader),
		"MultiReader":	ValueOf(io.MultiReader),
		"MultiWriter":	ValueOf(io.MultiWriter),
		"NewSectionReader":	ValueOf(io.NewSectionReader),
		"Pipe":	ValueOf(io.Pipe),
		"ReadAtLeast":	ValueOf(io.ReadAtLeast),
		"ReadFull":	ValueOf(io.ReadFull),
		"SeekCurrent":	ValueOf(io.SeekCurrent),
		"SeekEnd":	ValueOf(io.SeekEnd),
		"SeekStart":	ValueOf(io.SeekStart),
		"TeeReader":	ValueOf(io.TeeReader),
		"WriteString":	ValueOf(io.WriteString),
	},
	Types: map[string]Type{
		"ByteReader":	TypeOf((*io.ByteReader)(nil)).Elem(),
		"ByteScanner":	TypeOf((*io.ByteScanner)(nil)).Elem(),
		"ByteWriter":	TypeOf((*io.ByteWriter)(nil)).Elem(),
		"Closer":	TypeOf((*io.Closer)(nil)).Elem(),
		"LimitedReader":	TypeOf((*io.LimitedReader)(nil)).Elem(),
		"PipeReader":	TypeOf((*io.PipeReader)(nil)).Elem(),
		"PipeWriter":	TypeOf((*io.PipeWriter)(nil)).Elem(),
		"ReadCloser":	TypeOf((*io.ReadCloser)(nil)).Elem(),
		"ReadSeeker":	TypeOf((*io.ReadSeeker)(nil)).Elem(),
		"ReadWriteCloser":	TypeOf((*io.ReadWriteCloser)(nil)).Elem(),
		"ReadWriteSeeker":	TypeOf((*io.ReadWriteSeeker)(nil)).Elem(),
		"ReadWriter":	TypeOf((*io.ReadWriter)(nil)).Elem(),
		"Reader":	TypeOf((*io.Reader)(nil)).Elem(),
		"ReaderAt":	TypeOf((*io.ReaderAt)(nil)).Elem(),
		"ReaderFrom":	TypeOf((*io.ReaderFrom)(nil)).Elem(),
		"RuneReader":	TypeOf((*io.RuneReader)(nil)).Elem(),
		"RuneScanner":	TypeOf((*io.RuneScanner)(nil)).Elem(),
		"SectionReader":	TypeOf((*io.SectionReader)(nil)).Elem(),
		"Seeker":	TypeOf((*io.Seeker)(nil)).Elem(),
		"WriteCloser":	TypeOf((*io.WriteCloser)(nil)).Elem(),
		"WriteSeeker":	TypeOf((*io.WriteSeeker)(nil)).Elem(),
		"Writer":	TypeOf((*io.Writer)(nil)).Elem(),
		"WriterAt":	TypeOf((*io.WriterAt)(nil)).Elem(),
		"WriterTo":	TypeOf((*io.WriterTo)(nil)).Elem(),
	},
	Proxies: map[string]Type{
		"ByteReader":	TypeOf((*ByteReader_io)(nil)).Elem(),
		"ByteScanner":	TypeOf((*ByteScanner_io)(nil)).Elem(),
		"ByteWriter":	TypeOf((*ByteWriter_io)(nil)).Elem(),
		"Closer":	TypeOf((*Closer_io)(nil)).Elem(),
		"ReadCloser":	TypeOf((*ReadCloser_io)(nil)).Elem(),
		"ReadSeeker":	TypeOf((*ReadSeeker_io)(nil)).Elem(),
		"ReadWriteCloser":	TypeOf((*ReadWriteCloser_io)(nil)).Elem(),
		"ReadWriteSeeker":	TypeOf((*ReadWriteSeeker_io)(nil)).Elem(),
		"ReadWriter":	TypeOf((*ReadWriter_io)(nil)).Elem(),
		"Reader":	TypeOf((*Reader_io)(nil)).Elem(),
		"ReaderAt":	TypeOf((*ReaderAt_io)(nil)).Elem(),
		"ReaderFrom":	TypeOf((*ReaderFrom_io)(nil)).Elem(),
		"RuneReader":	TypeOf((*RuneReader_io)(nil)).Elem(),
		"RuneScanner":	TypeOf((*RuneScanner_io)(nil)).Elem(),
		"Seeker":	TypeOf((*Seeker_io)(nil)).Elem(),
		"WriteCloser":	TypeOf((*WriteCloser_io)(nil)).Elem(),
		"WriteSeeker":	TypeOf((*WriteSeeker_io)(nil)).Elem(),
		"Writer":	TypeOf((*Writer_io)(nil)).Elem(),
		"WriterAt":	TypeOf((*WriterAt_io)(nil)).Elem(),
		"WriterTo":	TypeOf((*WriterTo_io)(nil)).Elem(),
	},
	Wrappers: map[string][]string{
	} }
}

// --------------- proxy for io.ByteReader ---------------
type ByteReader_io struct {
	Object	interface{}
	ReadByte_	func() (byte, error)
}
func (Proxy *ByteReader_io) ReadByte() (byte, error) {
	return Proxy.ReadByte_()
}

// --------------- proxy for io.ByteScanner ---------------
type ByteScanner_io struct {
	Object	interface{}
	ReadByte_	func() (byte, error)
	UnreadByte_	func() error
}
func (Proxy *ByteScanner_io) ReadByte() (byte, error) {
	return Proxy.ReadByte_()
}
func (Proxy *ByteScanner_io) UnreadByte() error {
	return Proxy.UnreadByte_()
}

// --------------- proxy for io.ByteWriter ---------------
type ByteWriter_io struct {
	Object	interface{}
	WriteByte_	func(c byte) error
}
func (Proxy *ByteWriter_io) WriteByte(c byte) error {
	return Proxy.WriteByte_(c)
}

// --------------- proxy for io.Closer ---------------
type Closer_io struct {
	Object	interface{}
	Close_	func() error
}
func (Proxy *Closer_io) Close() error {
	return Proxy.Close_()
}

// --------------- proxy for io.ReadCloser ---------------
type ReadCloser_io struct {
	Object	interface{}
	Close_	func() error
	Read_	func(p []byte) (n int, err error)
}
func (Proxy *ReadCloser_io) Close() error {
	return Proxy.Close_()
}
func (Proxy *ReadCloser_io) Read(p []byte) (n int, err error) {
	return Proxy.Read_(p)
}

// --------------- proxy for io.ReadSeeker ---------------
type ReadSeeker_io struct {
	Object	interface{}
	Read_	func(p []byte) (n int, err error)
	Seek_	func(offset int64, whence int) (int64, error)
}
func (Proxy *ReadSeeker_io) Read(p []byte) (n int, err error) {
	return Proxy.Read_(p)
}
func (Proxy *ReadSeeker_io) Seek(offset int64, whence int) (int64, error) {
	return Proxy.Seek_(offset, whence)
}

// --------------- proxy for io.ReadWriteCloser ---------------
type ReadWriteCloser_io struct {
	Object	interface{}
	Close_	func() error
	Read_	func(p []byte) (n int, err error)
	Write_	func(p []byte) (n int, err error)
}
func (Proxy *ReadWriteCloser_io) Close() error {
	return Proxy.Close_()
}
func (Proxy *ReadWriteCloser_io) Read(p []byte) (n int, err error) {
	return Proxy.Read_(p)
}
func (Proxy *ReadWriteCloser_io) Write(p []byte) (n int, err error) {
	return Proxy.Write_(p)
}

// --------------- proxy for io.ReadWriteSeeker ---------------
type ReadWriteSeeker_io struct {
	Object	interface{}
	Read_	func(p []byte) (n int, err error)
	Seek_	func(offset int64, whence int) (int64, error)
	Write_	func(p []byte) (n int, err error)
}
func (Proxy *ReadWriteSeeker_io) Read(p []byte) (n int, err error) {
	return Proxy.Read_(p)
}
func (Proxy *ReadWriteSeeker_io) Seek(offset int64, whence int) (int64, error) {
	return Proxy.Seek_(offset, whence)
}
func (Proxy *ReadWriteSeeker_io) Write(p []byte) (n int, err error) {
	return Proxy.Write_(p)
}

// --------------- proxy for io.ReadWriter ---------------
type ReadWriter_io struct {
	Object	interface{}
	Read_	func(p []byte) (n int, err error)
	Write_	func(p []byte) (n int, err error)
}
func (Proxy *ReadWriter_io) Read(p []byte) (n int, err error) {
	return Proxy.Read_(p)
}
func (Proxy *ReadWriter_io) Write(p []byte) (n int, err error) {
	return Proxy.Write_(p)
}

// --------------- proxy for io.Reader ---------------
type Reader_io struct {
	Object	interface{}
	Read_	func(p []byte) (n int, err error)
}
func (Proxy *Reader_io) Read(p []byte) (n int, err error) {
	return Proxy.Read_(p)
}

// --------------- proxy for io.ReaderAt ---------------
type ReaderAt_io struct {
	Object	interface{}
	ReadAt_	func(p []byte, off int64) (n int, err error)
}
func (Proxy *ReaderAt_io) ReadAt(p []byte, off int64) (n int, err error) {
	return Proxy.ReadAt_(p, off)
}

// --------------- proxy for io.ReaderFrom ---------------
type ReaderFrom_io struct {
	Object	interface{}
	ReadFrom_	func(r io.Reader) (n int64, err error)
}
func (Proxy *ReaderFrom_io) ReadFrom(r io.Reader) (n int64, err error) {
	return Proxy.ReadFrom_(r)
}

// --------------- proxy for io.RuneReader ---------------
type RuneReader_io struct {
	Object	interface{}
	ReadRune_	func() (r rune, size int, err error)
}
func (Proxy *RuneReader_io) ReadRune() (r rune, size int, err error) {
	return Proxy.ReadRune_()
}

// --------------- proxy for io.RuneScanner ---------------
type RuneScanner_io struct {
	Object	interface{}
	ReadRune_	func() (r rune, size int, err error)
	UnreadRune_	func() error
}
func (Proxy *RuneScanner_io) ReadRune() (r rune, size int, err error) {
	return Proxy.ReadRune_()
}
func (Proxy *RuneScanner_io) UnreadRune() error {
	return Proxy.UnreadRune_()
}

// --------------- proxy for io.Seeker ---------------
type Seeker_io struct {
	Object	interface{}
	Seek_	func(offset int64, whence int) (int64, error)
}
func (Proxy *Seeker_io) Seek(offset int64, whence int) (int64, error) {
	return Proxy.Seek_(offset, whence)
}

// --------------- proxy for io.WriteCloser ---------------
type WriteCloser_io struct {
	Object	interface{}
	Close_	func() error
	Write_	func(p []byte) (n int, err error)
}
func (Proxy *WriteCloser_io) Close() error {
	return Proxy.Close_()
}
func (Proxy *WriteCloser_io) Write(p []byte) (n int, err error) {
	return Proxy.Write_(p)
}

// --------------- proxy for io.WriteSeeker ---------------
type WriteSeeker_io struct {
	Object	interface{}
	Seek_	func(offset int64, whence int) (int64, error)
	Write_	func(p []byte) (n int, err error)
}
func (Proxy *WriteSeeker_io) Seek(offset int64, whence int) (int64, error) {
	return Proxy.Seek_(offset, whence)
}
func (Proxy *WriteSeeker_io) Write(p []byte) (n int, err error) {
	return Proxy.Write_(p)
}

// --------------- proxy for io.Writer ---------------
type Writer_io struct {
	Object	interface{}
	Write_	func(p []byte) (n int, err error)
}
func (Proxy *Writer_io) Write(p []byte) (n int, err error) {
	return Proxy.Write_(p)
}

// --------------- proxy for io.WriterAt ---------------
type WriterAt_io struct {
	Object	interface{}
	WriteAt_	func(p []byte, off int64) (n int, err error)
}
func (Proxy *WriterAt_io) WriteAt(p []byte, off int64) (n int, err error) {
	return Proxy.WriteAt_(p, off)
}

// --------------- proxy for io.WriterTo ---------------
type WriterTo_io struct {
	Object	interface{}
	WriteTo_	func(w io.Writer) (n int64, err error)
}
func (Proxy *WriterTo_io) WriteTo(w io.Writer) (n int64, err error) {
	return Proxy.WriteTo_(w)
}
