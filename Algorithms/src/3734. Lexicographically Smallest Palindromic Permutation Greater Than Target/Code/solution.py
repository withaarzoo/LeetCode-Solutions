class Solution:

    def build_palindrome(self, half: str, middle: str) -> str:
        # The right half is forced to be the reverse of the left half.
        return half + middle + half[::-1]

    def smallest_greater_or_equal(self, original_count, target_half: str) -> str:
        # Work on a copy because the search changes character frequencies.
        count = original_count[:]
        k = len(target_half)
        matched = 0

        # Match the target prefix exactly for as long as possible.
        while matched < k and count[ord(target_half[matched]) - ord('a')] > 0:
            count[ord(target_half[matched]) - ord('a')] -= 1
            matched += 1

        # The complete target prefix itself can be formed.
        if matched == k:
            return target_half

        # Move backward and try to increase the rightmost possible position.
        for pos in range(matched, -1, -1):
            # Restore a character when backtracking over a matched position.
            if pos < matched:
                count[ord(target_half[pos]) - ord('a')] += 1

            current = ord(target_half[pos]) - ord('a')

            # Choose the smallest available character strictly larger than target_half[pos].
            for c in range(current + 1, 26):
                if count[c] == 0:
                    continue

                result = target_half[:pos] + chr(ord('a') + c)
                count[c] -= 1

                # Fill every remaining position in ascending order.
                for ch in range(26):
                    result += chr(ord('a') + ch) * count[ch]

                return result

        # No permutation can reach or exceed target_half.
        return ""

    def next_permutation(self, chars) -> bool:
        # Find the rightmost position that can be increased.
        pivot = len(chars) - 2

        while pivot >= 0 and chars[pivot] >= chars[pivot + 1]:
            pivot -= 1

        # The sequence is already the largest lexicographical permutation.
        if pivot < 0:
            return False

        swap_pos = len(chars) - 1

        # Find the smallest character larger than chars[pivot].
        while chars[swap_pos] <= chars[pivot]:
            swap_pos -= 1

        # Swap the pivot with that character.
        chars[pivot], chars[swap_pos] = chars[swap_pos], chars[pivot]

        # Reverse the suffix to make the next permutation as small as possible.
        chars[pivot + 1:] = reversed(chars[pivot + 1:])

        return True

    def lexPalindromicPermutation(self, s: str, target: str) -> str:
        # Count the frequency of every lowercase English letter.
        frequency = [0] * 26

        for ch in s:
            frequency[ord(ch) - ord('a')] += 1

        middle = ""  # Stores the only possible middle character.
        odd_count = 0

        # A valid palindrome can have at most one odd-frequency character.
        for c in range(26):
            if frequency[c] % 2 == 1:
                odd_count += 1
                middle = chr(ord('a') + c)

        if odd_count > 1:
            return ""

        # Only one character from every pair belongs to the first half.
        half_count = [count // 2 for count in frequency]
        k = len(s) // 2
        target_half = target[:k]

        # Find the smallest first-half permutation that is at least target_half.
        half = self.smallest_greater_or_equal(half_count, target_half)

        if not half and k > 0:
            return ""

        # Build the smallest candidate using that first half.
        candidate = self.build_palindrome(half, middle)

        if candidate > target:
            return candidate

        # If equality was not enough, move to the next first-half permutation.
        chars = list(half)

        if not self.next_permutation(chars):
            return ""

        # This is the smallest palindrome with a strictly larger first half.
        return self.build_palindrome("".join(chars), middle)