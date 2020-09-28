package optimizer

import (
	"github.com/pkg/errors"
	balancer "github.com/struckoff/sfcframework"
	"log"
	"math"
	"sort"
)

func CapPowerOptimizer(s *balancer.Space) (res []*balancer.CellGroup, err error) {
	defer func() {
		cgs := s.CellGroups()
		for _, cg := range cgs {
			log.Println("CapPowerOptimizer", cg.Node().ID(), cg.Node().Hash(), ": ", cg.Range().Min, cg.Range().Max, cg.Range().Len, cg.TotalLoad(), len(cg.Cells()))
		}
	}()

	// Distribute cells into segments based on power
	cgs, err := powerRangeOptimizer(s)
	if err != nil {
		return nil, err
	}

	// Sort by hash to make nodes connected to curve
	sort.Slice(cgs, func(i, j int) bool { return cgs[i].Node().Hash() < cgs[j].Node().Hash() })

	if len(cgs) >= 2 {
		for cgIdx := 0; cgIdx < len(cgs)-1; cgIdx++ {
			// transfer nodes from one cell group to another if load of group exceeds capacity
			if err := equalizer(cgs[cgIdx], cgs[cgIdx+1], s); err != nil {
				return nil, err
			}
		}
	}

	//check cell groups in case if there capacity overload
	for cgIdx := range cgs {
		c, err := cgs[cgIdx].Node().Capacity().Get()
		if err != nil {
			return nil, err
		}
		if float64(cgs[cgIdx].TotalLoad()) > c {
			return nil, errors.New("out of capacity")
		}
	}

	return cgs, nil
}

func powerRangeOptimizer(s *balancer.Space) (res []*balancer.CellGroup, err error) {
	totalPower := s.TotalPower()
	cgs := s.CellGroups()
	if len(cgs) == 0 {
		return res, nil
	}
	var max, min uint64

	sort.Slice(cgs, func(i, j int) bool { return cgs[i].Node().Hash() < cgs[j].Node().Hash() })

	for iter := 0; iter < len(cgs); iter++ {
		min = max
		p := cgs[iter].Node().Power().Get() / totalPower
		max = min + uint64(math.Ceil(float64(s.Capacity())*p))
		if max > s.Capacity()+1 {
			max = s.Capacity() + 1
		}
		if err := cgs[iter].SetRange(min, max, s); err != nil {
			return nil, errors.Wrap(err, "range power optimizer error")
		}
	}
	if max < s.Capacity() {
		if err := cgs[len(cgs)-1].SetRange(min, s.Capacity()+1, s); err != nil {
			return nil, errors.Wrap(err, "range power optimizer error")
		}
	}
	return cgs, nil
}

func equalizer(lcg, rcg *balancer.CellGroup, s *balancer.Space) error {
	if lcg.Range().Max != rcg.Range().Min {
		return errors.New("wrong group pair")
	}

	lc, err := lcg.Node().Capacity().Get()
	if err != nil {
		return err
	}
	//transfer cells from left cell group to right
	nbf := true
	for float64(lcg.TotalLoad()) > lc {
		nbf = false
		if err := lcg.SetRange(lcg.Range().Min, lcg.Range().Max-1, s); err != nil {
			return err
		}
		if err := rcg.SetRange(rcg.Range().Min-1, rcg.Range().Max, s); err != nil {
			return err
		}
	}

	// transfer cells from right cell group to left if no cell were transferred from left
	// and if left cell group has capacity to handle cells
	if nbf {
		rc, err := rcg.Node().Capacity().Get()
		if err != nil {
			return err
		}
		for float64(rcg.TotalLoad()) > rc && float64(lcg.TotalLoad()) <= lc {
			if err := lcg.SetRange(lcg.Range().Min, lcg.Range().Max+1, s); err != nil {
				return err
			}
			if err := rcg.SetRange(rcg.Range().Min+1, rcg.Range().Max, s); err != nil {
				return err
			}
		}
	}

	return nil
}

//func StaticCapacityStateful(s *balancer.Space) (res []*balancer.CellGroup, err error) {
//	defer func() {
//		cgs := s.CellGroups()
//		for _, cg := range cgs {
//			log.Println("StaticCapacityStateful", cg.Node().ID(), ": ", cg.Range().Min, cg.Range().Max, cg.Range().Len, cg.TotalLoad(), len(cg.Cells()))
//		}
//	}()
//	//TODO: reduce Capacity calls
//
//	//caps := make([]float64, len(cgs))
//	//for iter := range cgs {
//	//	caps[iter] = float64(cgs[iter].TotalLoad()) / float64(s.TotalLoad())
//	//}
//
//	cells := s.Cells()
//	totalPower := s.TotalPower()
//	//totalFree := s.TotalCapacity() - float64(s.TotalLoad())
//	cgs := s.CellGroups()
//	if len(cgs) == 0 {
//		return res, nil
//	}
//	var max, min uint64
//
//	var tc float64
//
//	//tcs := make([]float64, len(cgs))
//	for iter := range cgs {
//		c, err := cgs[iter].Node().Capacity().Get()
//		if err != nil {
//			return nil, err
//		}
//		tc += c - float64(cgs[iter].TotalLoad())
//	}
//
//	//sort.Slice(cgs, func(i, j int) bool {
//	//	capI, _ := cgs[i].Node().Capacity().Get()
//	//	capJ, _ := cgs[j].Node().Capacity().Get()
//	//	return (capI - float64(cgs[i].TotalLoad())) < (capJ - float64(cgs[j].TotalLoad()))
//	//})
//
//	for iter := 0; iter < len(cgs); iter++ {
//		min = max
//		p := cgs[iter].Node().Power().Get() / totalPower
//		f, err := cgs[iter].Node().Capacity().Get()
//		if err != nil {
//			return nil, err
//		}
//		c := (f - float64(cgs[iter].TotalLoad())) / tc
//		max = min + uint64(math.Round(float64(s.Capacity())*p*c))
//
//		for citer := 0; citer < len(cells); citer++ {
//			if cells[citer].ID() > max {
//				break
//			}
//			if cells[citer].ID() >= min {
//				cgs[iter].AddCell(cells[citer], true)
//			}
//		}
//		if err := cgs[iter].SetRange(min, max, nil); err != nil {
//			return nil, errors.Wrap(err, "power range optimizer error")
//		}
//	}
//
//	if max < s.Capacity() {
//		if err := cgs[len(cgs)-1].SetRange(min, s.Capacity()+1, nil); err != nil {
//			return nil, errors.Wrap(err, "range optimizer error")
//		}
//		for citer := range cells {
//			if cells[citer].ID() >= max {
//				cgs[len(cgs)-1].AddCell(cells[citer], true)
//			}
//		}
//	}
//
//	//for _, cg := range cgs {
//	//	log.Print(cg.ID(), ":", cg.TotalLoad(), cg.Range().Min, cg.Range().Max, cg.Range().Len, len(cells))
//	//}
//
//	return cgs, nil
//}
//
//func DynamicCapacityStateful(s *balancer.Space) (res []*balancer.CellGroup, err error) {
//	//TODO: reduce Capacity calls
//
//	cells := s.Cells()
//	totalPower := s.TotalPower()
//	//totalFree := s.TotalCapacity() - float64(s.TotalLoad())
//	cgs := s.CellGroups()
//	if len(cgs) == 0 {
//		return res, nil
//	}
//	var max, min uint64
//
//	caps := make([]float64, len(cgs))
//	for iter := range cgs {
//		caps[iter], err = cgs[iter].Node().Capacity().Get()
//		if err != nil {
//			return nil, err
//		}
//	}
//
//	sort.Slice(cgs, func(i, j int) bool {
//		capI, err := cgs[i].Node().Capacity().Get()
//		if err != nil {
//			return false
//		}
//		capJ, err := cgs[j].Node().Capacity().Get()
//		if err != nil {
//			return false
//		}
//		return capI < capJ
//	})
//
//	for iter := 0; iter < len(cgs); iter++ {
//		min = max
//		p := cgs[iter].Node().Power().Get() / totalPower
//		f, err := cgs[iter].Node().Capacity().Get()
//		if err != nil {
//			return nil, err
//		}
//		max = min + uint64(math.Round(float64(s.Capacity())*p))
//
//		for citer := 0; citer < len(cells); citer++ {
//			if cells[citer].ID() > max {
//				break
//			}
//			if cells[citer].ID() >= min {
//				f -= float64(cells[citer].Load())
//				if f <= 0 {
//					c := citer - 1
//					if c < 0 {
//						c = 0
//					}
//					max = cells[citer].ID()
//					break
//				}
//				cgs[iter].AddCell(cells[citer], true)
//			}
//		}
//		if err := cgs[iter].SetRange(min, max, nil); err != nil {
//			return nil, errors.Wrap(err, "power range optimizer error")
//		}
//	}
//
//	if max < s.Capacity() {
//		if err := cgs[len(cgs)-1].SetRange(min, s.Capacity()+1, nil); err != nil {
//			return nil, errors.Wrap(err, "range optimizer error")
//		}
//		for citer := range cells {
//			if cells[citer].ID() >= max {
//				cgs[len(cgs)-1].AddCell(cells[citer], true)
//			}
//		}
//	}
//
//	return cgs, nil
//}
//
//func LoadPowerOptimizer(s *balancer.Space) (res []*balancer.CellGroup, err error) {
//	defer func() {
//		cgs := s.CellGroups()
//		for _, cg := range cgs {
//			log.Println("LoadPowerOptimizer", cg.Node().ID(), cg.Node().Hash(), ": ", cg.Range().Min, cg.Range().Max, cg.Range().Len, cg.TotalLoad(), len(cg.Cells()))
//		}
//	}()
//
//	cells := s.Cells()
//	totalLoad := s.TotalLoad()
//	if len(cells) == 0 || totalLoad == 0 {
//		return powerRangeOptimizer(s)
//	}
//	totalPower := s.TotalPower()
//	cgs := s.CellGroups()
//	if len(cgs) == 0 {
//		return res, nil
//	}
//
//	// sort.Slice(cgs, func(i, j int) bool { return cgs[i].Node().Power().Get() > cgs[j].Node().Power().Get() })
//	sort.Slice(cgs, func(i, j int) bool { return cgs[i].Node().Hash() < cgs[j].Node().Hash() })
//	sort.Slice(cells, func(i, j int) bool { return cells[i].ID() < cells[j].ID() })
//
//	cgIdx := 0
//	cl := uint64(0)
//	p := cgs[cgIdx].Node().Power().Get() / totalPower
//	l := uint64(float64(totalLoad) * p)
//	var max, min uint64
//	c := true
//	for i := range cells {
//		id := cells[i].ID()
//		cl += cells[i].Load()
//		cp, err := cgs[cgIdx].Node().Capacity().Get()
//		if err != nil {
//			return nil, err
//		}
//		if cl > uint64(cp) {
//			min = max
//			max = id
//			if err := cgs[cgIdx].SetRange(min, max, nil); err != nil {
//				return nil, errors.Wrap(err, "range optimizer error")
//			}
//			cgIdx++
//			if cgIdx > (len(cgs)-1) && i < (len(cells)-1) {
//				return nil, errors.New("load is larger than cumulative capacity of all nodes")
//			}
//		}
//		if c && cl > l {
//			min = max
//			max = id
//			if err := cgs[cgIdx].SetRange(min, max, nil); err != nil {
//				return nil, errors.Wrap(err, "range optimizer error")
//			}
//			cgIdx++
//			cgs[cgIdx].AddCell(cells[i], true)
//			cl = cells[i].Load()
//			p = cgs[cgIdx].Node().Power().Get() / totalPower
//			l = uint64(float64(totalLoad) * p)
//			if cgIdx == (len(cgs) - 1) {
//				if err := cgs[cgIdx].SetRange(max, s.Capacity()+1, nil); err != nil {
//					return nil, errors.Wrap(err, "range optimizer error")
//				}
//				c = false
//			}
//		} else {
//			cgs[cgIdx].AddCell(cells[i], true)
//		}
//	}
//
//	return cgs, nil
//}
//
//
//func CapacityOptimizer(s *balancer.Space) (res []*balancer.CellGroup, err error) {
//	defer func() {
//		cgs := s.CellGroups()
//		for _, cg := range cgs {
//			log.Println("CapacityOptimizer", cg.Node().ID(), ": ", cg.Range().Min, cg.Range().Max, cg.Range().Len, cg.TotalLoad(), len(cg.Cells()))
//		}
//	}()
//	cells := s.Cells()
//	cgs := s.CellGroups()
//	if len(cgs) == 0 {
//		return res, nil
//	}
//
//	//sort.Slice(cgs, func(i, j int) bool {
//	//	ci, _ := cgs[i].Node().Capacity().Get()
//	//	cj, _ := cgs[j].Node().Capacity().Get()
//	//	return ci > cj
//	//})
//
//	sort.Slice(cgs, func(i, j int) bool { return cgs[i].Node().Hash() < cgs[j].Node().Hash() })
//
//	cgIdx := 0
//	cl := uint64(0)
//	c, err := cgs[cgIdx].Node().Capacity().Get()
//	if err != nil {
//		return nil, err
//	}
//	l := uint64(c)
//	var max, min uint64
//	//cll := uint64(0)
//	//for _, cell := range cells {
//	//	//cll += cell.Load()
//	//	log.Println(cell.ID(), cell.Load())
//	//}
//	//log.Println(cll)
//	//for _, cg := range cgs {
//	//	log.Println(cg.Node().Capacity().Get())
//	//}
//	for i := range cells {
//		id := cells[i].ID()
//		cl += cells[i].Load()
//		max = id
//		if cl > l {
//			min = max
//			if err := cgs[cgIdx].SetRange(min, max, nil); err != nil {
//				return nil, errors.Wrap(err, "capacity optimizer error")
//			}
//			cgIdx++
//			if cgIdx > (len(cgs)-1) && i < (len(cells)-1) {
//				return nil, errors.New("load is larger than cumulative capacity of all nodes")
//			}
//			if cgIdx >= len(cgs) {
//				break
//			}
//			cgs[cgIdx].AddCell(cells[i], true)
//			log.Println(cl, cells[i].Load())
//			cl = cells[i].Load()
//			c, err := cgs[cgIdx].Node().Capacity().Get()
//			if err != nil {
//				return nil, err
//			}
//			l = uint64(c)
//		} else {
//			cgs[cgIdx].AddCell(cells[i], true)
//		}
//	}
//	if cgIdx >= len(cgs) {
//		cgIdx = len(cgs) - 1
//	}
//	if max < s.Capacity() {
//		if err := cgs[cgIdx].SetRange(min, s.Capacity()+1, nil); err != nil {
//			return nil, errors.Wrap(err, "capacity optimizer error")
//		}
//		for i := cgIdx + 1; i < len(cgs); i++ {
//			if err := cgs[i].SetRange(math.MaxUint64, math.MaxUint64, nil); err != nil {
//				return nil, errors.Wrap(err, "capacity optimizer error")
//			}
//		}
//	}
//	return cgs, nil
//}
