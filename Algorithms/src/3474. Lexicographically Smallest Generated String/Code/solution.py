class Solution:
    def generateString(self, str1: str, str2: str) -> str:
        n = len(str1)
        m = len(str2)
        length = n + m - 1

        ans = ['?'] * length
        fixed = [False] * length

        # Apply all 'T' constraints
        for i in range(n):
            if str1[i] == 'T':
                for j in range(m):
                    pos = i + j

                    if ans[pos] != '?' and ans[pos] != str2[j]:
                        return ""

                    ans[pos] = str2[j]
                    fixed[pos] = True

        # Fill remaining positions with 'a'
        for i in range(length):
            if ans[i] == '?':
                ans[i] = 'a'

        # Process all 'F' constraints
        for i in range(n):
            if str1[i] == 'F':
                same = True

                for j in range(m):
                    if ans[i + j] != str2[j]:
                        same = False
                        break

                if not same:
                    continue

                changed = False

                # Try changing from right to left
                for j in range(m - 1, -1, -1):
                    pos = i + j

                    if fixed[pos]:
                        continue

                    for c in range(ord('a'), ord('z') + 1):
                        ch = chr(c)

                        if ch != ans[pos] and ch != str2[j]:
                            ans[pos] = ch
                            changed = True
                            break

                    if changed:
                        break

                if not changed:
                    return ""

        return ''.join(ans)