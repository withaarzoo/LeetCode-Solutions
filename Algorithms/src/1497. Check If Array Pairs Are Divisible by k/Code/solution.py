class Solution:
    def canArrange(self, arr: List[int], k: int) -> bool:
        # Frequency array to store the count of remainders
        remainderFreq = [0] * k
        
        # Step 1: Calculate the remainder for each element and store the frequency
        for num in arr:
            remainder = (num % k + k) % k  # Ensure non-negative remainder
            remainderFreq[remainder] += 1
        
        # Step 2: Check if the pairing condition holds
        for i in range(k // 2 + 1):
            if i == 0:
                # Elements with remainder 0 must pair among themselves
                if remainderFreq[i] % 2 != 0:
                    return False
            else:
                # Remainder i must pair with remainder k-i
                if remainderFreq[i] != remainderFreq[k - i]:
                    return False
        
        return True