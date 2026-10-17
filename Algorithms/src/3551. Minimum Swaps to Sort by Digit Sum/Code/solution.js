/**
 * @param {number[]} nums
 * @return {number}
 */
var minSwaps = function(nums) {
    const digitSum = (x) => {
        let s = 0;
        while (x > 0) {
            s += x % 10;
            x = Math.floor(x / 10);
        }
        return s;
    };
    const n = nums.length;
    const sorted = [...nums].sort((a, b) => {
        const sa = digitSum(a), sb = digitSum(b);
        if (sa !== sb) return sa - sb;
        return a - b;
    });
    const pos = new Map();
    for (let i = 0; i < n; i++) pos.set(sorted[i], i);
    const visited = new Array(n).fill(false);
    let cycles = 0;
    for (let i = 0; i < n; i++) {
        if (!visited[i]) {
            cycles++;
            let j = i;
            while (!visited[j]) {
                visited[j] = true;
                j = pos.get(nums[j]);
            }
        }
    }
    return n - cycles;
};