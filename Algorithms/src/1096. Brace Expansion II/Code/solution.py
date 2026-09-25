class Solution:
    def braceExpansionII(self, expression: str) -> list[str]:
        self.i = 0
        self.expr = expression
        res = self.parse()
        return sorted(res)
    def parse(self):
        res = set()
        cur = {""}
        while self.i < len(self.expr) and self.expr[self.i] != '}':
            if self.expr[self.i] == '{':
                self.i += 1
                nxt = self.parse()
                self.i += 1
                cur = self.product(cur, nxt)
            elif self.expr[self.i] == ',':
                res |= cur
                cur = {""}
                self.i += 1
            else:
                nxt = {self.expr[self.i]}
                self.i += 1
                cur = self.product(cur, nxt)
        res |= cur
        return res
    def product(self, a, b):
        return {x + y for x in a for y in b}