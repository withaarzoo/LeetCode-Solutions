package main

import (
	"container/heap"
)

// helper key for maps: pack shop and movie into int64
func key(shop int, movie int) int64 {
	return (int64(shop) << 32) | int64(uint32(movie))
}

// Available heap entry: (price, shop, version)
type AvailEntry struct{ price, shop, ver int }
type AvailHeap []AvailEntry
func (h AvailHeap) Len() int { return len(h) }
func (h AvailHeap) Less(i, j int) bool {
	if h[i].price != h[j].price { return h[i].price < h[j].price }
	return h[i].shop < h[j].shop
}
func (h AvailHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h *AvailHeap) Push(x interface{}) { *h = append(*h, x.(AvailEntry)) }
func (h *AvailHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

// Rented heap entry: (price, shop, movie, version)
type RentEntry struct{ price, shop, movie, ver int }
type RentHeap []RentEntry
func (h RentHeap) Len() int { return len(h) }
func (h RentHeap) Less(i, j int) bool {
	if h[i].price != h[j].price { return h[i].price < h[j].price }
	if h[i].shop != h[j].shop { return h[i].shop < h[j].shop }
	return h[i].movie < h[j].movie
}
func (h RentHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h *RentHeap) Push(x interface{}) { *h = append(*h, x.(RentEntry)) }
func (h *RentHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

type MovieRentingSystem struct {
	price       map[int64]int
	version     map[int64]int
	rentedState map[int64]bool

	avail map[int]*AvailHeap // movie -> heap
	rent  *RentHeap
}

func Constructor(n int, entries [][]int) MovieRentingSystem {
	m := MovieRentingSystem{
		price:       make(map[int64]int),
		version:     make(map[int64]int),
		rentedState: make(map[int64]bool),
		avail:       make(map[int]*AvailHeap),
		rent:        &RentHeap{},
	}
	heap.Init(m.rent)
	for _, e := range entries {
		shop, movie, p := e[0], e[1], e[2]
		k := key(shop, movie)
		m.price[k] = p
		m.version[k] = 0
		m.rentedState[k] = false
		if _, ok := m.avail[movie]; !ok {
			h := &AvailHeap{}
			heap.Init(h)
			m.avail[movie] = h
		}
		heap.Push(m.avail[movie], AvailEntry{p, shop, 0})
	}
	return m
}

func (this *MovieRentingSystem) Search(movie int) []int {
	res := []int{}
	h, ok := this.avail[movie]
	if !ok { return res }
	tmp := []AvailEntry{}
	for len(res) < 5 && h.Len() > 0 {
		top := heap.Pop(h).(AvailEntry)
		k := key(top.shop, movie)
		if this.version[k] == top.ver && !this.rentedState[k] {
			res = append(res, top.shop)
			tmp = append(tmp, top)
		}
	}
	for _, v := range tmp { heap.Push(h, v) }
	return res
}

func (this *MovieRentingSystem) Rent(shop int, movie int) {
	k := key(shop, movie)
	this.version[k] = this.version[k] + 1
	this.rentedState[k] = true
	p := this.price[k]
	heap.Push(this.rent, RentEntry{p, shop, movie, this.version[k]})
}

func (this *MovieRentingSystem) Drop(shop int, movie int) {
	k := key(shop, movie)
	this.version[k] = this.version[k] + 1
	this.rentedState[k] = false
	p := this.price[k]
	if _, ok := this.avail[movie]; !ok {
		h := &AvailHeap{}
		heap.Init(h)
		this.avail[movie] = h
	}
	heap.Push(this.avail[movie], AvailEntry{p, shop, this.version[k]})
}

func (this *MovieRentingSystem) Report() [][]int {
	res := [][]int{}
	tmp := []RentEntry{}
	for len(res) < 5 && this.rent.Len() > 0 {
		top := heap.Pop(this.rent).(RentEntry)
		k := key(top.shop, top.movie)
		if this.version[k] == top.ver && this.rentedState[k] {
			res = append(res, []int{top.shop, top.movie})
			tmp = append(tmp, top)
		}
	}
	for _, v := range tmp { heap.Push(this.rent, v) }
	return res
}

/**
 * Usage:
 * obj := Constructor(n, entries);
 * param_1 := obj.Search(movie);
 * obj.Rent(shop,movie);
 * obj.Drop(shop,movie);
 * param_4 := obj.Report();
 */
