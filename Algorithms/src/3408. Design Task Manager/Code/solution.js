var TaskManager = function (tasks) {
  // Define MaxHeap inside so it doesn't clash between runs
  class MaxHeap {
    constructor() {
      this.heap = [];
    }
    size() {
      return this.heap.length;
    }
    _cmp(a, b) {
      if (a.priority !== b.priority) return a.priority > b.priority;
      return a.taskId > b.taskId;
    }
    _swap(i, j) {
      [this.heap[i], this.heap[j]] = [this.heap[j], this.heap[i]];
    }
    push(item) {
      this.heap.push(item);
      this._siftUp(this.heap.length - 1);
    }
    pop() {
      if (this.heap.length === 0) return null;
      const top = this.heap[0];
      const last = this.heap.pop();
      if (this.heap.length > 0) {
        this.heap[0] = last;
        this._siftDown(0);
      }
      return top;
    }
    _siftUp(idx) {
      while (idx > 0) {
        const parent = Math.floor((idx - 1) / 2);
        if (this._cmp(this.heap[idx], this.heap[parent])) {
          this._swap(idx, parent);
          idx = parent;
        } else break;
      }
    }
    _siftDown(idx) {
      const n = this.heap.length;
      while (true) {
        let largest = idx;
        const l = 2 * idx + 1,
          r = 2 * idx + 2;
        if (l < n && this._cmp(this.heap[l], this.heap[largest])) largest = l;
        if (r < n && this._cmp(this.heap[r], this.heap[largest])) largest = r;
        if (largest === idx) break;
        this._swap(idx, largest);
        idx = largest;
      }
    }
  }

  this.map = new Map(); // taskId -> {userId, priority}
  this.heap = new MaxHeap();

  for (const t of tasks) {
    if (t.length < 3) continue;
    const [user, task, pr] = t;
    this.map.set(task, { userId: user, priority: pr });
    this.heap.push({ priority: pr, taskId: task });
  }
};

TaskManager.prototype.add = function (userId, taskId, priority) {
  this.map.set(taskId, { userId, priority });
  this.heap.push({ priority, taskId });
};

TaskManager.prototype.edit = function (taskId, newPriority) {
  const cur = this.map.get(taskId);
  cur.priority = newPriority;
  this.heap.push({ priority: newPriority, taskId });
};

TaskManager.prototype.rmv = function (taskId) {
  this.map.delete(taskId);
};

TaskManager.prototype.execTop = function () {
  while (this.heap.size() > 0) {
    const top = this.heap.pop();
    const rec = this.map.get(top.taskId);
    if (!rec) continue; // removed
    if (rec.priority !== top.priority) continue; // stale
    this.map.delete(top.taskId);
    return rec.userId;
  }
  return -1;
};

/**
 * Your TaskManager object will be instantiated and called as such:
 * var obj = new TaskManager(tasks)
 * obj.add(userId,taskId,priority)
 * obj.edit(taskId,newPriority)
 * obj.rmv(taskId)
 * var param_4 = obj.execTop()
 */
