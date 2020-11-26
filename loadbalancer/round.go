package loadbalancer

import "errors"

type RoundBalance struct {
	currentIndex int
	ipList       []string
}

func (r *RoundBalance) Add(params ...string) error {
	if len(params) == 0 {
		return errors.New("param len 1 at least")
	}
	addr := params[0]
	r.ipList = append(r.ipList, addr)
	return nil
}

func (r *RoundBalance) Next() string {
	if len(r.ipList) == 0 {
		return ""
	}
	lens := len(r.ipList)
	if r.currentIndex > lens {
		r.currentIndex = 0
	}

	currentAddr := r.ipList[r.currentIndex]
	r.currentIndex = (r.currentIndex + 1) % lens
	return currentAddr
}

func (r *RoundBalance) Get(key string) (string, error) {
	return r.Next(), nil
}
