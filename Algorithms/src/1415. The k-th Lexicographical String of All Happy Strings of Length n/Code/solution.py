class Solution:
    def getHappyString(self, n: int, k: int) -> str:
        
        self.count = 0
        self.result = ""

        def dfs(curr):

            if self.result:
                return

            if len(curr) == n:
                self.count += 1
                if self.count == k:
                    self.result = curr
                return

            for c in ['a','b','c']:

                if curr and curr[-1] == c:
                    continue

                dfs(curr + c)

        dfs("")
        return self.result