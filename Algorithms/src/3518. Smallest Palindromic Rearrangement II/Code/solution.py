class Solution:
    def smallestPalindrome(self, s: str, k: int) -> str:
        from collections import Counter
        import math
        
        # Build raw frequencies of characters from the string
        freq = Counter(s)
        half = {}
        mid = ""
        m = 0
        
        # Calculate exactly half the occurrences and locate the odd character
        for char in "abcdefghijklmnopqrstuvwxyz":
            if freq[char] > 0:
                if freq[char] % 2 != 0:
                    mid += char
                half[char] = freq[char] // 2
                m += half[char]
        
        # Computes combinations using Python's native fast math
        def get_ways(f, target_k):
            ways = 1
            curr_len = 0
            for char in "abcdefghijklmnopqrstuvwxyz":
                count = f.get(char, 0)
                if count > 0:
                    curr_len += count
                    ways *= math.comb(curr_len, count)
                    # Returning early avoids computing huge factorials unnecessarily
                    if ways > target_k:
                        return target_k + 1
            return ways
            
        # Return blank string if k permutations do not exist
        if get_ways(half, k) < k:
            return ""
            
        first_half = []
        # Find exactly the right character for each step of the half length
        for _ in range(m):
            for char in "abcdefghijklmnopqrstuvwxyz":
                if half.get(char, 0) > 0:
                    half[char] -= 1
                    ways = get_ways(half, k)
                    
                    # Accept branch and move to next slot if ways encompasses k
                    if ways >= k:
                        first_half.append(char)
                        break
                    else:
                        # Otherwise, shift k and give back the char to check the next one
                        k -= ways
                        half[char] += 1
                        
        # Tie the pieces together in order
        first_str = "".join(first_half)
        return first_str + mid + first_str[::-1]