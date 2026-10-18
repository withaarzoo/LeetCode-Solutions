/**
 * @param {number[]} nums
 * @return {number}
 */
var numOfWays = function(nums) {
    const MOD = 1000000007n;
    const n = nums.length;
    const fact = new Array(n + 1).fill(1n);
    const inv = new Array(n + 1).fill(1n);
    for (let i = 1; i <= n; i++) fact[i] = fact[i - 1] * BigInt(i) % MOD;
    inv[n] = modPow(fact[n], MOD - 2n, MOD);
    for (let i = n - 1; i >= 0; i--) inv[i] = inv[i + 1] * BigInt(i + 1) % MOD;
    function comb(a, b) {
        if (b < 0 || b > a) return 0n;
        return fact[a] * inv[b] % MOD * inv[a - b] % MOD;
    }
    function ways(arr) {
        const m = arr.length;
        if (m <= 2) return 1n;
        const left = [], right = [];
        for (const x of arr) {
            if (x < arr[0]) left.push(x);
            else if (x > arr[0]) right.push(x);
        }
        return comb(m - 1, left.length) * ways(left) % MOD * ways(right) % MOD;
    }
    return Number((ways(nums) - 1n + MOD) % MOD);
};
function modPow(base, exp, mod) {
    let res = 1n;
    while (exp > 0n) {
        if (exp & 1n) res = res * base % mod;
        base = base * base % mod;
        exp >>= 1n;
    }
    return res;
}