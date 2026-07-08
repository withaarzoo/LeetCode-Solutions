class Solution {
    public int[] sumAndMultiply(String s, int[][] queries) {
        // I use long values because multiplication can exceed the int range.
        final long MOD = 1_000_000_007L;
        int n = s.length();

        // nonZeroCount[i] = number of non-zero digits in s[0..i-1].
        int[] nonZeroCount = new int[n + 1];

        // I first count all non-zero digits so I can allocate the exact array size.
        for (int i = 0; i < n; i++) {
            // I extend the previous prefix count by one only for a non-zero digit.
            nonZeroCount[i + 1] = nonZeroCount[i] + (s.charAt(i) != '0' ? 1 : 0);
        }

        int k = nonZeroCount[n];

        // digits stores the string after removing every zero.
        int[] digits = new int[k];
        int index = 0;

        // I fill the compressed digit array in the original order.
        for (int i = 0; i < n; i++) {
            if (s.charAt(i) != '0') {
                digits[index++] = s.charAt(i) - '0';
            }
        }

        // prefixValue[i] stores the first i compressed digits as a number modulo MOD.
        long[] prefixValue = new long[k + 1];

        // prefixSum[i] stores the sum of the first i compressed digits.
        long[] prefixSum = new long[k + 1];

        // power10[i] stores 10^i modulo MOD.
        long[] power10 = new long[k + 1];
        power10[0] = 1;

        // I build all prefix arrays in one pass over the compressed digits.
        for (int i = 0; i < k; i++) {
            // I append the current digit to the previous prefix number.
            prefixValue[i + 1] = (prefixValue[i] * 10 + digits[i]) % MOD;

            // I add the current digit to the previous digit sum.
            prefixSum[i + 1] = prefixSum[i] + digits[i];

            // I compute the next power of 10 for later range extraction.
            power10[i + 1] = (power10[i] * 10) % MOD;
        }

        // I create one result slot for every query.
        int[] answer = new int[queries.length];

        // I answer each query using only prefix-array lookups.
        for (int i = 0; i < queries.length; i++) {
            int l = queries[i][0];
            int r = queries[i][1];

            // left counts non-zero digits before the query.
            int left = nonZeroCount[l];

            // right counts non-zero digits through the end of the query.
            int right = nonZeroCount[r + 1];

            // len is the number of digits in the compressed query number.
            int len = right - left;

            // I subtract the shifted earlier prefix to isolate the required number.
            long x = (prefixValue[right]
                    - (prefixValue[left] * power10[len]) % MOD
                    + MOD) % MOD;

            // I subtract digit-sum prefixes to get the query digit sum.
            long sum = prefixSum[right] - prefixSum[left];

            // I store x multiplied by its digit sum under modulo.
            answer[i] = (int) ((x * sum) % MOD);
        }

        return answer;
    }
}