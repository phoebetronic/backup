package mod

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"math"
	"math/rand"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/phoebetron/backup/pkg/ind"
)

const (
	fratra = 0.70
	frates = 0.15
	fraval = 0.15
)

func (r *run) ful(pat string) {
	{
		r.dir(filepath.Join(pat, "csv"))
	}

	var ndx ind.Index
	{
		ndx = ind.Read(pat)
	}

	var enc map[string]string
	{
		enc = ndx.EncI()
	}

	var ful [][]string

	for _, i := range ndx.Lis.Shfl() {
		f, err := ioutil.ReadFile(i.Fil)
		if err != nil {
			panic(err)
		}

		c, err := csv.NewReader(bytes.NewReader(f)).ReadAll()
		if err != nil {
			panic(err)
		}

		for _, v := range c {
			ful = append(ful, r.enc(enc, v))
		}
	}

	// Shuffle and split the full file for all combined and one hot encoded
	// training data.
	{
		rand.Seed(time.Now().UnixNano())
		rand.Shuffle(len(ful), func(i, j int) { ful[i], ful[j] = ful[j], ful[i] })
	}

	{
		r.dir(filepath.Join(pat, "ful"))
	}

	{
		w := 0
		x := int(float64(len(ful)) * (fratra))
		y := int(float64(len(ful)) * (fratra + frates))
		z := int(float64(len(ful)) * (fratra + frates + fraval))

		r.wri(ful[w:x], filepath.Join(pat, "ful", "tra.csv"))
		r.wri(ful[x:y], filepath.Join(pat, "ful", "tes.csv"))
		r.wri(ful[y:z], filepath.Join(pat, "ful", "val.csv"))
	}

	for _, s := range ndx.Sta {
		var b float32
		var c int
		{
			b = s.Buc
			c = s.Cou
		}

		// Format the bucket string representation for comparision during the
		// process of appending to the training data file. The precision of 5 is
		// important at this point because the first string in each raw CSV file
		// has a precision of 5 already.
		var n string
		{
			n = strconv.FormatFloat(float64(b), 'f', 5, 32)
		}

		{
			fmt.Printf("preparing training data for bucket %s\n", n)
		}

		// Calculate the amount of samples we need to collect for each of the
		// other buckets.
		var o int
		{
			o = int(math.Min(float64(c), float64(ndx.SumX(b)))) / (len(ndx.Sta) - 1)
		}

		// Prepare the target map we can draw from while filling up records of
		// training data.
		t := map[float32]int{}
		{
			t[b] = c

			for _, z := range ndx.Sta {
				if z.Buc != b {
					t[z.Buc] = o
				}
			}
		}

		var p string
		{
			p = filepath.Join(pat, "csv", nam(b, "ful"))
		}

		// Shuffle the index list so that we do not always take from the top
		// alone and iterate through the whole shuffled list.
		for _, i := range ndx.Lis.Shfl() {
			// Ignore files for buckets that we already collected enough
			// training data for.
			if t[i.Buc] <= 0 {
				continue
			}

			{
				t[i.Buc] -= i.Cou
			}

			{
				r.app(p, i.Fil, n)
			}
		}

		// At this point the full file for the current bucket got created. Now
		// shuffle and split the final training data files.
		var l [][]string
		{
			f, err := ioutil.ReadFile(p)
			if err != nil {
				panic(err)
			}

			l, err = csv.NewReader(bytes.NewReader(f)).ReadAll()
			if err != nil {
				panic(err)
			}
		}

		{
			rand.Seed(time.Now().UnixNano())
			rand.Shuffle(len(l), func(i, j int) { l[i], l[j] = l[j], l[i] })
		}

		{
			w := 0
			x := int(float64(len(l)) * (fratra))
			y := int(float64(len(l)) * (fratra + frates))
			z := int(float64(len(l)) * (fratra + frates + fraval))

			r.wri(l[w:x], filepath.Join(pat, "csv", nam(b, "tra")))
			r.wri(l[x:y], filepath.Join(pat, "csv", nam(b, "tes")))
			r.wri(l[y:z], filepath.Join(pat, "csv", nam(b, "val")))
		}

		{
			err := os.Remove(p)
			if err != nil {
				panic(err)
			}
		}
	}
}
