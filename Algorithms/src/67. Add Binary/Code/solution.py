class Solution:
    def addBinary(self, a: str, b: str) -> str:
        i = len(a) - 1   # pointer for a
        j = len(b) - 1   # pointer for b
        carry = 0
        result = []
        
        while i >= 0 or j >= 0 or carry:
            total = carry
            
            # Add digit from a
            if i >= 0:
                total += int(a[i])
                i -= 1
            
            # Add digit from b
            if j >= 0:
                total += int(b[j])
                j -= 1
            
            result.append(str(total % 2))
            carry = total // 2
        
        # Reverse the result to correct order
        return "".join(result[::-1])
