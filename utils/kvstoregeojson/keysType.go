package main

import (
	"bytes"
	"encoding/json"
)

type Keys struct {
	Nodes map[string][]Point
}

type Point struct {
	Lon          float64
	Lat          float64
	ClusterColor string
	Cluster      string
	Key          string `json:"-"`
}

func (p *Point) UnmarshalJSON(data []byte) error {
	type cp Point
	data = bytes.ReplaceAll(data, []byte("\\"), nil)
	data = bytes.ReplaceAll(data, []byte("\"{"), []byte("{"))
	data = bytes.ReplaceAll(data, []byte("}\""), []byte("}"))
	if err := json.Unmarshal(data, (*cp)(p)); err != nil {
		return err
	}
	b, err := json.Marshal(p)
	if err != nil {
		return err
	}
	p.Key = string(b)
	return nil
}
