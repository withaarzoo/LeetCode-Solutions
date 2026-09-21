function resultArray(nums: number[], k: number): number[] {
    let result = new Array(k).fill(0);
    let curr = new Array(k).fill(0);
    for (let num of nums) {
        let m = num % k;
        let next = new Array(k).fill(0);
        next[m] = 1;
        for (let prev = 0; prev < k; prev++) {
            if (curr[prev] > 0) {
                let nr = (prev * m) % k;
                next[nr] += curr[prev];
            }
        }
        for (let r = 0; r < k; r++) {
            result[r] += next[r];
        }
        curr = next;
    }
    return result;
};