class Solution:
    def resultArray(self, nums: List[int], k: int) -> List[int]:
        result = [0] * k
        curr = [0] * k
        for num in nums:
            m = num % k
            next = [0] * k
            next[m] = 1
            for prev in range(k):
                if curr[prev] > 0:
                    nr = (prev * m) % k
                    next[nr] += curr[prev]
            for r in range(k):
                result[r] += next[r]
            curr = next
        return result