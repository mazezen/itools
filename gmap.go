package itools

import "sync"

type GMap struct {
	M map[string]any
	l sync.RWMutex
}

var Gm *GMap

func init() {
	Gm = &GMap{
		M: make(map[string]any),
	}
}

func (g *GMap) Set(key string, value interface{}) {
	g.l.Lock()
	defer g.l.Unlock()
	g.M[key] = value
}

func (g *GMap) Has(key string) bool {
	g.l.RLock()
	defer g.l.RUnlock()
	_, ok := g.M[key]
	return ok
}

func (g *GMap) Delete(key string) {
	g.l.Lock()
	defer g.l.Unlock()
	delete(g.M, key)
}

func (g *GMap) Get(key string) interface{} {
	g.l.RLock()
	defer g.l.RUnlock()
	if v, ok := g.M[key]; ok {
		return v
	}
	return nil
}
