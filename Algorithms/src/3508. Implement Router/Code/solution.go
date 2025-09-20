package main

import "fmt"

type Packet struct {
    s, d, t int
}

type Router struct {
    memoryLimit int
    queue       []Packet
    qHead       int
    size        int
    seen        map[string]bool
    destTimes   map[int][]int
    destHeads   map[int]int
}

func makeKey(s, d, t int) string {
    return fmt.Sprintf("%d#%d#%d", s, d, t)
}

func Constructor(memoryLimit int) Router {
    return Router{
        memoryLimit: memoryLimit,
        queue:       make([]Packet, 0),
        qHead:       0,
        size:        0,
        seen:        make(map[string]bool),
        destTimes:   make(map[int][]int),
        destHeads:   make(map[int]int),
    }
}

func (this *Router) AddPacket(source int, destination int, timestamp int) bool {
    key := makeKey(source, destination, timestamp)
    if this.seen[key] {
        return false
    }

    for this.size >= this.memoryLimit {
        old := this.queue[this.qHead]
        this.qHead++
        this.size--
        delete(this.seen, makeKey(old.s, old.d, old.t))
        this.destHeads[old.d]++
    }

    this.queue = append(this.queue, Packet{source, destination, timestamp})
    this.size++
    this.seen[key] = true
    this.destTimes[destination] = append(this.destTimes[destination], timestamp)
    return true
}

func (this *Router) ForwardPacket() []int {
    if this.size == 0 {
        return []int{}
    }
    pkt := this.queue[this.qHead]
    this.qHead++
    this.size--
    delete(this.seen, makeKey(pkt.s, pkt.d, pkt.t))
    this.destHeads[pkt.d]++
    return []int{pkt.s, pkt.d, pkt.t}
}

func lowerBound(arr []int, target int, lo int) int {
    l, r := lo, len(arr)
    for l < r {
        m := (l + r) / 2
        if arr[m] < target {
            l = m + 1
        } else {
            r = m
        }
    }
    return l
}

func upperBound(arr []int, target int, lo int) int {
    l, r := lo, len(arr)
    for l < r {
        m := (l + r) / 2
        if arr[m] <= target {
            l = m + 1
        } else {
            r = m
        }
    }
    return l
}

func (this *Router) GetCount(destination int, startTime int, endTime int) int {
    arr, ok := this.destTimes[destination]
    if !ok {
        return 0
    }
    h := this.destHeads[destination]
    L := lowerBound(arr, startTime, h)
    R := upperBound(arr, endTime, h)
    return R - L
}

/**
 * Your Router object will be instantiated and called as such:
 * obj := Constructor(memoryLimit);
 * param_1 := obj.AddPacket(source,destination,timestamp);
 * param_2 := obj.ForwardPacket();
 * param_3 := obj.GetCount(destination,startTime,endTime);
 */
