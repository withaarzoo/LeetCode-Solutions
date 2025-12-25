class Solution:
    def maximumHappinessSum(self, happiness: List[int], k: int) -> int:
        # Sort in descending order
        happiness.sort(reverse=True)
        
        ans = 0
        
        for i in range(k):
            curr = happiness[i] - i
            if curr > 0:
                ans += curr
        
        return ans
