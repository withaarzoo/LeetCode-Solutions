class Solution:
    def hasAllCodes(self, s: str, k: int) -> bool:
        n = len(s)
        
        if n < k:
            return False
        
        total = 1 << k
        if n - k + 1 < total:
            return False
        
        seen = [False] * total
        mask = total - 1
        
        curr = 0
        count = 0
        
        # First window
        for i in range(k):
            curr = (curr << 1) | int(s[i])
        
        if not seen[curr]:
            seen[curr] = True
            count += 1
        
        # Sliding window
        for i in range(k, n):
            curr = ((curr << 1) & mask) | int(s[i])
            
            if not seen[curr]:
                seen[curr] = True
                count += 1
                if count == total:
                    return True
        
        return count == total