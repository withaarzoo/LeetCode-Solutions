const MOD int64 = 1e9 + 7

type Fancy struct {
    seq []int64
    mul int64
    add int64
}

func modPow(a, b int64) int64 {
    res := int64(1)
    for b > 0 {
        if b&1 == 1 {
            res = res * a % MOD
        }
        a = a * a % MOD
        b >>= 1
    }
    return res
}

func Constructor() Fancy {
    return Fancy{[]int64{}, 1, 0}
}

func (this *Fancy) Append(val int) {
    inv := modPow(this.mul, MOD-2)
    stored := ((int64(val)-this.add+MOD)%MOD * inv) % MOD
    this.seq = append(this.seq, stored)
}

func (this *Fancy) AddAll(inc int) {
    this.add = (this.add + int64(inc)) % MOD
}

func (this *Fancy) MultAll(m int) {
    this.mul = this.mul * int64(m) % MOD
    this.add = this.add * int64(m) % MOD
}

func (this *Fancy) GetIndex(idx int) int {
    if idx >= len(this.seq) {
        return -1
    }
    return int((this.seq[idx]*this.mul%MOD + this.add) % MOD)
}