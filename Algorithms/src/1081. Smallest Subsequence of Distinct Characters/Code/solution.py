class Solution:
    def smallestSubsequence(self, s: str) -> str:

        # Count remaining occurrences of every character
        freq = [0] * 26
        for ch in s:
            freq[ord(ch) - ord('a')] += 1

        # Track characters already inside the stack
        in_stack = [False] * 26

        # Stack to build the answer
        stack = []

        for ch in s:

            # Current occurrence has been processed
            freq[ord(ch) - ord('a')] -= 1

            # Skip duplicate characters
            if in_stack[ord(ch) - ord('a')]:
                continue

            # Remove larger characters that appear again later
            while (
                stack and
                stack[-1] > ch and
                freq[ord(stack[-1]) - ord('a')] > 0
            ):
                in_stack[ord(stack.pop()) - ord('a')] = False

            # Add current character
            stack.append(ch)
            in_stack[ord(ch) - ord('a')] = True

        # Convert stack into a string
        return "".join(stack)