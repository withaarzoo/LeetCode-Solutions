class ProductOfNumbers {
    pre: number[];
    constructor() {
        this.pre = [1];
    }

    add(num: number): void {
        if (num === 0) {
            this.pre = [1];
        } else {
            this.pre.push(this.pre[this.pre.length - 1] * num);
        }
    }

    getProduct(k: number): number {
        if (k >= this.pre.length) return 0;
        return this.pre[this.pre.length - 1] / this.pre[this.pre.length - 1 - k];
    }
}