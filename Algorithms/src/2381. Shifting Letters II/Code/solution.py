class Solution:
    def shiftingLetters(self, s: str, shifts: List[List[int]]) -> str:
        n = len(s)
        diff = [0] * (n + 1)

        # Build the difference array
        for start, end, direction in shifts:
            delta = 1 if direction == 1 else -1
            diff[start] += delta
            if end + 1 < n:
                diff[end + 1] -= delta

        # Calculate cumulative shifts
        shift = 0
        result = list(s)
        for i in range(n):
            shift += diff[i]
            shift = (shift % 26 + 26) % 26  # Normalize shift to [0, 25]
            result[i] = chr((ord(result[i]) - ord('a') + shift) % 26 + ord('a'))

        return ''.join(result)
