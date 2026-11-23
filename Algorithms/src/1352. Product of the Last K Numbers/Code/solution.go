type ProductOfNumbers struct {
    pre []int
}


func Constructor() ProductOfNumbers {
    return ProductOfNumbers{pre: []int{1}}
}


func (this *ProductOfNumbers) Add(num int)  {
    if num == 0 {
        this.pre = []int{1}
    } else {
        this.pre = append(this.pre, this.pre[len(this.pre)-1]*num)
    }
}


func (this *ProductOfNumbers) GetProduct(k int) int {
    if k >= len(this.pre) {
        return 0
    }
    return this.pre[len(this.pre)-1] / this.pre[len(this.pre)-1-k]
}