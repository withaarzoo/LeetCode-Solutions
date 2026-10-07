class Solution:
    def removeInvalidParentheses(self, s: str) -> list[str]:
        left = right = 0
        for c in s:
            if c == '(':
                left += 1
            elif c == ')':
                if left > 0:
                    left -= 1
                else:
                    right += 1
        res = set()
        path = []
        def dfs(i, left_rem, right_rem, open_cnt):
            if i == len(s):
                if left_rem == 0 and right_rem == 0 and open_cnt == 0:
                    res.add(''.join(path))
                return
            c = s[i]
            if c != '(' and c != ')':
                path.append(c)
                dfs(i + 1, left_rem, right_rem, open_cnt)
                path.pop()
                return
            if c == '(':
                if left_rem > 0:
                    dfs(i + 1, left_rem - 1, right_rem, open_cnt)
                path.append(c)
                dfs(i + 1, left_rem, right_rem, open_cnt + 1)
                path.pop()
            else:
                if right_rem > 0:
                    dfs(i + 1, left_rem, right_rem - 1, open_cnt)
                if open_cnt > 0:
                    path.append(c)
                    dfs(i + 1, left_rem, right_rem, open_cnt - 1)
                    path.pop()
        dfs(0, left, right, 0)
        return list(res)