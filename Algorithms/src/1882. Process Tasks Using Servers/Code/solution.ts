function assignTasks(servers: number[], tasks: number[]): number[] {
    const free = new MinPriorityQueue({ compare: (a: number[], b: number[]) => a[0] !== b[0] ? a[0] - b[0] : a[1] - b[1] });
    for (let i = 0; i < servers.length; i++) free.enqueue([servers[i], i]);
    const busy = new MinPriorityQueue({ compare: (a: number[], b: number[]) => {
        if (a[0] !== b[0]) return a[0] - b[0];
        if (a[1] !== b[1]) return a[1] - b[1];
        return a[2] - b[2];
    }});
    const m = tasks.length;
    const ans = new Array(m);
    for (let j = 0; j < m; j++) {
        while (!busy.isEmpty() && busy.front()[0] <= j) {
            const cur = busy.dequeue();
            free.enqueue([cur[1], cur[2]]);
        }
        if (!free.isEmpty()) {
            const cur = free.dequeue();
            ans[j] = cur[1];
            busy.enqueue([j + tasks[j], cur[0], cur[1]]);
        } else {
            const cur = busy.dequeue();
            ans[j] = cur[2];
            busy.enqueue([cur[0] + tasks[j], cur[1], cur[2]]);
        }
    }
    return ans;
};