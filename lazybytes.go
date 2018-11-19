// Package lazybytes implements a bytes.Reader which is initialized lazily.
// Otherwise it performs exactly like bytes.Reader.
//
// The constructor takes a function which is called to perform initialization
// when the Reader is accessed first time.
//
// This may be useful with http.ServeContent. By wrapping your content
// template rendering in lazybytes.Reader the rendering will be skipped if
// the content is not needed (for example if the client makes an
// If-Modified-Since request).
package lazybytes

import (
	"bytes"
	"io"
)

// Reader implements the io.Reader, io.ReaderAt, io.WriterTo, io.Seeker,
// io.ByteScanner and io.RuneScanner interfaces.
type Reader struct {
	br       *bytes.Reader
	initFunc func() []byte
}

// NewReader takes a function and returns a Reader. The supplied
// initialization function is called to initialize the Reader contents at
// the time when when it is first used.
func NewReader(f func() []byte) *Reader {
	return &Reader{
		initFunc: f,
	}
}

func (r *Reader) initialize() {
	// note that this is not thread safe
	// change to use sync.Once if thread safety is required
	if r.br == nil {
		r.br = bytes.NewReader(r.initFunc())
	}
}

func (r *Reader) Len() int {
	r.initialize()
	return r.br.Len()
}

func (r *Reader) Read(b []byte) (n int, err error) {
	r.initialize()
	return r.br.Read(b)
}

func (r *Reader) ReadAt(b []byte, off int64) (n int, err error) {
	r.initialize()
	return r.br.ReadAt(b, off)
}

func (r *Reader) ReadByte() (b byte, err error) {
	r.initialize()
	return r.br.ReadByte()
}

func (r *Reader) ReadRune() (ch rune, size int, err error) {
	r.initialize()
	return r.br.ReadRune()
}

func (r *Reader) Seek(offset int64, whence int) (int64, error) {
	r.initialize()
	return r.br.Seek(offset, whence)
}

func (r *Reader) Size() int64 {
	r.initialize()
	return r.br.Size()
}

func (r *Reader) UnreadByte() error {
	r.initialize()
	return r.br.UnreadByte()
}

func (r *Reader) UnreadRune() error {
	r.initialize()
	return r.br.UnreadRune()
}

func (r *Reader) WriteTo(w io.Writer) (n int64, err error) {
	r.initialize()
	return r.br.WriteTo(w)
}
