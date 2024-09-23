class Solution:
    def minExtraChar(self, s: str, dictionary: List[str]) -> int:
        # Convert the dictionary list to a set for O(1) time complexity lookups
        dict_set = set(dictionary) 
        
        # Length of the input string s
        n = len(s) 
        
        # Create a DP array where dp[i] represents the minimum number of extra characters
        # needed to parse the substring s[0:i]. Initialize with the maximum value (n) 
        # since the worst case is that all characters are extra.
        dp = [n] * (n + 1) 
        
        # Base case: No extra characters needed for an empty string
        dp[0] = 0 
        
        # Iterate through the string s, starting from index 1 up to n
        for i in range(1, n + 1):
            # For each i, check every possible substring s[j:i], where j < i
            for j in range(i):
                # Get the substring s[j:i]
                sub = s[j:i]  
                
                # If the substring is in the dictionary, it means we don't need to consider
                # these characters as extra, so we minimize dp[i] by considering dp[j].
                if sub in dict_set:
                    dp[i] = min(dp[i], dp[j])  # Update dp[i] if substring is found in dictionary
            
            # If no valid substring is found, we treat the current character s[i-1] as extra.
            # Therefore, we update dp[i] by considering the case where we treat the current 
            # character as extra and add 1 to dp[i-1].
            dp[i] = min(dp[i], dp[i - 1] + 1)  
        
        # The final answer will be stored in dp[n], which represents the minimum extra characters
        # needed for the entire string s.
        return dp[n]
