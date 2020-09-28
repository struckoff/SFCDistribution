package main

import (
	"flag"
	geojson "github.com/paulmach/go.geojson"
	"github.com/struckoff/sfcdistribution/dataset"
	"github.com/struckoff/sfcframework/curve"
	"os"
	"strings"
)

func main() {
	crv := flag.String("curve", "hilbert", "hilbert/morton")
	dims := flag.Uint64("dims", 2, "Number of dimensions")
	bits := flag.Uint64("bits", 4, "Number of bits to encode each dimension")

	flag.Parse()

	if err := run(*crv, *dims, *bits); err != nil {
		panic(err)
	}
}

func run(crv string, dims, bits uint64) error {
	//var geos []*geojson.Feature
	var points [][]float64

	var crvType curve.CurveType

	switch strings.ToLower(crv) {
	case "hilbert":
		crvType = curve.Hilbert
	case "morton":
		crvType = curve.Morton
	}

	sfc, err := curve.NewCurve(crvType, dims, bits)
	if err != nil {
		return err
	}

	dis, err := dataset.RawCurve(sfc)
	if err != nil {
		return err
	}

	for _, di := range dis {
		points = append(points, []float64{di.Values()[1].(float64), di.Values()[0].(float64)})
	}

	f := geojson.NewLineStringFeature(points)
	//geos = append(geos, f)

	b, err := f.MarshalJSON()
	if err != nil {
		return err
	}

	if _, err := os.Stdout.Write(b); err != nil {
		return err
	}

	return nil
}
