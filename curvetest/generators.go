package curvetest

import (
	"fmt"
	"github.com/struckoff/kvstore/router/nodehasher"
	"github.com/struckoff/sfcdistribution/node"
	balancer "github.com/struckoff/sfcframework"
	"github.com/struckoff/sfcframework/curve"
	"github.com/struckoff/sfcframework/optimizer"
	"github.com/struckoff/sfcframework/transform"
)

func BitsMortonTest(step, max uint64, nodesAmount int, capacity float64, dis []balancer.DataItem) (res []CurveTestConfig) {
	nameForm := "Morton.Stateless[2x%d]"

	tpl := CurveTestConfig{
		Name:      "",
		Stateless: true,
		Curve:     curve.Morton,
		Dims:      2,
		Bits:      0,
		Tf:        transform.SpaceTransform,
		Of:        optimizer.RangeOptimizer,
		Nodes:     nil,
		Dis:       dis,
	}

	for tpl.Bits < max {
		tpl.Bits += step
		tpl.Name = fmt.Sprintf(nameForm, 1<<tpl.Bits)
		hsfc, err := curve.NewCurve(tpl.Curve, 2, tpl.Bits)
		if err != nil {
			panic(err)
		}

		hasher := nodehasher.NewGeoSfc(hsfc)
		tpl.Nodes = node.GeoEqualSetup(nodesAmount, capacity, hasher)
		res = append(res, tpl)
	}
	return res
}

func StateMortonTest(bits uint64, nodesAmount int, capacity float64, dis []balancer.DataItem) (res []CurveTestConfig) {
	nameForm := "Morton.%s[2x%d]"

	tpl := CurveTestConfig{
		Name:  "",
		Curve: curve.Morton,
		Dims:  2,
		Bits:  bits,
		Tf:    transform.SpaceTransform,
		Nodes: nil,
		Dis:   dis,
	}

	hsfc, err := curve.NewCurve(tpl.Curve, 2, bits)
	if err != nil {
		panic(err)
	}
	hasher := nodehasher.NewGeoSfc(hsfc)
	tpl.Nodes = node.GeoEqualSetup(nodesAmount, capacity, hasher)

	tpl.Stateless = true
	tpl.Of = optimizer.RangeOptimizer
	tpl.Name = fmt.Sprintf(nameForm, "Stateless", 1<<bits)

	res = append(res, tpl)

	tpl.Stateless = false
	tpl.Of = optimizer.PowerRangeOptimizer
	tpl.Name = fmt.Sprintf(nameForm, "Statefull", 1<<bits)

	res = append(res, tpl)
	return res
}
func StateHilbertTest(bits uint64, nodesAmount int, capacity float64, dis []balancer.DataItem) (res []CurveTestConfig) {
	nameForm := "Hilbert.%s[2x%d]"

	tpl := CurveTestConfig{
		Name:  "",
		Curve: curve.Hilbert,
		Dims:  2,
		Bits:  bits,
		Tf:    transform.SpaceTransform,
		Dis:   dis,
	}

	hsfc, err := curve.NewCurve(tpl.Curve, 2, bits)
	if err != nil {
		panic(err)
	}
	hasher := nodehasher.NewGeoSfc(hsfc)
	tpl.Nodes = node.GeoEqualSetup(nodesAmount, capacity, hasher)

	tpl.Stateless = true
	tpl.Of = optimizer.RangeOptimizer
	tpl.Name = fmt.Sprintf(nameForm, "Stateless", 1<<bits)

	res = append(res, tpl)

	tpl.Stateless = false
	tpl.Of = optimizer.PowerRangeOptimizer
	tpl.Name = fmt.Sprintf(nameForm, "Statefull", 1<<bits)

	res = append(res, tpl)
	return res
}
