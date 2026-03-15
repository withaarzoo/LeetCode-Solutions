class Fancy:

    MOD = 10**9 + 7

    def __init__(self):
        self.seq = []
        self.mul = 1
        self.add = 0

    def mod_pow(self, a, b):
        res = 1
        while b:
            if b & 1:
                res = res * a % self.MOD
            a = a * a % self.MOD
            b >>= 1
        return res

    def append(self, val: int) -> None:
        inv = self.mod_pow(self.mul, self.MOD - 2)
        stored = ((val - self.add) % self.MOD) * inv % self.MOD
        self.seq.append(stored)

    def addAll(self, inc: int) -> None:
        self.add = (self.add + inc) % self.MOD

    def multAll(self, m: int) -> None:
        self.mul = self.mul * m % self.MOD
        self.add = self.add * m % self.MOD

    def getIndex(self, idx: int) -> int:
        if idx >= len(self.seq):
            return -1
        return (self.seq[idx] * self.mul + self.add) % self.MOD