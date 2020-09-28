package node

type Power float64

func (p Power) Get() float64 {
	return float64(p)
}
