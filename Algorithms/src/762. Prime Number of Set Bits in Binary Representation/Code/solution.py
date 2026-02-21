class Solution:
    def countPrimeSetBits(self, left: int, right: int) -> int:
        # Prime set bits possible up to 20
        primes = {2,3,5,7,11,13,17,19}
        
        ans = 0
        
        for num in range(left, right + 1):
            # Count set bits
            set_bits = bin(num).count('1')
            
            if set_bits in primes:
                ans += 1
        
        return ans