package apicliaws

import (
	"fmt"
	"sync/atomic"

	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
)

type Writer struct {
	siz int64
	tot int64
	wri *manager.WriteAtBuffer
}

func (w *Writer) WriteAt(byt []byte, off int64) (int, error) {
	atomic.AddInt64(&w.tot, int64(len(byt)))

	fmt.Printf("\rbuffered %d%%", int(float32(w.tot*100)/float32(w.siz)))

	return w.wri.WriteAt(byt, off)
}
