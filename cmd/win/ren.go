package win

import (
	"bytes"
	"strconv"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/vg/draw"
)

func ren(ind int, sin int, lef []float32, rig []float32) []byte {
	var bot plotter.XYs
	{
		bot = append(bot, plotter.XY{X: float64(len(lef) + 1), Y: float64(rig[0])})
	}

	var top plotter.XYs
	{
		top = append(top, plotter.XY{X: float64(len(lef) + sin), Y: float64(rig[sin])})
	}

	var xys plotter.XYs
	for _, v := range lef {
		xys = append(xys, plotter.XY{X: float64(len(xys)), Y: float64(v)})
	}
	for _, v := range rig {
		xys = append(xys, plotter.XY{X: float64(len(xys)), Y: float64(v)})
	}

	act := bytes.NewBuffer([]byte{})
	plo := plot.New()

	plo.Title.Text = strconv.Itoa(ind) + ".gold.svg"
	plo.X.Label.Text = "T"
	plo.Y.Label.Text = "P"

	err := plotutil.AddLines(plo, xys)
	if err != nil {
		panic(err)
	}

	{
		b, err := plotter.NewScatter(bot)
		if err != nil {
			panic(err)
		}
		b.GlyphStyle.Color = plotutil.Color(1)
		b.GlyphStyle.Radius = 5
		b.GlyphStyle.Shape = draw.PyramidGlyph{}
		plo.Add(b)
		plo.Legend.Add("bot", b)
	}

	{
		t, err := plotter.NewScatter(top)
		if err != nil {
			panic(err)
		}
		t.GlyphStyle.Color = plotutil.Color(2)
		t.GlyphStyle.Radius = 5
		t.GlyphStyle.Shape = draw.PyramidGlyph{}
		plo.Add(t)
		plo.Legend.Add("top", t)
	}

	wri, err := plo.WriterTo(6*vg.Inch, 6*vg.Inch, "svg")
	if err != nil {
		panic(err)
	}

	_, err = wri.WriteTo(act)
	if err != nil {
		panic(err)
	}

	return act.Bytes()
}
