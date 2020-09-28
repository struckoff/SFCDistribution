package main

import (
	"errors"
	"github.com/fogleman/gg"
	balancer "github.com/struckoff/sfcframework"
	"github.com/struckoff/sfcframework/curve"
	balancernode "github.com/struckoff/sfcframework/node"
	"log"
	"math"
	"math/rand"
)

func main() {
	//if err := capacityOptimizer(); err != nil {
	//	panic(err)
	//}
	//if err := DrawSplitCurve(curve.Morton, 4, []float64{0.1, 0.2, 0.3, 0.4, 0.5, 0.6, 0.7, 0.8, 0.9, 1.0}, "mor.png"); err != nil {
	//	panic(err)
	//}
	if err := DrawSplitCurve(curve.Morton, 4, []float64{0.1, 0.2, 0.3, 0.4, 0.5, 0.6, 0.7, 0.8, 0.9, 1.0}, "mor.png"); err != nil {
		panic(err)
	}
	return
}

func DrawCurve(cType curve.CurveType, bits uint64, op string) error {
	dims := uint64(2)
	c, err := curve.NewCurve(cType, dims, bits)
	if err != nil {
		return err
	}
	dcSize := 2048
	dc := gg.NewContext(dcSize, dcSize)
	dc.SetRGB(1, 1, 1)
	dc.SetLineWidth(10)
	maxSize := 1 << bits
	cSize := float64(dcSize / maxSize)
	maxCode := uint64((1 << (dims * bits)) - 1)
	sx, sy := -1.0, -1.0
	for idx := uint64(0); idx <= maxCode; idx++ {
		cs, err := c.Decode(idx)
		if err != nil {
			return err
		}
		x := float64(cs[0])*cSize + cSize/2
		y := float64(cs[1])*cSize + cSize/2
		if sx != -1 {
			dc.DrawLine(sx, sy, x, y)
			dc.Stroke()
		}
		sx = x
		sy = y
	}
	return dc.SavePNG(op)
}

func DrawSplitCurve(cType curve.CurveType, bits uint64, splits []float64, op string) error {
	dims := uint64(2)
	c, err := curve.NewCurve(cType, dims, bits)
	if err != nil {
		return err
	}
	dcSize := 512
	dc := gg.NewContext(dcSize, dcSize)
	dc.SetRGB(1, 1, 1)
	dc.DrawRectangle(0, 0, float64(dcSize), float64(dcSize))
	dc.Fill()
	dc.SetLineWidth(7)
	maxSize := (1 << bits)
	cSize := float64(dcSize / maxSize)
	maxCode := uint64((1 << (dims * bits)) - 1)
	sx, sy := -1.0, -1.0
	si := 0
	r := rand.Float64()
	g := rand.Float64()
	b := rand.Float64()
	dc.SetRGB(r, g, b)
	for idx := uint64(0); idx <= maxCode; idx++ {
		p := float64(idx) / float64(maxCode)
		if splits != nil && p > splits[si] {
			r := rand.Float64()
			g := rand.Float64()
			b := rand.Float64()
			dc.SetRGB(r, g, b)
			si++
		}
		cs, err := c.Decode(idx)
		if err != nil {
			return err
		}
		x := float64(cs[0])*cSize + cSize/2
		y := float64(cs[1])*cSize + cSize/2
		if sx != -1 {
			dc.DrawLine(sx, sy, x, y)
			dc.Stroke()
		}
		sx = x
		sy = y
	}
	dc.SavePNG(op)
	return nil
}

//func capacityOptimizer() error {
//	stat := make(map[string]int)
//	defer func() {
//		fmt.Println(stat)
//	}()
//
//	//Prepare hasher
//	sfc, err := curve.NewCurve(curve.Morton, 2, 4)
//	if err != nil {
//		return err
//	}
//
//	hasher := nodehasher.NewGeoSfc(sfc)
//
//	//Build nodes with different capacities
//	nodes := node.GeoSetupCaps([]float64{1500, 1500, 1500, 1500, 1500, 1500, 1500, 1500, 1500, 1500, 1500}, hasher)
//	//Build nodes with equal capacities
//	//nodes := node.GeoEqualSetup(11, 1000, hasher)
//
//	//Build balancer
//	bal, err := balancer.NewBalancer(curve.Hilbert, 2, 16, transform.SpaceTransform, optimizer.CapPowerOptimizer, nil)
//	if err != nil {
//		return err
//	}
//
//	//Add nodes with loads to balancer
//	for _, node := range nodes {
//		//Generate 400 DataItems
//		if err != nil {
//			return err
//		}
//		fmt.Println("===================")
//		if err := bal.AddNode(node, true); err != nil {
//			return err
//		}
//		fmt.Println("===================")
//		if err := bal.Optimize(); err != nil {
//			return err
//		}
//	}
//	dis, err := dataset.GeoClustersWithNoise("4", 20000, .3)
//	for iter, di := range dis {
//		//add data considering cell loads and perform repalcing if needed
//		if n, _, err := addData(bal, di); err != nil {
//			log.Println("ERRR", iter, err.Error())
//			return err
//		} else {
//			stat[n.ID()]++
//		}
//	}
//	return nil
//}

//func capacityOptimizer_preLoad() error {
//	//Prepare hasher
//	sfc, err := curve.NewCurve(curve.Hilbert, 2, 4)
//	if err != nil {
//		return err
//	}
//
//	hasher := nodehasher.NewGeoSfc(sfc)
//
//	//Build nodes with different capacities
//	nodes := node.GeoSetupCaps([]float64{1000, 500, 300, 100, 10000}, hasher)
//	//Build nodes with equal capacities
//	//nodes := node.GeoEqualSetup(11, 1000, hasher)
//
//	//Build balancer
//	bal, err := balancer.NewBalancer(curve.Hilbert, 2, 16, transform.SpaceTransform, optimizer.CapPowerOptimizer, nil)
//	if err != nil {
//		return err
//	}
//
//	//Add nodes with loads to balancer
//	for _, node := range nodes {
//		//Generate 400 DataItems
//		dis, err := dataset.GeoClustersWithNoise("4", 400, .3)
//		if err != nil {
//			return err
//		}
//		fmt.Println("===================")
//		if err := bal.AddNode(node, true); err != nil {
//			return err
//		}
//		for _, di := range dis {
//			//add data considering cell loads and perform repalcing if needed
//			if _, _, err := addData(bal, di); err != nil {
//				return err
//			}
//		}
//		fmt.Println("===================")
//		if err := bal.Optimize(); err != nil {
//			return err
//		}
//	}
//	return nil
//}

func addData(bal *balancer.Balancer, di balancer.DataItem) (balancernode.Node, uint64, error) {
	nb, cid, err := bal.LocateData(di)
	if err != nil {
		return nil, 0, err
	}

	ok, err := checkNodeCapacity(bal, nb, di)
	if err != nil {
		return nil, 0, err
	}
	if !ok {
		ncID, err := findBetterCell(bal, di, cid)
		if err != nil {
			return nil, 0, err
		}
		nb, cid, err = bal.RelocateData(di, ncID)
		if err != nil {
			return nil, 0, err
		}
	} else {
		nb, cid, err = bal.AddData(di)
		if err != nil {
			return nil, 0, err
		}
	}
	n, ok := nb.(balancernode.Node)
	if !ok {
		return nil, 0, errors.New("wrong node type")
	}
	return n, cid, nil
}

func checkNodeCapacity(bal *balancer.Balancer, n balancernode.Node, di balancer.DataItem) (bool, error) {
	cgs := bal.Space().CellGroups()
	c, err := n.Capacity().Get()
	if err != nil {
		return false, err
	}
	nf := true
	for iter := range cgs {
		if cgs[iter].Node().ID() == n.ID() {
			nf = false
			diff := c - float64(cgs[iter].TotalLoad()) - float64(di.Size())
			if diff > 0 {
				return true, err
			}
			return false, nil
		}
	}
	if nf {
		return false, errors.New("cell group not found")
	}
	return false, nil
}

func findBetterCell(bal *balancer.Balancer, di balancer.DataItem, cid uint64) (uint64, error) {
	dis := math.MaxInt64
	ncID := cid
	//var cg *balancer.CellGroup

	cgs := bal.Space().CellGroups()
	for iter := range cgs {
		l := cgs[iter].TotalLoad()
		c, err := cgs[iter].Node().Capacity().Get()
		if err != nil {
			log.Println(err)
			continue
		}
		dc := c - float64(l) - float64(di.Size())
		if dc < 0 {
			continue
		}
		if cgs[iter].Range().Len <= 0 {
			continue
		}

		// find closest cell to filled group
		if cgs[iter].Range().Max <= cid {
			if lft := int(cid) - int(cgs[iter].Range().Max); lft < dis {
				dis = lft
				ncID = cgs[iter].Range().Max - 1 //closest cell in available group
			}
		} else if cgs[iter].Range().Min > cid {
			if rght := int(cgs[iter].Range().Min) - int(cid); rght < dis {
				dis = rght
				ncID = cgs[iter].Range().Min //closest cell in available group
			}
		}
	}

	if dis == math.MaxInt64 {
		return 0, errors.New("out of capacity")
	}
	if cid == ncID {
		return 0, errors.New("appropriate cell not found")
	}

	return ncID, nil
}
