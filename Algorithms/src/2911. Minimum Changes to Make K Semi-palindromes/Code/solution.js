/**
 * @param {string} s
 * @param {number} k
 * @return {number}
 */
var minimumChanges = function(s, k) {
    const n = s.length;
    const INF = 1e9;
    const cost = Array.from({length: n}, () => Array(n).fill(INF));
    for (let i = 0; i < n; i++) {
        for (let j = i + 1; j < n; j++) {
            const L = j - i + 1;
            let best = INF;
            for (let d = 1; d < L; d++) {
                if (L % d !== 0) continue;
                let cur = 0;
                const m = L / d;
                for (let r = 0; r < d; r++) {
                    for (let t = 0; t < Math.floor(m / 2); t++) {
                        if (s[i + r + t * d] !== s[i + r + (m - 1 - t) * d])
                            cur++;
                    }
                }
                best = Math.min(best, cur);
            }
            cost[i][j] = best;
        }
    }
    const dp = Array.from({length: n + 1}, () => Array(k + 1).fill(INF));
    dp[0][0] = 0;
    for (let i = 1; i <= n; i++) {
        for (let p = 1; p <= k; p++) {
            for (let prev = 0; prev <= i - 2; prev++) {
                if (dp[prev][p - 1] !== INF && cost[prev][i - 1] !== INF)
                    dp[i][p] = Math.min(dp[i][p], dp[prev][p - 1] + cost[prev][i - 1]);
            }
        }
    }
    return dp[n][k];
};