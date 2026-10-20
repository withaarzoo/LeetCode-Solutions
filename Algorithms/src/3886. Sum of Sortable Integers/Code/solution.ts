function sortableIntegers(nums: number[]): number {
    const n = nums.length;
    const sorted = [...nums].sort((a, b) => a - b);
    const divs: number[] = [];
    for (let i = 1; i * i <= n; i++) {
        if (n % i === 0) {
            divs.push(i);
            if (i !== n / i) divs.push(n / i);
        }
    }
    let ans = 0;
    for (const k of divs) {
        let ok = true;
        for (let s = 0; s < n && ok; s += k) {
            let desc = 0, pos = -1;
            for (let j = 0; j < k; j++) {
                const a = nums[s + j];
                const b = nums[s + (j + 1) % k];
                if (a > b) {
                    desc++;
                    pos = j;
                    if (desc > 1) break;
                }
            }
            if (desc > 1) {
                ok = false;
                break;
            }
            const r = desc === 0 ? 0 : (pos + 1) % k;
            for (let j = 0; j < k; j++) {
                if (nums[s + (r + j) % k] !== sorted[s + j]) {
                    ok = false;
                    break;
                }
            }
        }
        if (ok) ans += k;
    }
    return ans;
}