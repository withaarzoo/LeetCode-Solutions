var minimumDeleteSum = function(s1, s2) {
    const n = s1.length, m = s2.length;
    let dp = new Array(m + 1).fill(0);

    for (let j = m - 1; j >= 0; j--) {
        dp[j] = dp[j + 1] + s2.charCodeAt(j);
    }

    for (let i = n - 1; i >= 0; i--) {
        let prev = dp[m];
        dp[m] += s1.charCodeAt(i);

        for (let j = m - 1; j >= 0; j--) {
            let temp = dp[j];
            if (s1[i] === s2[j]) {
                dp[j] = prev;
            } else {
                dp[j] = Math.min(
                    s1.charCodeAt(i) + dp[j],
                    s2.charCodeAt(j) + dp[j + 1]
                );
            }
            prev = temp;
        }
    }
    return dp[0];
};
