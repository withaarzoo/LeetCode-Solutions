class Solution:
    def findThePrefixCommonArray(self, A: List[int], B: List[int]) -> List[int]:
        
        n = len(A)
        
        # Frequency array to count appearances
        freq = [0] * (n + 1)
        
        # Final answer array
        ans = [0] * n
        
        # Stores count of common elements
        common = 0
        
        for i in range(n):
            
            # Add current element from A
            freq[A[i]] += 1
            
            # If frequency becomes 2,
            # number exists in both arrays
            if freq[A[i]] == 2:
                common += 1
            
            # Add current element from B
            freq[B[i]] += 1
            
            # Same logic for B
            if freq[B[i]] == 2:
                common += 1
            
            # Store answer for this prefix
            ans[i] = common
        
        return ans