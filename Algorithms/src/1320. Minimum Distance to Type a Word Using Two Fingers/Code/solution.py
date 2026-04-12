class Solution:
    def minimumDistance(self, word: str) -> int:
        from functools import lru_cache

        # Calculate Manhattan distance between two letters
        def get_dist(a, b):
            # 26 means finger is not placed yet
            if a == 26 or b == 26:
                return 0

            row1, col1 = divmod(a, 6)
            row2, col2 = divmod(b, 6)

            return abs(row1 - row2) + abs(col1 - col2)

        @lru_cache(None)
        def solve(idx, f1, f2):
            # If all characters are typed
            if idx == len(word):
                return 0

            cur = ord(word[idx]) - ord('A')

            # Option 1: Use finger 1
            use_finger1 = get_dist(f1, cur) + solve(idx + 1, cur, f2)

            # Option 2: Use finger 2
            use_finger2 = get_dist(f2, cur) + solve(idx + 1, f1, cur)

            return min(use_finger1, use_finger2)

        # Both fingers initially not placed
        return solve(0, 26, 26)