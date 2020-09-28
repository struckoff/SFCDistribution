package report

import (
	"errors"
	geojson "github.com/paulmach/go.geojson"
	balancer "github.com/struckoff/sfcframework"
)

func NewGeoReport(colors map[string][2]string) *GeoReport {
	return &GeoReport{
		fc:     geojson.NewFeatureCollection(),
		colors: colors,
	}
}

type GeoReport struct {
	fc     *geojson.FeatureCollection
	colors map[string][2]string
}

func (r *GeoReport) Add(node balancer.Node, di balancer.DataItem) error {
	vals := di.Values()
	if len(vals) < 2 {
		return errors.New("not enough values in dataitem")
	}
	lat, ok := vals[0].(float64)
	if !ok {
		return errors.New("wrong lon type")
	}
	lon, ok := vals[1].(float64)
	if !ok {
		return errors.New("wrong lat type")
	}
	f := geojson.NewPointFeature([]float64{lon, lat})
	f.SetProperty("node", node.ID())
	f.SetProperty("circle-color", r.colors[node.ID()][0])
	f.SetProperty("circle-altcolor", r.colors[node.ID()][1])
	r.fc.AddFeature(f)
	return nil
}

func (r *GeoReport) Print(opts ...string) (string, error) {
	b, err := r.fc.MarshalJSON()
	if err != nil {
		return "", err
	}
	return string(b), nil
}
