var CustomStack = function (maxSize) {
  this.maxSize = maxSize;
  this.stack = [];
  this.inc = [];
};

CustomStack.prototype.push = function (x) {
  if (this.stack.length < this.maxSize) {
    this.stack.push(x);
    this.inc.push(0); // Initialize increment for this element
  }
};

CustomStack.prototype.pop = function () {
  if (this.stack.length === 0) {
    return -1;
  }
  let idx = this.stack.length - 1;
  let result = this.stack[idx] + this.inc[idx]; // Apply any pending increments
  if (idx > 0) {
    this.inc[idx - 1] += this.inc[idx]; // Propagate increment to the next element
  }
  this.stack.pop();
  this.inc.pop();
  return result;
};

CustomStack.prototype.increment = function (k, val) {
  let limit = Math.min(k, this.stack.length) - 1;
  if (limit >= 0) {
    this.inc[limit] += val; // Add increment to the bottom k-th element
  }
};
