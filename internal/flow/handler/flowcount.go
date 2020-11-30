package handler

import (
	"go-gateway/internal/flow/service"
	"sync"
	"time"
)

var FlowCounterHandler *FlowCounter

type FlowCounter struct {
	RedisFlowCountMap   map[string]*service.RedisFlowCountService
	RedisFlowCountSlice []*service.RedisFlowCountService
	Locker              sync.RWMutex
}

func init() {
	FlowCounterHandler = NewFlowCounter()
}

func NewFlowCounter() *FlowCounter {
	return &FlowCounter{
		RedisFlowCountMap:   map[string]*service.RedisFlowCountService{},
		RedisFlowCountSlice: []*service.RedisFlowCountService{},
		Locker:              sync.RWMutex{},
	}
}

func (counter *FlowCounter) GetCounter(serverName string) (*service.RedisFlowCountService, error) {
	for _, item := range counter.RedisFlowCountSlice {
		if item.AppID == serverName {
			return item, nil
		}
	}
	newCounter := service.NewRedisFlowCountService(serverName, 1*time.Second)
	counter.RedisFlowCountSlice = append(counter.RedisFlowCountSlice, newCounter)
	counter.Locker.Lock()
	defer counter.Locker.Unlock()
	counter.RedisFlowCountMap[serverName] = newCounter
	return newCounter, nil
}
