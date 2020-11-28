package loadbalancer

import (
	"errors"
	"math/rand"
)

type RandomBalance struct {
	currentIndex int
	ipList       []string
}

func (r *RandomBalance) Add(params ...string) error {
	if len(params) == 0 {
		return errors.New("param len 1 at least")
	}
	addr := params[0]

	r.ipList = append(r.ipList, addr)
	return nil
}

func (r *RandomBalance) Next() string {
	if len(r.ipList) == 0 {
		return ""
	}
	r.currentIndex = rand.Intn(len(r.ipList))
	return r.ipList[r.currentIndex]
}

func (r *RandomBalance) Get(key string) (string, error) {
	return r.Next(), nil
}
