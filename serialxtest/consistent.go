package serialxtest

import (
	"github.com/cespare/xxhash"
	"github.com/serialx/hashring"
	"github.com/struckoff/sfcdistribution/report"
	balancer "github.com/struckoff/sfcframework"
)

func ConsistentReport(dis []balancer.DataItem, nodes []balancer.Node, rpt report.Report) report.ReportFunc {

	nodeMap := make(map[string]balancer.Node)
	weights := make(map[string]int)
	for _, n := range nodes {
		nodeMap[n.ID()] = n
		weights[n.ID()] = int(n.Power().Get())
	}
	ring, err := hashring.NewWithHashAndWeights(weights, xxhash.New())
	if err != nil {
		panic(err)
	}

	return func(rCh chan<- report.Report) {

		//var k hashring.HashKey
		for _, di := range dis {
			//if k != ring.GenKey(di.ID()) {
			//	k = ring.GenKey(di.ID())
			//	fmt.Println(k, di.ID())
			//}
			name, _ := ring.GetNode(di.ID())
			rpt.Add(nodeMap[name], di)
		}
		rCh <- rpt
	}
}
