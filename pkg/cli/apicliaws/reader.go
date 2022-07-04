package apicliaws

import (
	"fmt"
	"os"
	"sync"
)

type Reader struct {
	fp      *os.File
	size    int64
	read    int64
	signMap map[int64]struct{}
	mux     sync.Mutex
}

func (r *Reader) Read(p []byte) (int, error) {
	return r.fp.Read(p)
}

func (r *Reader) ReadAt(p []byte, off int64) (int, error) {
	n, err := r.fp.ReadAt(p, off)
	if err != nil {
		return n, err
	}

	r.mux.Lock()
	// Ignore the first signature call
	if _, ok := r.signMap[off]; ok {
		// Got the length have read( or means has uploaded), and you can construct your message
		r.read += int64(n)
		fmt.Printf("\rtotal read:%d    progress:%d%%", r.read, int(float32(r.read*100)/float32(r.size)))
	} else {
		r.signMap[off] = struct{}{}
	}
	r.mux.Unlock()
	return n, err
}

func (r *Reader) Seek(offset int64, whence int) (int64, error) {
	return r.fp.Seek(offset, whence)
}
