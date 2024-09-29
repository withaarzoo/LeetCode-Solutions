type AllOne struct {
    count map[string]int
    set   map[int]map[string]bool
}

func Constructor() AllOne {
    return AllOne{
        count: make(map[string]int),
        set:   make(map[int]map[string]bool),
    }
}

func (this *AllOne) Inc(key string) {
    n := this.count[key]
    this.count[key]++
    if this.set[n] != nil {
        delete(this.set[n], key)
        if len(this.set[n]) == 0 {
            delete(this.set, n)
        }
    }
    if this.set[n+1] == nil {
        this.set[n+1] = make(map[string]bool)
    }
    this.set[n+1][key] = true
}

func (this *AllOne) Dec(key string) {
    n := this.count[key]
    this.count[key]--
    delete(this.set[n], key)
    if len(this.set[n]) == 0 {
        delete(this.set, n)
    }
    if this.count[key] == 0 {
        delete(this.count, key)
    } else {
        if this.set[n-1] == nil {
            this.set[n-1] = make(map[string]bool)
        }
        this.set[n-1][key] = true
    }
}

func (this *AllOne) GetMaxKey() string {
    if len(this.set) == 0 {
        return ""
    }
    for n := len(this.set); n > 0; n-- {
        for key := range this.set[n] {
            return key
        }
    }
    return ""
}

func (this *AllOne) GetMinKey() string {
    if len(this.set) == 0 {
        return ""
    }
    for n := 1; n <= len(this.set); n++ {
        for key := range this.set[n] {
            return key
        }
    }
    return ""
}