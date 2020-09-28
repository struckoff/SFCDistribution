package dataset

import (
	"github.com/pkg/errors"
	"github.com/struckoff/sfcdistribution/dataset/dataitem"
	balancer "github.com/struckoff/sfcframework"
	"github.com/struckoff/sfcframework/curve"
	"math/rand"
)

const (
	MaxLongtitude = 180.0
	MinLongtitude = -180.0
	MaxLatitude   = 90.0
	MinLatitude   = -90.0
)

var clusterModes = map[string][][4]float64{
	"world": {
		{MinLatitude, MaxLatitude, MinLongtitude, MaxLongtitude},
	},
	"2": {
		{MinLatitude, 0, MinLongtitude, 0},
		{0, MaxLatitude, 0, MaxLongtitude},
	},
	"4v": {
		{MinLatitude, MaxLatitude, 0, 0},
		{0, 0, MinLongtitude, MaxLongtitude},
		{MinLatitude, 0, MinLongtitude, 0},
		{0, MaxLatitude, 0, MaxLongtitude},
	},
	"4": {
		{0, MaxLatitude, MinLongtitude, 0},
		{0, MaxLatitude, 0, MaxLongtitude},
		{MinLatitude, 0, MinLongtitude, 0},
		{MinLatitude, 0, 0, MaxLongtitude},
	},
}

func GeoPoint(lat, lon float64) (balancer.DataItem, error) {
	di, err := dataitem.NewSpace(lat, lon)
	if err != nil {
		return nil, err
	}
	return di, nil
}

func RandGeoPoint(minLat, maxLat, minLon, maxLon float64) (balancer.DataItem, error) {
	lat := minLat + rand.Float64()*(maxLat-minLat)
	lon := minLon + rand.Float64()*(maxLon-minLon)
	di, err := dataitem.NewSpace(lat, lon)
	if err != nil {
		return nil, err
	}
	return di, nil
}
func GeoClusters(mode string, amount int) (res []balancer.DataItem, err error) {
	sets, ok := clusterModes[mode]
	if !ok {
		return nil, errors.New("mode not found")
	}
	res = make([]balancer.DataItem, amount)
	for iter := range res {
		ps := sets[iter%len(sets)]
		res[iter], err = RandGeoPoint(ps[0], ps[1], ps[2], ps[3])
		if err != nil {
			return nil, err
		}

	}
	return res, nil
}

func GeoClustersWithNoise(mode string, amount int, noiserate float64) (res []balancer.DataItem, err error) {
	res = make([]balancer.DataItem, amount)
	var ps [4]float64

	sets := clusterModes[mode]
	for iter := range res {
		if rand.Float64() < noiserate {
			ps = clusterModes["world"][0]
		} else {
			setIdx := rand.Intn(len(sets))
			ps = sets[setIdx]
		}
		res[iter], err = RandGeoPoint(ps[0], ps[1], ps[2], ps[3])
		if err != nil {
			return nil, err
		}
	}
	return res, nil
}

const latStep = 90.0
const lonStep = 180.0

func RawCurve(sfc curve.Curve) (res []balancer.DataItem, err error) {
	for cid := uint64(0); cid < sfc.Length(); cid++ {
		coords, _ := sfc.Decode(cid)
		dimSize := sfc.DimensionSize()
		lat := float64(coords[0])/float64(dimSize)*2*latStep - latStep
		lon := float64(coords[1])/float64(dimSize)*2*lonStep - lonStep

		p, err := GeoPoint(lat, lon)
		if err != nil {
			return nil, err
		}

		res = append(res, p)
	}
	return
}
