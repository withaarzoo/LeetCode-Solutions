class Solution:
    def findTheWinner(self, n: int, k: int) -> int:
        # Call the josephus function to find the winner in 0-indexed form,
        # then add 1 to convert it to 1-indexed as required by the problem.
        return self.josephus(n, k) + 1
    
    def josephus(self, n: int, k: int) -> int:
        # Base case: If there is only one person left, that person is the winner (0-indexed).
        if n == 1:
            return 0
        # Recursive case: Reduce the problem size by one (n-1), find the position of the winner
        # in the smaller problem, adjust by k steps, and take modulo n to wrap around the circle.
        return (self.josephus(n - 1, k) + k) % n
