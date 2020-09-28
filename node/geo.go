package node

import (
	"fmt"
	"github.com/OneOfOne/xxhash"
	"github.com/struckoff/kvstore/router/nodehasher"
	"github.com/struckoff/kvstore/router/rpcapi"
	"github.com/struckoff/sfcdistribution/dataset"
	balancer "github.com/struckoff/sfcframework"
)

type GeoNode struct {
	id     string
	geo    *rpcapi.GeoData
	p      float64
	c      float64
	hasher nodehasher.Hasher
}

func (n *GeoNode) ID() string {
	//return fmt.Sprintf("%f,%f", n.geo.Latitude, n.geo.Longitude)
	return n.id
}
func (n *GeoNode) Power() balancer.Power {
	return Power(n.p)
}
func (n *GeoNode) Capacity() balancer.Capacity {
	return Capacity(n.c)
}
func (n *GeoNode) Hash() uint64 {
	if n.geo == nil {
		hasher := xxhash.New64()
		_, err := hasher.WriteString(n.id)
		if err != nil {
			panic(err)
		}
		return hasher.Sum64()
	}
	code, err := n.hasher.Sum(&rpcapi.NodeMeta{
		ID:       n.id,
		Power:    n.p,
		Capacity: n.c,
		Geo:      n.geo,
	})
	if err != nil {
		panic(err)
	}
	return code
}

func NewGeoNode(id string, p, c float64, geo *rpcapi.GeoData, hasher nodehasher.Hasher) *GeoNode {
	return &GeoNode{id, geo, p, c, hasher}
}

func GeoEqualSetup(n int, capacity float64, hasher nodehasher.Hasher) []balancer.Node {
	ns := make([]balancer.Node, n)
	latIter := (dataset.MaxLatitude - dataset.MinLatitude) / float64(n)
	lonIter := (dataset.MaxLatitude - dataset.MinLatitude) / float64(n)
	for iter := range ns {
		geo := &rpcapi.GeoData{
			Latitude:  dataset.MinLatitude + latIter*float64(iter),
			Longitude: dataset.MinLongtitude + lonIter*float64(iter),
		}
		ns[iter] = NewGeoNode(fmt.Sprintf("node-%d", iter), 1, capacity, geo, hasher)
	}
	return ns
}

func GeoSetupCaps(caps []float64, hasher nodehasher.Hasher) []balancer.Node {
	ns := make([]balancer.Node, len(caps))
	latIter := (dataset.MaxLatitude - dataset.MinLatitude) / float64(len(caps))
	lonIter := (dataset.MaxLatitude - dataset.MinLatitude) / float64(len(caps))
	for iter := range ns {
		geo := &rpcapi.GeoData{
			Latitude:  dataset.MinLatitude + latIter*float64(iter),
			Longitude: dataset.MinLongtitude + lonIter*float64(iter),
		}
		ns[iter] = NewGeoNode(fmt.Sprintf("node-%d", iter), 1, caps[iter], geo, hasher)
	}
	return ns
}
