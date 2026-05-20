class Solution {
    public int[] findThePrefixCommonArray(int[] A, int[] B) {
        
        int n = A.length;
        
        // Frequency array to count appearances
        int[] freq = new int[n + 1];
        
        // Final answer array
        int[] ans = new int[n];
        
        // Current count of common elements
        int common = 0;
        
        for (int i = 0; i < n; i++) {
            
            // Process current element from A
            freq[A[i]]++;
            
            // If frequency becomes 2,
            // this number exists in both arrays
            if (freq[A[i]] == 2) {
                common++;
            }
            
            // Process current element from B
            freq[B[i]]++;
            
            // Same logic for B
            if (freq[B[i]] == 2) {
                common++;
            }
            
            // Store result for this prefix
            ans[i] = common;
        }
        
        return ans;
    }
}