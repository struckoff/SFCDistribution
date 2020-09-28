package dataset

import (
	"fmt"
	balancer "github.com/struckoff/sfcframework"
)

type MatrixCell struct {
	r uint64
	c uint64
}

func (m MatrixCell) ID() string {
	return fmt.Sprintf("[%d, %d]", m.r, m.c)
}

func (m MatrixCell) Size() uint64 {
	return 1
}

func (m MatrixCell) Values() []interface{} {
	return []interface{}{m.r, m.c}
}

func Matix(r, c uint64) (res []balancer.DataItem) {
	res = make([]balancer.DataItem, 0, r*c)
	for ri := uint64(0); ri < r; ri++ {
		for ci := uint64(0); ci < c; ci++ {
			res = append(res, MatrixCell{ri, ci})
		}
	}
	return res
}
