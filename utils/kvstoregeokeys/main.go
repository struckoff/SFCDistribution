package main

import (
	"encoding/json"
	"flag"
	"math/rand"
	"net/http"
	"os"
	"sort"
	"strings"
)

func main() {
	srv := flag.String("address", "http://localhost:9190", "Address of kvstore node")
	amount := flag.Int("n", 100, "amount of key sequences to produce")
	ksLen := flag.Int("l", 100, "amount of keys in key sequence")
	flag.Parse()

	ks, err := download(*srv + "/list")
	if err != nil {
		panic(err)
	}
	points := unpack(ks)
	sort.Slice(points, func(i, j int) bool { return compare(points[j], points[i]) })
	for iter := 0; iter < *amount; iter++ {
		k := keyseq(points, *ksLen)
		if _, err := os.Stdout.WriteString(k); err != nil {
			panic(err)
		}
		if _, err := os.Stdout.WriteString("\n"); err != nil {
			panic(err)
		}
	}
}

func download(url string) (Keys, error) {
	var ks Keys
	r, err := http.Get(url)
	if err != nil {
		return ks, err
	}
	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(&ks.Nodes); err != nil {
		return ks, err
	}
	return ks, nil
}

func unpack(ks Keys) []Point {
	var count int
	for _, node := range ks.Nodes {
		count += len(node)
	}
	points := make([]Point, count)
	shift := 0
	for _, node := range ks.Nodes {
		copy(points[shift:], node)
		shift += len(node)
	}
	return points
}

func compare(p0, p1 Point) bool {
	if p0.Lon > p1.Lon {
		return true
	}
	if p0.Lat > p1.Lat {
		return true
	}
	return false
}

func keyseq(points []Point, ln int) string {
	var buf strings.Builder
	shift := rand.Intn(len(points) - ln)
	for iter := shift; iter < shift+ln; iter++ {
		buf.WriteString("/")
		buf.WriteString(points[iter].Key)
	}
	//strings.Join(points[shift:shift+ln], "/")
	return buf.String()
}
