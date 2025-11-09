class Solution:
    def countOperations(self, num1: int, num2: int) -> int:
        a, b = num1, num2
        ops = 0
        while a > 0 and b > 0:
            if a < b:
                a, b = b, a            # ensure a >= b
            ops += a // b               # count batched subtractions
            a %= b                      # remainder becomes new a
        return ops
