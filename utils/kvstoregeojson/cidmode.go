package main

import (
	geojson "github.com/paulmach/go.geojson"
	"os"
	"sort"
	"strconv"
)

func cidmode(srv *string) {
	ks, err := download(*srv + "/cid")
	if err != nil {
		panic(err)
	}

	var geos []*geojson.Feature
	geos = cidunpack(ks)
	fc := geojson.NewFeatureCollection()
	for iter := range geos {
		fc.AddFeature(geos[iter])
	}
	b, err := fc.MarshalJSON()
	if err != nil {
		panic(err)
	}

	if _, err := os.Stdout.Write(b); err != nil {
		panic(err)
	}
}

func cellcompress(ps []Point) (p Point) {
	if len(ps) == 0 {
		return
	}

	sort.Slice(ps, func(i, j int) bool {
		return ps[i].Lat < ps[j].Lat
	})

	midLat := ps[len(ps)/2].Lat

	sort.Slice(ps, func(i, j int) bool {
		return ps[i].Lon < ps[j].Lon
	})

	midLon := ps[len(ps)/2].Lon

	return Point{Lon: midLon, Lat: midLat}
}

func cidunpack(ks Keys) []*geojson.Feature {
	var geos []*geojson.Feature

	ids := []string{}
	for nodeID := range ks.Nodes {
		ids = append(ids, nodeID)
	}
	sort.Slice(ids, func(i, j int) bool {
		a, err := strconv.Atoi(ids[i])
		if err != nil {
			panic(err)
		}
		b, err := strconv.Atoi(ids[j])
		if err != nil {
			panic(err)
		}
		return a < b
	})

	points := make([][]float64, len(ids))
	for iter, cid := range ids {
		p := cellcompress(ks.Nodes[cid])
		points[iter] = []float64{p.Lon, p.Lat}
	}

	g := geojson.NewLineStringFeature(points)
	geos = append(geos, g)
	return geos
}
