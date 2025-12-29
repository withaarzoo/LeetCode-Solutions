class Solution:
    def pyramidTransition(self, bottom: str, allowed: List[str]) -> bool:
        from collections import defaultdict

        rules = defaultdict(set)
        for a, b, c in allowed:
            rules[a + b].add(c)

        bad = set()

        def dfs(row, idx, nxt):
            if len(row) == 1:
                return True

            if idx == len(row) - 1:
                if nxt in bad:
                    return False
                ok = dfs(nxt, 0, "")
                if not ok:
                    bad.add(nxt)
                return ok

            key = row[idx:idx + 2]
            for c in rules[key]:
                if dfs(row, idx + 1, nxt + c):
                    return True
            return False

        return dfs(bottom, 0, "")
