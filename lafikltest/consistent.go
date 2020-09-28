package lafikltest

import (
	"github.com/lafikl/consistent"
	"github.com/struckoff/sfcdistribution/report"
	balancer "github.com/struckoff/sfcframework"
)

func ConsistentReport(dis []balancer.DataItem, nodes []balancer.Node, rpt report.Report) report.ReportFunc {
	ring := consistent.New()

	nodeMap := make(map[string]balancer.Node)
	for _, n := range nodes {
		nodeMap[n.ID()] = n
		ring.Add(n.ID())
	}

	return func(rCh chan<- report.Report) {
		//rpt := report.NewStatReport("lafikl/consistent", nodes)
		//var k hashring.HashKey
		for _, di := range dis {
			//if k != ring.GenKey(di.ID()) {
			//	k = ring.GenKey(di.ID())
			//	fmt.Println(k, di.ID())
			//}
			name, err := ring.Get(di.ID())
			if err != nil {
				panic(err)
			}
			if err := rpt.Add(nodeMap[name], di); err != nil {
				panic(err)
			}
		}
		rCh <- rpt
	}
}
func ConsistentReportWithLoad(dis []balancer.DataItem, nodes []balancer.Node, rpt report.Report) report.ReportFunc {
	ring := consistent.New()

	nodeMap := make(map[string]balancer.Node)
	for _, n := range nodes {
		nodeMap[n.ID()] = n
		ring.Add(n.ID())
	}

	return func(rCh chan<- report.Report) {

		for _, di := range dis {
			name, err := ring.Get(di.ID())
			if err != nil {
				panic(err)
			}
			ring.Inc(name)
			rpt.Add(nodeMap[name], di)
		}
		rCh <- rpt
	}
}
func ConsistentLeastReport(dis []balancer.DataItem, nodes []balancer.Node, rpt report.Report) report.ReportFunc {
	ring := consistent.New()

	nodeMap := make(map[string]balancer.Node)
	for _, n := range nodes {
		nodeMap[n.ID()] = n
		ring.Add(n.ID())
	}

	return func(rCh chan<- report.Report) {
		//rpt := report.NewStatReport("lafikl/consistent+least", nodes)
		//var k hashring.HashKey
		for _, di := range dis {
			name, err := ring.GetLeast(di.ID())
			if err != nil {
				panic(err)
			}
			rpt.Add(nodeMap[name], di)
		}
		rCh <- rpt
	}
}
func ConsistentLeastReportWithLoad(dis []balancer.DataItem, nodes []balancer.Node, rpt report.Report) report.ReportFunc {
	ring := consistent.New()

	nodeMap := make(map[string]balancer.Node)
	for _, n := range nodes {
		nodeMap[n.ID()] = n
		ring.Add(n.ID())
	}

	return func(rCh chan<- report.Report) {
		//rpt := report.NewStatReport("lafikl/consistent+least+load", nodes)
		//var k hashring.HashKey
		for _, di := range dis {
			name, err := ring.GetLeast(di.ID())
			if err != nil {
				panic(err)
			}
			ring.Inc(name)
			rpt.Add(nodeMap[name], di)
		}

		rCh <- rpt
	}
}
func ConsistentLeastReportWithLoadAndDone(dis []balancer.DataItem, nodes []balancer.Node, rpt report.Report) report.ReportFunc {
	ring := consistent.New()

	nodeMap := make(map[string]balancer.Node)
	for _, n := range nodes {
		nodeMap[n.ID()] = n
		ring.Add(n.ID())
	}

	return func(rCh chan<- report.Report) {
		//rpt := report.NewStatReport("lafikl/consistent+least+load+done", nodes)
		//var k hashring.HashKey
		for _, di := range dis {
			name, err := ring.GetLeast(di.ID())
			if err != nil {
				panic(err)
			}
			ring.Inc(name)
			ring.Done(name)
			rpt.Add(nodeMap[name], di)
		}

		rCh <- rpt
	}
}
