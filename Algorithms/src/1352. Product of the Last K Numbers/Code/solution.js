var ProductOfNumbers = function() {
    this.pre = [1];
};

ProductOfNumbers.prototype.add = function(num) {
    if (num === 0) {
        this.pre = [1];
    } else {
        this.pre.push(this.pre[this.pre.length - 1] * num);
    }
};

ProductOfNumbers.prototype.getProduct = function(k) {
    if (k >= this.pre.length) return 0;
    return this.pre[this.pre.length - 1] / this.pre[this.pre.length - 1 - k];
};