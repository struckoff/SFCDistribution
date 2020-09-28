package report

import (
	"errors"
	"fmt"
	balancer "github.com/struckoff/sfcframework"
	"sort"
	"strings"
	"sync"
)

func NewStatReport(name string, nodes []balancer.Node) *StatReport {
	return &StatReport{
		s:     make(map[balancer.Node]float64),
		name:  name,
		nodes: nodes,
	}
}

type StatReportItems []StatReportItem

func (ris *StatReportItems) Sort() *StatReportItems {
	sort.Slice((*ris), func(i, j int) bool {
		return strings.Compare((*ris)[i].Node.ID(), (*ris)[j].Node.ID()) < 0
	})
	return ris
}

func (ris StatReportItems) Print(mode, groupBy string) string {
	switch mode {
	case "tsv":
		return ris.csv(groupBy, "\t")
	case "csv":
		return ris.csv(groupBy, ",")
	}
	return ""
}

func (ris StatReportItems) csv(groupBy, sep string) string {
	headerForm := "node%[1]scount%[1]srate to all%[1]srate to node\n"
	if len(groupBy) > 0 {
		data := ris.groupByNode(groupBy)
		return ris.csvline(data, sep) + "\n"
	}
	lines := ris.csvtable(sep)
	var buf strings.Builder
	buf.WriteString(fmt.Sprintf(headerForm, sep))
	for _, line := range lines {
		buf.WriteString(line)
		buf.WriteString("\n")
	}
	return buf.String()
}

func (ris StatReportItems) csvtable(sep string) []string {
	form := "%s" + sep + "%.f" + sep + "%f" + sep + "%f"
	res := make([]string, len(ris))
	for iter, item := range ris {
		res[iter] = fmt.Sprintf(form, item.Node.ID(), item.Count, item.RateToAll, item.RateToNode)
	}
	return res
}

func (ris StatReportItems) csvline(data map[string]float64, sep string) string {
	var buf strings.Builder
	for _, item := range ris {
		val := data[item.Node.ID()]
		buf.WriteString(fmt.Sprintf("%f%s", val, sep))
	}
	return strings.TrimLeft(buf.String(), sep)
}

func (ris StatReportItems) groupByNode(field string) map[string]float64 {
	res := make(map[string]float64)
	for _, item := range ris {
		switch strings.ToLower(field) {
		case "count":
			res[item.Node.ID()] = item.Count
		case "ratetonode":
			res[item.Node.ID()] = item.RateToNode
		case "ratetoall":
			res[item.Node.ID()] = item.RateToAll
		}
	}
	return res
}

type StatReportItem struct {
	Node       balancer.Node
	RateToAll  float64
	RateToNode float64
	Count      float64
}

type StatReport struct {
	mu sync.RWMutex
	s  map[balancer.Node]float64

	name  string
	nodes []balancer.Node
	count float64
}

func (r *StatReport) Add(node balancer.Node, _ balancer.DataItem) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.count++
	r.s[node]++
	return nil
}

func (r *StatReport) Print(opts ...string) (string, error) {
	if len(opts) < 2 {
		panic(errors.New("not enough arguments"))
	}
	mode, groupBy := opts[0], opts[1]
	if len(groupBy) == 0 {
		return r.Build().Sort().Print(mode, groupBy), nil
	}
	var sep string
	switch mode {
	case "tsv":
		sep = "\t"
	case "csv":
		sep = ","
	}
	return r.name + sep + r.Build().Sort().Print(mode, groupBy), nil
}

func (r *StatReport) Build() *StatReportItems {
	r.mu.Lock()
	defer r.mu.Unlock()
	res := make(StatReportItems, 0, len(r.nodes))
	for _, node := range r.nodes {
		if c, ok := r.s[node]; ok {
			res = append(res, StatReportItem{
				Node:       node,
				RateToAll:  c / r.count,
				RateToNode: (c / r.count) * float64(len(r.nodes)),
				Count:      c,
			})
			continue
		}
		res = append(res, StatReportItem{
			Node:       node,
			RateToAll:  0,
			RateToNode: 0,
			Count:      0,
		})
	}
	return &res
}
