package apicliaws

import (
	"bytes"
	"fmt"
	"sync"
)

type Reader struct {
	mux sync.Mutex
	rea bytes.Reader
	siz int64
	tot int64
}

func (r *Reader) Read(byt []byte) (int, error) {
	return r.rea.Read(byt)
}

func (r *Reader) ReadAt(byt []byte, off int64) (int, error) {
	num, err := r.rea.ReadAt(byt, off)
	if err != nil {
		return num, err
	}

	r.mux.Lock()
	r.tot += int64(num)
	fmt.Printf("\ruploaded %d%%", int(float32(r.tot*100)/float32(r.siz)))
	r.mux.Unlock()

	return num, nil
}

func (r *Reader) Seek(off int64, whe int) (int64, error) {
	return r.rea.Seek(off, whe)
}
