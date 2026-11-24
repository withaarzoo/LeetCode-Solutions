function maxLength(nums: number[]): number {
    const n = nums.length;
    let ans = 0;
    for (let i = 0; i < n; i++) {
        let prod = 1;
        let g = 0;
        let lc = 1;
        for (let j = i; j < n; j++) {
            const x = nums[j];
            if (g === 0) g = x;
            else {
                let a = g, b = x;
                while (b) {
                    const t = b;
                    b = a % b;
                    a = t;
                }
                g = a;
            }
            let gg = lc;
            let b = x;
            while (b) {
                const t = b;
                b = gg % b;
                gg = t;
            }
            if (lc / gg > Number.MAX_SAFE_INTEGER / x) break;
            lc = lc / gg * x;
            if (prod > Number.MAX_SAFE_INTEGER / x) break;
            prod *= x;
            if (prod === lc * g) ans = Math.max(ans, j - i + 1);
        }
    }
    return ans;
};