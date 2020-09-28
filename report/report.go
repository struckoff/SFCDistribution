package report

import balancer "github.com/struckoff/sfcframework"

type Report interface {
	Add(balancer.Node, balancer.DataItem) error
	Print(...string) (string, error)
}
