package node

type Capacity float64

func (c Capacity) Get() (float64, error) {
	return float64(c), nil
}
