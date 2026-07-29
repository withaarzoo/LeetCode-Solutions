class Solution {
    public String smallestPalindrome(String s, int k) {
        // Track global character frequencies
        int[] freq = new int[26];
        for (char c : s.toCharArray()) {
            freq[c - 'a']++;
        }

        // Track half frequencies and identify the middle character
        int[] half = new int[26];
        StringBuilder mid = new StringBuilder();
        int m = 0;

        for (int i = 0; i < 26; ++i) {
            if (freq[i] % 2 != 0) {
                mid.append((char) (i + 'a'));
            }
            half[i] = freq[i] / 2;
            m += half[i];
        }

        // Initial permutation check to see if target k is reachable
        if (getWays(half, k) < k) {
            return "";
        }

        StringBuilder firstHalf = new StringBuilder();
        // Construct left half from left to right
        for (int i = 0; i < m; ++i) {
            for (int c = 0; c < 26; ++c) {
                if (half[c] > 0) {
                    // Temporarily claim the character
                    half[c]--;
                    long ways = getWays(half, k);

                    // Target string is within the branches of this character choice
                    if (ways >= k) {
                        firstHalf.append((char) (c + 'a'));
                        break;
                    } else {
                        // Skip character and decrease our k target
                        k -= ways;
                        half[c]++;
                    }
                }
            }
        }

        // Mirror the first half to complete the palindrome
        StringBuilder res = new StringBuilder(firstHalf);
        res.append(mid);
        res.append(firstHalf.reverse());
        return res.toString();
    }

    // Helper to calculate total permutations of the remaining multiset
    private long getWays(int[] f, long targetK) {
        long ways = 1;
        int currLen = 0;
        for (int count : f) {
            if (count > 0) {
                currLen += count;
                long n = currLen;
                long r = count;

                if (r > n - r)
                    r = n - r;
                long curNCr = 1;

                // Calculate combinations with an early termination mechanism
                for (int i = 1; i <= r; ++i) {
                    curNCr = curNCr * (n - i + 1) / i;
                    if (curNCr > targetK) {
                        curNCr = targetK + 1;
                        break;
                    }
                }
                ways *= curNCr;
                if (ways > targetK)
                    return targetK + 1;
            }
        }
        return ways;
    }
}