from typing import List

class Solution:
    def finalValueAfterOperations(self, operations: List[str]) -> int:
        X = 0  # initial value
        for op in operations:
            # if '+' in op then it's an increment, otherwise decrement
            if '+' in op:
                X += 1
            else:
                X -= 1
        return X
