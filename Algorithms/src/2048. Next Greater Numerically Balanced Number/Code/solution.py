class Solution:
    def isBalanced(self, x: int) -> bool:
        cnt = [0]*10
        t = x
        while t > 0:
            cnt[t % 10] += 1
            t //= 10
        # digit 0 cannot appear
        if cnt[0] > 0:
            return False
        for d in range(1, 10):
            if cnt[d] != 0 and cnt[d] != d:
                return False
        return True

    def nextBeautifulNumber(self, n: int) -> int:
        x = n + 1
        while True:
            if self.isBalanced(x):
                return x
            x += 1
