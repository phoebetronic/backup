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

func ren(ind int, bin int, tin int, lef []float32, rig []float32) []byte {
	var sta plotter.XYs
	{
		sta = append(sta, plotter.XY{X: float64(len(lef) + 1), Y: float64(rig[0])})
	}

	var bot plotter.XYs
	{
		bot = append(bot, plotter.XY{X: float64(len(lef) + bin), Y: float64(rig[bin])})
	}

	var top plotter.XYs
	{
		top = append(top, plotter.XY{X: float64(len(lef) + tin), Y: float64(rig[tin])})
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
		addsca(plo, sta, 1, "sta")
		addsca(plo, bot, 2, "bot")
		addsca(plo, top, 3, "top")
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

func addsca(plo *plot.Plot, xys plotter.XYs, ind int, des string) {
	t, err := plotter.NewScatter(xys)
	if err != nil {
		panic(err)
	}
	t.GlyphStyle.Color = plotutil.Color(ind)
	t.GlyphStyle.Radius = 5
	t.GlyphStyle.Shape = draw.PyramidGlyph{}
	plo.Add(t)
	plo.Legend.Add(des, t)
}
