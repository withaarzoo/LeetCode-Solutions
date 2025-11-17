class Solution:
    def kLengthApart(self, nums: List[int], k: int) -> bool:
        prev = -1  # index of last seen 1; -1 means none yet
        for i, val in enumerate(nums):
            if val == 1:
                if prev != -1:
                    # zeros between = i - prev - 1
                    if i - prev - 1 < k:
                        return False
                prev = i
        return True
