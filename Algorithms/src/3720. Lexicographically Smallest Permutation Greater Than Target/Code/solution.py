class Solution:
    def lexGreaterPermutation(self, s: str, target: str) -> str:
        # Store the frequency of every character available in s.
        count = [0] * 26
        for ch in s:
            count[ord(ch) - ord('a')] += 1

        n = len(s)
        matched = 0

        # Match target from left to right for as long as possible.
        while matched < n and count[ord(target[matched]) - ord('a')] > 0:
            # Using the same character keeps the current prefix smallest.
            count[ord(target[matched]) - ord('a')] -= 1
            matched += 1

        # Start where matching failed, or at the last position if all matched.
        start = matched if matched < n else n - 1

        # Move backward because increasing a later position gives a smaller answer.
        for i in range(start, -1, -1):
            # Restore the character if it was previously used to match target.
            if i < matched:
                count[ord(target[i]) - ord('a')] += 1

            # Find the smallest available character greater than target[i].
            bigger = -1
            for ch in range(ord(target[i]) - ord('a') + 1, 26):
                if count[ch] > 0:
                    bigger = ch
                    break

            # Build the answer as soon as this position can be increased.
            if bigger != -1:
                # Consume the character used to make the answer greater.
                count[bigger] -= 1

                # Keep the prefix unchanged and place the larger character.
                answer = target[:i] + chr(ord('a') + bigger)

                # Append all remaining characters in sorted order.
                for ch in range(26):
                    answer += chr(ord('a') + ch) * count[ch]

                return answer

        # No permutation is strictly greater than target.
        return ""