type CustomStack struct {
    maxSize int
    stack   []int
    inc     []int
}

func Constructor(maxSize int) CustomStack {
    return CustomStack{
        maxSize: maxSize,
        stack:   []int{},
        inc:     []int{},
    }
}

func (this *CustomStack) Push(x int) {
    if len(this.stack) < this.maxSize {
        this.stack = append(this.stack, x)
        this.inc = append(this.inc, 0)  // Initialize increment for this element
    }
}

func (this *CustomStack) Pop() int {
    if len(this.stack) == 0 {
        return -1
    }
    idx := len(this.stack) - 1
    result := this.stack[idx] + this.inc[idx]  // Apply any pending increments
    if idx > 0 {
        this.inc[idx-1] += this.inc[idx]  // Propagate increment to the next element
    }
    this.stack = this.stack[:idx]
    this.inc = this.inc[:idx]
    return result
}

func (this *CustomStack) Increment(k int, val int) {
    limit := min(k, len(this.stack)) - 1
    if limit >= 0 {
        this.inc[limit] += val  // Add increment to the bottom k-th element
    }
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}