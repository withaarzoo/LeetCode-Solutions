class Solution:
    def processStr(self, s: str, k: int) -> str:
        n = len(s)

        # lengths[i] = length after processing s[i]
        lengths = [0] * n
        cur_len = 0

        for i, ch in enumerate(s):
            if 'a' <= ch <= 'z':
                # Append a character
                cur_len += 1
            elif ch == '*':
                # Remove last character if present
                if cur_len > 0:
                    cur_len -= 1
            elif ch == '#':
                # Duplicate the whole string
                cur_len *= 2
            else:
                # '%' only reverses, length stays same
                pass

            lengths[i] = cur_len

        # k is outside the final string
        if k >= cur_len:
            return '.'

        # Undo operations from right to left
        for i in range(n - 1, -1, -1):
            ch = s[i]
            before = 0 if i == 0 else lengths[i - 1]

            if 'a' <= ch <= 'z':
                # This letter was appended at index "before"
                if k == before:
                    return ch
            elif ch == '#':
                # Undo T + T
                if before > 0:
                    k %= before
            elif ch == '%':
                # Undo reverse
                k = before - 1 - k
            else:
                # '*' does not change surviving indices
                pass

        return '.'