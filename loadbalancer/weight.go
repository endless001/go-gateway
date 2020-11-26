package loadbalancer

import (
	"errors"
	"strconv"
)

type WeightBalance struct {
	currnetIndex int
	rss          []*WeightNode
	rsw          []int
}

type WeightNode struct {
	addr            string
	weight          int
	currnetWeight   int
	effectiveWeight int
}

func (r *WeightBalance) Add(params ...string) error {
	if len(params) != 2 {
		return errors.New("param len need 2")
	}
	parInt, err := strconv.ParseInt(params[1], 10, 64)
	if err != nil {
		return err
	}

	node := &WeightNode{addr: params[0], weight: int(parInt)}
	node.effectiveWeight = node.weight
	r.rss = append(r.rss, node)
	return nil
}

func (r *WeightBalance) Next() string {
	total := 0
	var best *WeightNode
	for i := 0; i < len(r.rss); i++ {
		w := r.rss[i]
		total += w.effectiveWeight

		w.currnetWeight += w.effectiveWeight
		if w.effectiveWeight < w.weight {
			w.effectiveWeight++
		}
		if best == nil || w.currnetWeight > best.currnetWeight {
			best = w
		}
	}
	if best == nil {
		return ""
	}
	best.currnetWeight -= total
	return best.addr
}

func (r *WeightBalance) Get(key string) (string, error) {
	return r.Next(), nil
}
