/**
 * @param {number[]} nums1
 * @param {number[]} nums2
 * @param {number} k1
 * @param {number} k2
 * @return {number}
 */
var minSumSquareDiff = function(nums1, nums2, k1, k2) {
    const n = nums1.length;
    const d = new Array(n);
    let total = 0;
    let mx = 0;
    for (let i = 0; i < n; ++i) {
        d[i] = Math.abs(nums1[i] - nums2[i]);
        total += d[i];
        mx = Math.max(mx, d[i]);
    }
    let k = k1 + k2;
    if (total <= k) return 0;
    let left = 0, right = mx;
    while (left < right) {
        const mid = left + ((right - left) >> 1);
        let need = 0;
        for (const v of d) need += Math.max(0, v - mid);
        if (need <= k) right = mid;
        else left = mid + 1;
    }
    for (let i = 0; i < n; ++i) {
        k -= Math.max(0, d[i] - left);
        d[i] = Math.min(d[i], left);
    }
    for (let i = 0; i < n && k > 0; ++i) {
        if (d[i] === left) {
            --d[i];
            --k;
        }
    }
    let ans = 0;
    for (const v of d) ans += v * v;
    return ans;
};