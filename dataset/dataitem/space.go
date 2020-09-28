package dataitem

import (
	"encoding/json"
	"fmt"
)

type Space struct {
	key string `json:"-"`
	Lat float64
	Lon float64
}

func NewSpace(lat, lon float64) (Space, error) {
	var item Space
	item.Lat = lat
	item.Lon = lon

	b, err := json.Marshal(item)
	if err != nil {
		return Space{}, err
	}

	item.key = string(b)
	return item, err
}

func (di Space) ID() string {
	return fmt.Sprintf("%f,%f", di.Lat, di.Lon)
	//return di.key
}

func (di Space) Size() uint64 {
	return 1
}

func (di Space) Values() []interface{} {
	res := make([]interface{}, 2)
	res[0] = di.Lat
	res[1] = di.Lon
	return res
}
