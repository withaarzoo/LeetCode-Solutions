class Solution:
    def evaluate(self, s: str, knowledge: list[list[str]]) -> str:
        mp = {k: v for k, v in knowledge}
        res = []
        n = len(s)
        i = 0
        while i < n:
            if s[i] == '(':
                j = i + 1
                while s[j] != ')':
                    j += 1
                key = s[i + 1:j]
                res.append(mp.get(key, "?"))
                i = j + 1
            else:
                res.append(s[i])
                i += 1
        return "".join(res)