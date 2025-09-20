/**
 * @param {number} memoryLimit
 */
var Router = function (memoryLimit) {
  this.memoryLimit = memoryLimit;
  this.queue = []; // array of [s,d,t]
  this.qHead = 0; // front index
  this.size = 0; // current stored count
  this.seen = new Set(); // keys "s#d#t"
  this.dest = new Map(); // dest -> {arr: [...timestamps...], head: number}
};

function makeKey(s, d, t) {
  return s + "#" + d + "#" + t;
}

/**
 * @param {number} source
 * @param {number} destination
 * @param {number} timestamp
 * @return {boolean}
 */
Router.prototype.addPacket = function (source, destination, timestamp) {
  const key = makeKey(source, destination, timestamp);
  if (this.seen.has(key)) return false;

  // Evict oldest until there is room
  while (this.size >= this.memoryLimit) {
    const old = this.queue[this.qHead++];
    this.size--;
    const oldKey = makeKey(old[0], old[1], old[2]);
    this.seen.delete(oldKey);
    const dObj = this.dest.get(old[1]);
    dObj.head += 1;
  }

  // Append new packet
  this.queue.push([source, destination, timestamp]);
  this.size++;
  this.seen.add(key);

  if (!this.dest.has(destination))
    this.dest.set(destination, { arr: [], head: 0 });
  this.dest.get(destination).arr.push(timestamp);
  return true;
};

/**
 * @return {number[]}
 */
Router.prototype.forwardPacket = function () {
  if (this.size === 0) return [];
  const pkt = this.queue[this.qHead++];
  this.size--;
  const key = makeKey(pkt[0], pkt[1], pkt[2]);
  this.seen.delete(key);
  const dObj = this.dest.get(pkt[1]);
  dObj.head += 1;
  return pkt;
};

// Binary search helpers that accept lo index
function lowerBound(arr, target, lo) {
  let l = lo,
    r = arr.length;
  while (l < r) {
    const m = (l + r) >> 1;
    if (arr[m] < target) l = m + 1;
    else r = m;
  }
  return l;
}
function upperBound(arr, target, lo) {
  let l = lo,
    r = arr.length;
  while (l < r) {
    const m = (l + r) >> 1;
    if (arr[m] <= target) l = m + 1;
    else r = m;
  }
  return l;
}

/**
 * @param {number} destination
 * @param {number} startTime
 * @param {number} endTime
 * @return {number}
 */
Router.prototype.getCount = function (destination, startTime, endTime) {
  if (!this.dest.has(destination)) return 0;
  const dObj = this.dest.get(destination);
  const arr = dObj.arr;
  const h = dObj.head;
  const L = lowerBound(arr, startTime, h);
  const R = upperBound(arr, endTime, h);
  return R - L;
};

/**
 * Your Router object will be instantiated and called as such:
 * var obj = new Router(memoryLimit)
 * var param_1 = obj.addPacket(source,destination,timestamp)
 * var param_2 = obj.forwardPacket()
 * var param_3 = obj.getCount(destination,startTime,endTime)
 */
