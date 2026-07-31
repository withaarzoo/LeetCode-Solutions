class Solution {
    public int minimumPushes(String word) {
        // Store the frequency of every lowercase letter
        int[] freq = new int[26];

        // Count each letter
        for (char ch : word.toCharArray()) {
            freq[ch - 'a']++;
        }

        // Sort frequencies in increasing order
        Arrays.sort(freq);

        int ans = 0;
        int index = 0;

        // Traverse from the largest frequency to the smallest
        for (int i = 25; i >= 0; i--) {
            // Ignore unused letters
            if (freq[i] == 0) break;

            // Every 8 letters need one extra push
            int pushes = (index / 8) + 1;

            // Add this letter's contribution
            ans += freq[i] * pushes;

            index++;
        }

        return ans;
    }
}