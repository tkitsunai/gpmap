package gpmap

import (
	"fmt"
	"sync"
)

type MapperFunc func(v PMapData) PMapData

type PMapData interface {
	Map() MapperFunc
}

type PmapContext struct {
	size   int
	target []PMapData
}

func New(list []PMapData) (*PmapContext, error) {
	if list == nil || (len(list) == 0) {
		return nil, fmt.Errorf("invalid argument")
	}

	return &PmapContext{
		target: list,
		size:   len(list),
	}, nil
}

func (p *PmapContext) Map(mapper MapperFunc) ([]PMapData, error) {
	wait := &sync.WaitGroup{}
	result := make([]PMapData, p.size, p.size)
	for i := 0; i < p.size; i++ {
		wait.Add(1)
		go func(idx int, item PMapData) {
			defer wait.Done()
			result[idx] = mapper(item)
		}(i, p.target[i])
	}
	wait.Wait()
	return result, nil
}

// for benchmark testing
func (p *PmapContext) SyncMap(mapper MapperFunc) ([]PMapData, error) {
	result := make([]PMapData, p.size, p.size)
	for i := 0; i < p.size; i++ {
		result[i] = mapper(p.target[i])
	}
	return result, nil
}
