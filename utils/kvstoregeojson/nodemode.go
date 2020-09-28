package main

import (
	"encoding/json"
	"fmt"
	geojson "github.com/paulmach/go.geojson"
	"net/http"
	"os"
)

func nodesmode(srv, node *string) {
	ks, err := download(*srv + "/list")
	if err != nil {
		panic(err)
	}

	var geos []*geojson.Feature
	if *node != "" {
		colors := nodeColors(len(ks.Nodes))
		cs := colors[*node]
		geos = features(*node, ks.Nodes[*node], cs[0], cs[1])
	} else {
		geos = unpack(ks)
	}
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

func download(url string) (Keys, error) {
	var ks Keys
	r, err := http.Get(url)
	if err != nil {
		return ks, err
	}
	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(&ks.Nodes); err != nil {
		return ks, err
	}
	return ks, nil
}

func unpack(ks Keys) []*geojson.Feature {
	var geos []*geojson.Feature
	colors := nodeColors(len(ks.Nodes))
	for nodeID := range ks.Nodes {
		cs := colors[nodeID]
		geos = append(geos, features(nodeID, ks.Nodes[nodeID], cs[0], cs[1])...)
	}
	return geos
}

func features(nodeID string, points []Point, colors, altcolor string) []*geojson.Feature {
	geos := make([]*geojson.Feature, 0)
	for iter := range points {
		var p *geojson.Feature
		p = geojson.NewPointFeature([]float64{points[iter].Lon, points[iter].Lat})
		p.SetProperty("title", nodeID)
		p.SetProperty("circle-color", colors)
		p.SetProperty("circle-altcolor", altcolor)
		p.SetProperty("cluster-color", fmt.Sprintf("#%s", points[iter].ClusterColor))
		p.SetProperty("cluster", points[iter].Cluster)
		geos = append(geos, p)
	}
	return geos
}

func nodeColors(n int) map[string][2]string {
	step := maxColor / (n + 1)
	color := 0
	res := make(map[string][2]string)
	for iter := 0; iter < n; iter++ {
		color += step
		nodeID := fmt.Sprintf("node-%d", iter)
		code := fmt.Sprintf("#%x", color)
		altcode := fmt.Sprintf("#%x", maxColor-color)
		res[nodeID] = [2]string{code, altcode}
		//fmt.Println(nodeID, ":", altcode)
	}
	return res
}
