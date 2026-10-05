class Solution:
    def scoreOfParentheses(self, s: str) -> int:
        score = depth = 0
        for i in range(len(s)):
            if s[i] == '(':
                depth += 1
            else:
                depth -= 1
                if s[i - 1] == '(':
                    score += 1 << depth
        return score