package curvetest

import (
	balancer "github.com/struckoff/sfcframework"
	"github.com/struckoff/sfcframework/curve"
)

type CurveTestConfig struct {
	Name      string
	Stateless bool
	Curve     curve.CurveType
	Dims      uint64
	Bits      uint64
	Tf        balancer.TransformFunc
	Of        balancer.OptimizerFunc
	Nodes     []balancer.Node
	Dis       []balancer.DataItem
}
