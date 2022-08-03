package win

import (
	"bytes"
	"encoding/csv"
	"io/ioutil"
	"os"
)

func (r *run) app(p string, n string, s string) {
	var a []byte
	{
		f, err := ioutil.ReadFile(n)
		if err != nil {
			panic(err)
		}

		c, err := csv.NewReader(bytes.NewReader(f)).ReadAll()
		if err != nil {
			panic(err)
		}

		var l [][]string
		for _, v := range c {
			if v[0] == s {
				v[0] = "1"
			} else {
				v[0] = "0"
			}

			{
				l = append(l, v)
			}
		}

		b := bytes.NewBuffer([]byte{})
		w := csv.NewWriter(b)

		for _, s := range l {
			err := w.Write(s)
			if err != nil {
				panic(err)
			}
		}

		{
			w.Flush()
			err := w.Error()
			if err != nil {
				panic(err)
			}
		}

		{
			a = b.Bytes()
		}
	}

	// Append the training data we collected to the bucket specific file.
	{
		f, err := os.OpenFile(p, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			panic(err)
		}

		_, err = f.Write(a)
		if err != nil {
			f.Close() // ignore error; Write error takes precedence
			panic(err)
		}

		err = f.Close()
		if err != nil {
			panic(err)
		}
	}
}
