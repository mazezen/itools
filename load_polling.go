package itools

import (
	"errors"
	"hash/crc32"
	"math/rand"
	"sort"
	"strconv"
	"sync"
)

var (
	ParamsNoEnoughError = errors.New("at least one parameter is required")
)

type RandomPolling struct {
	currentIndex int
	rss          []interface{}
}

func (p *RandomPolling) Add(params ...interface{}) error {
	if len(params) == 0 {
		return ParamsNoEnoughError
	}
	p.rss = append(p.rss, params...)
	return nil
}

func (p *RandomPolling) RandomNext() interface{} {
	if len(p.rss) == 0 {
		return ""
	}
	p.currentIndex = rand.Intn(len(p.rss))
	return p.rss[p.currentIndex]
}

type RotationPolling struct {
	currentIndex int
	rss          []interface{}
}

func (p *RotationPolling) Add(params ...interface{}) error {
	if len(params) == 0 {
		return ParamsNoEnoughError
	}
	p.rss = append(p.rss, params...)
	return nil
}

func (p *RotationPolling) RotationNext() interface{} {
	if len(p.rss) == 0 {
		return ""
	}

	var length = len(p.rss)
	if p.currentIndex >= length {
		p.currentIndex = 0
	}
	current := p.rss[p.currentIndex]
	p.currentIndex++

	return current
}

type Hash func(data []byte) uint32
type Uint32Slice []uint32

func (p Uint32Slice) Len() int           { return len(p) }
func (p Uint32Slice) Less(i, j int) bool { return p[i] < p[j] }
func (p Uint32Slice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

type ConsistentHashPolling struct {
	mux      sync.RWMutex
	hash     Hash
	replicas int
	keys     Uint32Slice
	hashMap  map[uint32]interface{}
}

func NewConsistentHashPolling(replicas int, fn Hash) *ConsistentHashPolling {
	m := &ConsistentHashPolling{
		hash:     fn,
		replicas: replicas,
		hashMap:  make(map[uint32]interface{}),
	}
	if m.hash == nil {
		m.hash = crc32.ChecksumIEEE
	}
	return m
}

func (p *ConsistentHashPolling) Add(params ...interface{}) error {
	if len(params) == 0 {
		return ParamsNoEnoughError
	}
	pa := params[0]
	p.mux.Lock()
	defer p.mux.Unlock()
	for i := 0; i < p.replicas; i++ {
		hash := p.hash([]byte(strconv.Itoa(i) + pa.(string)))
		p.keys = append(p.keys, hash)
		p.hashMap[hash] = pa
	}
	sort.Sort(p.keys)
	return nil
}

func (p *ConsistentHashPolling) Next(key string) (interface{}, error) {
	if len(p.keys) == 0 {
		return "", errors.New("is empty")
	}
	hash := p.hash([]byte(key))

	idx := sort.Search(len(p.keys), func(i int) bool {
		return p.keys[i] >= hash
	})

	if idx == len(p.keys) {
		idx = 0
	}
	p.mux.RLock()
	defer p.mux.RUnlock()
	return p.hashMap[p.keys[idx]], nil
}

type WeightPolling struct {
	curIndex int
	rss      []*WeightedNode
	rsw      []int
}

type WeightedNode struct {
	addr            string
	weight          int
	currentWeight   int
	effectiveWeight int
}

func (r *WeightPolling) Add(params ...string) error {
	if len(params) != 2 {
		return errors.New("param len need 2")
	}
	parInt, err := strconv.ParseInt(params[1], 10, 64)
	if err != nil {
		return nil
	}
	node := &WeightedNode{addr: params[0], weight: int(parInt)}
	node.effectiveWeight = node.weight
	r.rss = append(r.rss, node)
	return nil
}

func (r *WeightPolling) Next() string {
	total := 0
	var best *WeightedNode
	for i := 0; i < len(r.rss); i++ {
		w := r.rss[i]
		total += w.effectiveWeight
		w.currentWeight += w.effectiveWeight

		if w.effectiveWeight < w.weight {
			w.effectiveWeight++
		}
		if best == nil || w.currentWeight > best.currentWeight {
			best = w
		}
	}
	if best == nil {
		return ""
	}

	best.currentWeight -= total
	return best.addr
}
