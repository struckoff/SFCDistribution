package buraksezertest

import (
	"github.com/buraksezer/consistent"
	"github.com/cespare/xxhash"
	"github.com/struckoff/sfcdistribution/report"
	balancer "github.com/struckoff/sfcframework"
)

type Node string

func (n Node) String() string {
	return string(n)
}

// consistent package doesn't provide a default hashing function.
// You should provide a proper one to distribute keys/members uniformly.
type hasher struct{}

func (h hasher) Sum64(data []byte) uint64 {
	// you should use a proper hash function for uniformity.
	return xxhash.Sum64(data)
}

func ConsistentReport(dis []balancer.DataItem, nodes []balancer.Node, load float64, rpt report.Report) report.ReportFunc {
	cfg := consistent.Config{
		PartitionCount:    len(nodes),
		ReplicationFactor: 1,
		Load:              load,
		Hasher:            hasher{},
	}

	nodeMap := make(map[string]balancer.Node)
	var ns []consistent.Member
	for _, n := range nodes {
		nodeMap[n.ID()] = n
		ns = append(ns, Node(n.ID()))
	}
	c := consistent.New(ns, cfg)

	return func(rCh chan<- report.Report) {
		for _, di := range dis {
			n := c.LocateKey([]byte(di.ID()))
			rpt.Add(nodeMap[n.String()], di)
		}
		rCh <- rpt
	}
}
