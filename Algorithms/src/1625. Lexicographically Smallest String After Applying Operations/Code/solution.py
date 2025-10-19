from collections import deque

class Solution:
    def findLexSmallestString(self, s: str, a: int, b: int) -> str:
        n = len(s)
        seen = set([s])
        q = deque([s])
        ans = s

        while q:
            cur = q.popleft()
            if cur < ans:
                ans = cur

            # Operation 1: add a to odd indices
            arr = list(cur)
            for i in range(1, n, 2):
                arr[i] = str((int(arr[i]) + a) % 10)
            addOp = ''.join(arr)
            if addOp not in seen:
                seen.add(addOp)
                q.append(addOp)

            # Operation 2: rotate right by b
            rotOp = cur[-b:] + cur[:-b]
            if rotOp not in seen:
                seen.add(rotOp)
                q.append(rotOp)

        return ans
