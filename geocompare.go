package main

import ()

//
//import (
//	"encoding/json"
//	"fmt"
//	"github.com/buraksezer/consistent"
//	balancer "github.com/struckoff/sfcframework"
//	"github.com/struckoff/sfcframework/curve"
//	"github.com/struckoff/sfcframework/optimizer"
//	"github.com/struckoff/sfcframework/transform"
//	"math/rand"
//)
//
//const (
//	GEO_DIMS = 2
//	GEO_SIZE = 8
//	//GEO_KEYS = 100_000
//	GEO_KEYS = 100_000
//)
//
//type SpaceDataItem struct {
//	Key string
//	Lat float64
//	Lon float64
//}
//
//func NewSpaceDataItem(key string) (SpaceDataItem, error) {
//	var item SpaceDataItem
//	err := json.Unmarshal([]byte(key), &item)
//	item.Key = key
//	return item, err
//}
//
//func (di SpaceDataItem) ID() string {
//	return string(di.Key)
//}
//
//func (di SpaceDataItem) Size() uint64 {
//	return 1
//}
//
//func (di SpaceDataItem) Values() []interface{} {
//	res := make([]interface{}, 2)
//	res[0] = di.Lat
//	res[1] = di.Lon
//	return res
//}
//
////for iter := uint64(0); iter < 100_000/2; iter++ {
////lat := -90 + rand.Float64()*(0 - -90)
////lon := -180 + rand.Float64()*(0 - -180)
////key := fmt.Sprintf("{Lat:%f, Lon:%f}", lat, lon)
////
////lat := rand.Float64()*(90 - 0)
////lon := rand.Float64()*(180 - 0)
////key := fmt.Sprintf("{Lat:%f, Lon:%f}", lat, lon)
//
//func geocompare() {
//	rates := make(map[string]int)
//
//	nodes := map[string]int{
//		"node-0": 1,
//		"node-1": 1,
//		"node-2": 1,
//		"node-3": 1,
//		//"node-4": 1,
//		//"node-5": 1,
//		//"node-6": 1,
//		//"node-7": 1,
//	}
//
//	fmt.Println("========= STRAIGHT =========")
//	dis := straight()
//	rates = geoMorton64(dis, nodes)
//	fmt.Println("SFC Morton 64", rates)
//	rates = geoHilbert64(dis, nodes)
//	fmt.Println("SFC Hilbert 64", rates)
//	rates = geoMorton8(dis, nodes)
//	fmt.Println("SFC Morton 8", rates)
//	rates = geoHilbert8(dis, nodes)
//	fmt.Println("SFC Hilbert 8", rates)
//	rates = geoPowerMorton64(dis, nodes)
//	fmt.Println("SFC Morton 64 Statefull", rates)
//	rates = geoPowerHilbert64(dis, nodes)
//	fmt.Println("SFC Hilbert 64 Statefull", rates)
//	rates = geoPowerMorton8(dis, nodes)
//	fmt.Println("SFC Morton 8 Statefull", rates)
//	rates = geoPowerHilbert8(dis, nodes)
//	fmt.Println("SFC Hilbert 8 Statefull", rates)
//	rates = geoConsring(dis, nodes)
//	fmt.Println("Consistent", rates)
//
//	fmt.Println("========= 2 CLUSTERS =========")
//	dis = clusters()
//	rates = geoMorton64(dis, nodes)
//	fmt.Println("SFC Morton 64", rates)
//	rates = geoHilbert64(dis, nodes)
//	fmt.Println("SFC Hilbert 64", rates)
//	rates = geoMorton8(dis, nodes)
//	fmt.Println("SFC Morton 8", rates)
//	rates = geoHilbert8(dis, nodes)
//	fmt.Println("SFC Hilbert 8", rates)
//	rates = geoPowerMorton64(dis, nodes)
//	fmt.Println("SFC Morton 64 Statefull", rates)
//	rates = geoPowerHilbert64(dis, nodes)
//	fmt.Println("SFC Hilbert 64 Statefull", rates)
//	rates = geoPowerMorton8(dis, nodes)
//	fmt.Println("SFC Morton 8 Statefull", rates)
//	rates = geoPowerHilbert8(dis, nodes)
//	fmt.Println("SFC Hilbert 8 Statefull", rates)
//	rates = geoConsring(dis, nodes)
//	fmt.Println("Consistent", rates)
//
//	fmt.Println("========= COMPLICATED =========")
//	dis = complicated()
//	rates = geoMorton64(dis, nodes)
//	fmt.Println("SFC Morton 64", rates)
//	rates = geoHilbert64(dis, nodes)
//	fmt.Println("SFC Hilbert 64", rates)
//	rates = geoMorton8(dis, nodes)
//	fmt.Println("SFC Morton 8", rates)
//	rates = geoHilbert8(dis, nodes)
//	fmt.Println("SFC Hilbert 8", rates)
//	rates = geoPowerMorton64(dis, nodes)
//	fmt.Println("SFC Morton 64 Statefull", rates)
//	rates = geoPowerHilbert64(dis, nodes)
//	fmt.Println("SFC Hilbert 64 Statefull", rates)
//	rates = geoPowerMorton8(dis, nodes)
//	fmt.Println("SFC Morton 8 Statefull", rates)
//	rates = geoPowerHilbert8(dis, nodes)
//	fmt.Println("SFC Hilbert 8 Statefull", rates)
//	rates = geoConsring(dis, nodes)
//	fmt.Println("Consistent", rates)
//}
//
//func straight() []SpaceDataItem {
//	var dis []SpaceDataItem
//	for iter := uint64(0); iter < GEO_KEYS; iter++ {
//		lat := -90 + rand.Float64()*(90 - -90)
//		lon := -180 + rand.Float64()*(180 - -180)
//		key := fmt.Sprintf("{\"Lat\":%f, \"Lon\":%f}", lat, lon)
//		di, err := NewSpaceDataItem(key)
//		if err != nil {
//			panic(err)
//		}
//		//key := fmt.Sprintf("key-%d", rand.Int())
//		dis = append(dis, di)
//	}
//	return dis
//}
//
//func clusters() []SpaceDataItem {
//	var dis []SpaceDataItem
//	for iter := uint64(0); iter < GEO_KEYS/2; iter++ {
//		lat := -90 + rand.Float64()*(0 - -90)
//		lon := -180 + rand.Float64()*(0 - -180)
//		key := fmt.Sprintf("{\"Lat\":%f, \"Lon\":%f}", lat, lon)
//		di, err := NewSpaceDataItem(key)
//		if err != nil {
//			panic(err)
//		}
//		dis = append(dis, di)
//
//		lat = rand.Float64() * (90)
//		lon = rand.Float64() * (180)
//		key = fmt.Sprintf("{\"Lat\":%f, \"Lon\":%f}", lat, lon)
//		di, err = NewSpaceDataItem(key)
//		if err != nil {
//			panic(err)
//		}
//		//key := fmt.Sprintf("key-%d", rand.Int())
//		dis = append(dis, di)
//	}
//	return dis
//}
//
//func complicated() []SpaceDataItem {
//	var dis []SpaceDataItem
//	for iter := uint64(0); iter < GEO_KEYS/2; iter++ {
//		lat := -45 + rand.Float64()*(0 - -90)
//		lon := -180 + rand.Float64()*(0 - -180)
//		key := fmt.Sprintf("{\"Lat\":%f, \"Lon\":%f}", lat, lon)
//		di, err := NewSpaceDataItem(key)
//		if err != nil {
//			panic(err)
//		}
//		dis = append(dis, di)
//
//		lat = 45 + rand.Float64()*(0 - -45)
//		lon = 0 + rand.Float64()*(0 - -180)
//		key = fmt.Sprintf("{\"Lat\":%f, \"Lon\":%f}", lat, lon)
//		di, err = NewSpaceDataItem(key)
//		if err != nil {
//			panic(err)
//		}
//		//key := fmt.Sprintf("key-%d", rand.Int())
//		dis = append(dis, di)
//	}
//	return dis
//}
//
//
//func geoMorton64(dis []SpaceDataItem, nodes map[string]int) map[string]int {
//rand.Seed(42)
//var ns []balancer.Node
//for n, w := range nodes {
//	ns = append(ns, balancer.NewMockNode(n, float64(w), 20))
//}
//
//bal, err := balancer.NewBalancer(curve.Morton, GEO_DIMS, 64, transform.SpaceTransform,
//	optimizer.RangeOptimizer, ns)
//if err != nil {
//	panic(err)
//}
//
//if err := bal.Optimize(); err != nil {
//	panic(err)
//}
//
//rates := make(map[string]int)
//for n := range nodes {
//	rates[n] = 0
//}
//for _, di := range dis {
//	if n, err := bal.LocateData(di); err != nil {
//		panic(err)
//	} else {
//		rates[n.ID()]++
//	}
//}
//return rates
//}

//func geoMorton8(dis []SpaceDataItem, nodes map[string]int) map[string]int {
//	rand.Seed(42)
//	var ns []balancer.Node
//	for n, w := range nodes {
//		ns = append(ns, balancer.NewMockNode(n, float64(w), 20))
//	}
//
//	bal, err := balancer.NewBalancer(curve.Morton, GEO_DIMS, 8, transform.SpaceTransform,
//		optimizer.RangeOptimizer, ns)
//	if err != nil {
//		panic(err)
//	}
//
//	if err := bal.Optimize(); err != nil {
//		panic(err)
//	}
//
//	rates := make(map[string]int)
//	for n := range nodes {
//		rates[n] = 0
//	}
//	for _, di := range dis {
//		if n, err := bal.LocateData(di); err != nil {
//			panic(err)
//		} else {
//			rates[n.ID()]++
//		}
//	}
//	return rates
//}
//func geoHilbert64(dis []SpaceDataItem, nodes map[string]int) map[string]int {
//	rand.Seed(42)
//	var ns []balancer.Node
//	for n, w := range nodes {
//		ns = append(ns, balancer.NewMockNode(n, float64(w), 20))
//	}
//
//	bal, err := balancer.NewBalancer(curve.Hilbert, GEO_DIMS, 64, transform.SpaceTransform,
//		optimizer.RangeOptimizer, ns)
//	if err != nil {
//		panic(err)
//	}
//
//	if err := bal.Optimize(); err != nil {
//		panic(err)
//	}
//
//	rates := make(map[string]int)
//	for n := range nodes {
//		rates[n] = 0
//	}
//	for _, di := range dis {
//		if n, err := bal.LocateData(di); err != nil {
//			panic(err)
//		} else {
//			rates[n.ID()]++
//		}
//	}
//	return rates
//}
//func geoHilbert8(dis []SpaceDataItem, nodes map[string]int) map[string]int {
//	rand.Seed(42)
//	var ns []balancer.Node
//	for n, w := range nodes {
//		ns = append(ns, balancer.NewMockNode(n, float64(w), 20))
//	}
//
//	bal, err := balancer.NewBalancer(curve.Hilbert, GEO_DIMS, 8, transform.SpaceTransform,
//		optimizer.RangeOptimizer, ns)
//	if err != nil {
//		panic(err)
//	}
//
//	if err := bal.Optimize(); err != nil {
//		panic(err)
//	}
//
//	rates := make(map[string]int)
//	for n := range nodes {
//		rates[n] = 0
//	}
//	for _, di := range dis {
//		if n, err := bal.LocateData(di); err != nil {
//			panic(err)
//		} else {
//			rates[n.ID()]++
//		}
//	}
//	return rates
//}
//func geoPowerMorton64(dis []SpaceDataItem, nodes map[string]int) map[string]int {
//	rand.Seed(42)
//	var ns []balancer.Node
//	for n, w := range nodes {
//		ns = append(ns, balancer.NewMockNode(n, float64(w), 20))
//	}
//
//	bal, err := balancer.NewBalancer(curve.Morton, GEO_DIMS, 64, transform.SpaceTransform,
//		optimizer.PowerRangeOptimizer, ns)
//	if err != nil {
//		panic(err)
//	}
//
//	if err := bal.Optimize(); err != nil {
//		panic(err)
//	}
//
//	rates := make(map[string]int)
//	for n := range nodes {
//		rates[n] = 0
//	}
//	for _, di := range dis {
//		if n, err := bal.LocateData(di); err != nil {
//			panic(err)
//		} else {
//			rates[n.ID()]++
//		}
//		if err := bal.Optimize(); err != nil {
//			panic(err)
//		}
//	}
//	return rates
//}
//func geoPowerMorton8(dis []SpaceDataItem, nodes map[string]int) map[string]int {
//	rand.Seed(42)
//	var ns []balancer.Node
//	for n, w := range nodes {
//		ns = append(ns, balancer.NewMockNode(n, float64(w), 20))
//	}
//
//	bal, err := balancer.NewBalancer(curve.Morton, GEO_DIMS, 8, transform.SpaceTransform,
//		optimizer.PowerRangeOptimizer, ns)
//	if err != nil {
//		panic(err)
//	}
//
//	if err := bal.Optimize(); err != nil {
//		panic(err)
//	}
//
//	rates := make(map[string]int)
//	for n := range nodes {
//		rates[n] = 0
//	}
//	for _, di := range dis {
//		if n, err := bal.LocateData(di); err != nil {
//			panic(err)
//		} else {
//			rates[n.ID()]++
//		}
//		if err := bal.Optimize(); err != nil {
//			panic(err)
//		}
//	}
//	return rates
//}
//func geoPowerHilbert64(dis []SpaceDataItem, nodes map[string]int) map[string]int {
//	rand.Seed(42)
//	var ns []balancer.Node
//	for n, w := range nodes {
//		ns = append(ns, balancer.NewMockNode(n, float64(w), 20))
//	}
//
//	bal, err := balancer.NewBalancer(curve.Hilbert, GEO_DIMS, 64, transform.SpaceTransform,
//		optimizer.PowerRangeOptimizer, ns)
//	if err != nil {
//		panic(err)
//	}
//
//	if err := bal.Optimize(); err != nil {
//		panic(err)
//	}
//
//	rates := make(map[string]int)
//	for n := range nodes {
//		rates[n] = 0
//	}
//	for _, di := range dis {
//		if n, err := bal.LocateData(di); err != nil {
//			panic(err)
//		} else {
//			rates[n.ID()]++
//		}
//		if err := bal.Optimize(); err != nil {
//			panic(err)
//		}
//	}
//	return rates
//}
//func geoPowerHilbert8(dis []SpaceDataItem, nodes map[string]int) map[string]int {
//	rand.Seed(42)
//	var ns []balancer.Node
//	for n, w := range nodes {
//		ns = append(ns, balancer.NewMockNode(n, float64(w), 20))
//	}
//
//	bal, err := balancer.NewBalancer(curve.Hilbert, GEO_DIMS, 8, transform.SpaceTransform,
//		optimizer.PowerRangeOptimizer, ns)
//	if err != nil {
//		panic(err)
//	}
//
//	if err := bal.Optimize(); err != nil {
//		panic(err)
//	}
//
//	rates := make(map[string]int)
//	for n := range nodes {
//		rates[n] = 0
//	}
//	for _, di := range dis {
//		if n, err := bal.LocateData(di); err != nil {
//			panic(err)
//
//		} else {
//			rates[n.ID()]++
//		}
//		if err := bal.Optimize(); err != nil {
//			panic(err)
//		}
//	}
//	return rates
//}
//func geoConsring(dis []SpaceDataItem, nodes map[string]int) map[string]int {
//	cfg := consistent.Config{
//		PartitionCount:    len(nodes),
//		ReplicationFactor: 1,
//		Load:              1.25,
//		Hasher:            hasher{},
//	}
//
//	var ns []consistent.Member
//	for n := range nodes {
//		ns = append(ns, myMember(n))
//	}
//
//	c := consistent.New(ns, cfg)
//	// Add some members to the consistent hash table.
//	// Add function calculates average load and distributes partitions over members
//
//	rates := make(map[string]int)
//	for n := range nodes {
//		rates[n] = 0
//	}
//	for _, di := range dis {
//		owner := c.LocateKey([]byte(di.ID()))
//		rates[owner.String()]++
//	}
//
//	return rates
//}
