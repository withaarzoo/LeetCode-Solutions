class Solution:
    def makeLargestSpecial(self, s: str) -> str:
        parts = []
        count = 0
        start = 0
        
        for i, ch in enumerate(s):
            if ch == '1':
                count += 1
            else:
                count -= 1
            
            if count == 0:
                # Recursively solve inner substring
                inner = self.makeLargestSpecial(s[start+1:i])
                parts.append("1" + inner + "0")
                start = i + 1
        
        # Sort in descending order
        parts.sort(reverse=True)
        
        return "".join(parts)