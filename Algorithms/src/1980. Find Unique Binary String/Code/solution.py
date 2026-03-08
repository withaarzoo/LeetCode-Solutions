class Solution:
    def findDifferentBinaryString(self, nums: List[str]) -> str:
        n = len(nums)
        result = []

        # Flip diagonal bits
        for i in range(n):
            if nums[i][i] == '0':
                result.append('1')
            else:
                result.append('0')

        return "".join(result)