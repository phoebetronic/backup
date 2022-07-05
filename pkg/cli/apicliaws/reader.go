package apicliaws

import (
	"fmt"
	"os"
	"sync"
)

type Reader struct {
	fil *os.File
	mux sync.Mutex
	siz int64
	tot int64
}

func (r *Reader) Read(byt []byte) (int, error) {
	return r.fil.Read(byt)
}

func (r *Reader) ReadAt(byt []byte, off int64) (int, error) {
	num, err := r.fil.ReadAt(byt, off)
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
	return r.fil.Seek(off, whe)
}
