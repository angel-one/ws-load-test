package chart

import (
	"fmt"
	"github.com/wcharczuk/go-chart"
	"net/http"
	"time"
)

func makeRange(min, max int) []float64 {
	a := make([]float64, max-min+1)
	for i := range a {
		a[i] = float64(min) + float64(i)
	}
	return a
}

func DrawChart(res http.ResponseWriter, req *http.Request, latency []float64, timeSeries []time.Time) {

	graph := chart.Chart{
		XAxis: chart.XAxis{
			Name:      "Requests Index",
			NameStyle: chart.StyleShow(),
			Style:     chart.StyleShow(),
		},
		YAxis: chart.YAxis{
			Name:      "Latency Count",
			NameStyle: chart.StyleShow(),
			Style:     chart.StyleShow(),
		},
		Series: []chart.Series{
			chart.ContinuousSeries{
				Style: chart.Style{
					Show:        true,
					StrokeColor: chart.GetDefaultColor(0).WithAlpha(64),
					FillColor:   chart.GetDefaultColor(0).WithAlpha(64),
				},
				XValues: makeRange(1, len(latency)),
				YValues: latency,
			},
		},
	}

	res.Header().Set("Content-Type", "image/png")
	graph.Render(chart.PNG, res)
}

func renderHTML(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello Carrot!</h1>")
}
