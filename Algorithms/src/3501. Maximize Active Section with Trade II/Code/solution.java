class Solution {
    public List<Integer> maxActiveSectionsAfterTrade(String s, int[][] queries) {
        int n = s.length(); // Length of string
        int totalOnes = 0; // Base count of 1s in the entire string

        // Count base ones
        for (int i = 0; i < n; i++) {
            if (s.charAt(i) == '1') {
                totalOnes++;
            }
        }

        // Use dynamic lists to store compressed segment data
        List<Integer> typeList = new ArrayList<>();
        List<Integer> startList = new ArrayList<>();
        List<Integer> endList = new ArrayList<>();

        // Group consecutive identical characters
        for (int i = 0; i < n;) {
            int j = i;
            // Advance as long as the character matches the segment's start
            while (j < n && s.charAt(j) == s.charAt(i)) {
                j++;
            }
            // Save segment info
            typeList.add(s.charAt(i) - '0');
            startList.add(i);
            endList.add(j - 1);
            i = j; // Move to next segment
        }

        int N = typeList.size(); // Total segment count
        // Arrays are faster in Java, so cast our lists to primitive arrays
        int[] type = new int[N];
        int[] start = new int[N];
        int[] endIdx = new int[N];
        for (int i = 0; i < N; i++) {
            type[i] = typeList.get(i);
            start[i] = startList.get(i);
            endIdx[i] = endList.get(i);
        }

        // Create an index mapping for O(1) segment lookups
        int[] posToSeg = new int[n];
        for (int i = 0; i < N; i++) {
            for (int j = start[i]; j <= endIdx[i]; j++) {
                posToSeg[j] = i;
            }
        }

        // Precalculate maximum possible gain if we swap around a given 1-segment
        int[] ans = new int[N];
        for (int i = 1; i < N - 1; i++) {
            if (type[i] == 1) {
                // Formula: length of left 0-segment + length of right 0-segment
                ans[i] = (endIdx[i - 1] - start[i - 1] + 1) + (endIdx[i + 1] - start[i + 1] + 1);
            }
        }

        // Prepare log array for Sparse Table
        int[] logTable = new int[N + 1];
        for (int i = 2; i <= N; i++) {
            logTable[i] = logTable[i / 2] + 1;
        }

        int K = logTable[N] + 1; // Max power of 2
        // Build Sparse Table for O(1) Range Maximum Queries
        int[][] st = new int[K][N];
        for (int i = 0; i < N; i++) {
            st[0][i] = ans[i];
        }

        for (int j = 1; j < K; j++) {
            for (int i = 0; i + (1 << j) <= N; i++) {
                st[j][i] = Math.max(st[j - 1][i], st[j - 1][i + (1 << (j - 1))]);
            }
        }

        List<Integer> res = new ArrayList<>(); // Store results

        // Process each query independently
        for (int[] q : queries) {
            int L = q[0]; // Query left bound
            int R = q[1]; // Query right bound

            // Map character indices to segment IDs
            int segL = posToSeg[L];
            int segR = posToSeg[R];

            // Need at least a 0-1-0 segment structure
            if (segR - segL < 2) {
                res.add(totalOnes);
                continue;
            }

            int maxGain = 0;
            // Check boundary-adjacent 1-segments (their 0s might be chopped off)
            maxGain = Math.max(maxGain, evaluateEdge(segL + 1, L, R, segL, segR, type, start, endIdx));
            maxGain = Math.max(maxGain, evaluateEdge(segR - 1, L, R, segL, segR, type, start, endIdx));

            // Check fully enclosed 1-segments using Sparse Table
            if (segL + 2 <= segR - 2) {
                int L_idx = segL + 2;
                int R_idx = segR - 2;
                int j = logTable[R_idx - L_idx + 1];
                int rmqVal = Math.max(st[j][L_idx], st[j][R_idx - (1 << j) + 1]);
                maxGain = Math.max(maxGain, rmqVal);
            }

            // Add max possible extra 1s to the initial total
            res.add(totalOnes + maxGain);
        }

        return res;
    }

    // Helper to evaluate partial segments manually
    private int evaluateEdge(int i, int L, int R, int segL, int segR, int[] type, int[] start, int[] endIdx) {
        // Exclude segments that touch edges because they lack surrounding 0s
        if (i <= segL || i >= segR)
            return 0;
        // Must be a 1-segment
        if (type[i] == 0)
            return 0;

        int leftLen = 0;
        // Truncate left 0-segment if it crosses L
        if (i - 1 == segL)
            leftLen = Math.max(0, endIdx[i - 1] - L + 1);
        else
            leftLen = endIdx[i - 1] - start[i - 1] + 1;

        int rightLen = 0;
        // Truncate right 0-segment if it crosses R
        if (i + 1 == segR)
            rightLen = Math.max(0, R - start[i + 1] + 1);
        else
            rightLen = endIdx[i + 1] - start[i + 1] + 1;

        return leftLen + rightLen; // Total gain
    }
}