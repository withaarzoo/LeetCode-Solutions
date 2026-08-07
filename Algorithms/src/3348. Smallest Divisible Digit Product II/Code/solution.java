class Solution {
    public String smallestNumber(String num, long t) {
        int req2 = 0, req3 = 0, req5 = 0, req7 = 0;
        long temp = t;
        // Strip out allowed prime factors
        while (temp % 2 == 0) {
            temp /= 2;
            req2++;
        }
        while (temp % 3 == 0) {
            temp /= 3;
            req3++;
        }
        while (temp % 5 == 0) {
            temp /= 5;
            req5++;
        }
        while (temp % 7 == 0) {
            temp /= 7;
            req7++;
        }
        // Detect unsupported prime factors immediately
        if (temp > 1)
            return "-1";

        int[][] dp = new int[60][40];
        for (int i = 0; i < 60; ++i) {
            for (int j = 0; j < 40; ++j) {
                dp[i][j] = 1000000000;
            }
        }
        dp[0][0] = 0;

        // Apply digit transitions that generate 2s and 3s
        int[][] trans = { { 1, 0 }, { 0, 1 }, { 2, 0 }, { 1, 1 }, { 3, 0 }, { 0, 2 } };
        for (int i = 0; i < 60; ++i) {
            for (int j = 0; j < 40; ++j) {
                if (dp[i][j] == 1000000000)
                    continue;
                for (int[] tr : trans) {
                    int ni = Math.min(59, i + tr[0]);
                    int nj = Math.min(39, j + tr[1]);
                    dp[ni][nj] = Math.min(dp[ni][nj], dp[i][j] + 1);
                }
            }
        }

        // Propagate backwards so looking up dp[i][j] covers needing AT LEAST i and j
        for (int i = 59; i >= 0; --i) {
            for (int j = 39; j >= 0; --j) {
                if (i < 59)
                    dp[i][j] = Math.min(dp[i][j], dp[i + 1][j]);
                if (j < 39)
                    dp[i][j] = Math.min(dp[i][j], dp[i][j + 1]);
            }
        }

        int[] F2 = { 0, 0, 1, 0, 2, 0, 1, 0, 3, 0 };
        int[] F3 = { 0, 0, 0, 1, 0, 0, 1, 0, 0, 2 };
        int[] F5 = { 0, 0, 0, 0, 0, 1, 0, 0, 0, 0 };
        int[] F7 = { 0, 0, 0, 0, 0, 0, 0, 1, 0, 0 };

        int n = num.length();
        boolean hasZero = false;
        int firstZero = n;
        for (int i = 0; i < n; ++i) {
            if (num.charAt(i) == '0') {
                hasZero = true;
                firstZero = i;
                break;
            }
        }

        // Ensure the input naturally fulfills the requirements without any changes
        if (!hasZero) {
            int r2 = req2, r3 = req3, r5 = req5, r7 = req7;
            for (int i = 0; i < n; i++) {
                int d = num.charAt(i) - '0';
                r2 = Math.max(0, r2 - F2[d]);
                r3 = Math.max(0, r3 - F3[d]);
                r5 = Math.max(0, r5 - F5[d]);
                r7 = Math.max(0, r7 - F7[d]);
            }
            if (r2 == 0 && r3 == 0 && r5 == 0 && r7 == 0)
                return num;
        }

        // Limit the scope to indices before a '0' occurs
        int limit = Math.min(n - 1, firstZero);
        int p2 = 0, p3 = 0, p5 = 0, p7 = 0;
        for (int i = 0; i < limit; ++i) {
            int d = num.charAt(i) - '0';
            p2 += F2[d];
            p3 += F3[d];
            p5 += F5[d];
            p7 += F7[d];
        }

        // Step backwards to modify the rightmost possible viable digit
        for (int i = limit; i >= 0; --i) {
            int startD = (num.charAt(i) - '0') + 1;
            for (int d = startD; d <= 9; ++d) {
                int n2 = Math.max(0, req2 - p2 - F2[d]);
                int n3 = Math.max(0, req3 - p3 - F3[d]);
                int n5 = Math.max(0, req5 - p5 - F5[d]);
                int n7 = Math.max(0, req7 - p7 - F7[d]);
                int L = n - 1 - i;

                // If factors fit, finalize string dynamically
                if (n7 + n5 + dp[n2][n3] <= L) {
                    StringBuilder ans = new StringBuilder(num.substring(0, i));
                    ans.append(d);
                    int rem2 = n2, rem3 = n3, rem5 = n5, rem7 = n7;
                    // Choose lexicographically smallest valid digit iteratively
                    for (int pos = 0; pos < L; ++pos) {
                        for (int x = 1; x <= 9; ++x) {
                            int nn2 = Math.max(0, rem2 - F2[x]);
                            int nn3 = Math.max(0, rem3 - F3[x]);
                            int nn5 = Math.max(0, rem5 - F5[x]);
                            int nn7 = Math.max(0, rem7 - F7[x]);
                            if (nn7 + nn5 + dp[nn2][nn3] <= L - 1 - pos) {
                                ans.append(x);
                                rem2 = nn2;
                                rem3 = nn3;
                                rem5 = nn5;
                                rem7 = nn7;
                                break;
                            }
                        }
                    }
                    return ans.toString();
                }
            }
            if (i > 0) {
                int d = num.charAt(i - 1) - '0';
                p2 -= F2[d];
                p3 -= F3[d];
                p5 -= F5[d];
                p7 -= F7[d];
            }
        }

        // Start from scratch if no modified prefix fits the target constraint
        int minLenNeeded = req7 + req5 + dp[req2][req3];
        int M = Math.max(n + 1, minLenNeeded);
        StringBuilder ans = new StringBuilder();
        int rem2 = req2, rem3 = req3, rem5 = req5, rem7 = req7;

        for (int pos = 0; pos < M; ++pos) {
            for (int x = 1; x <= 9; ++x) {
                int nn2 = Math.max(0, rem2 - F2[x]);
                int nn3 = Math.max(0, rem3 - F3[x]);
                int nn5 = Math.max(0, rem5 - F5[x]);
                int nn7 = Math.max(0, rem7 - F7[x]);
                if (nn7 + nn5 + dp[nn2][nn3] <= M - 1 - pos) {
                    ans.append(x);
                    rem2 = nn2;
                    rem3 = nn3;
                    rem5 = nn5;
                    rem7 = nn7;
                    break;
                }
            }
        }
        return ans.toString();
    }
}