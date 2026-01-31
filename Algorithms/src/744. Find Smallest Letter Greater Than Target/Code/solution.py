class Solution:
    def nextGreatestLetter(self, letters: List[str], target: str) -> str:
        left, right = 0, len(letters) - 1
        answer = letters[0]  # wrap-around default

        while left <= right:
            mid = (left + right) // 2

            if letters[mid] > target:
                answer = letters[mid]
                right = mid - 1
            else:
                left = mid + 1

        return answer
