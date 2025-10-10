class Solution:
    def maximumEnergy(self, energy: List[int], k: int) -> int:
        n = len(energy)
        ans = -10**18  # very small initial value
        for r in range(k):
            cur = 0
            # last index in this residue class
            last = r + ((n - 1 - r) // k) * k
            i = last
            while i >= r:
                cur += energy[i]   # suffix sum starting at i
                ans = max(ans, cur)
                i -= k
        return ans
