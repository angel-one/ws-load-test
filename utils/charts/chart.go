package charts

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

func DrawChart(data []float64, timeSeries []time.Time, xAxis string, yAxis string) *chart.Chart{

	graph := chart.Chart{
		XAxis: chart.XAxis{
			Name:      xAxis,
			NameStyle: chart.StyleShow(),
			Style:     chart.StyleShow(),
		},
		YAxis: chart.YAxis{
			Name:      yAxis,
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
				XValues: makeRange(1, len(data)),
				YValues: data,
			},
		},
	}

	return &graph
}

func RenderHTML(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Test!")
}
