class Solution:
    def minOperations(self, boxes: str) -> List[int]:
        n = len(boxes)
        answer = [0] * n

        # Left-to-right pass
        balls = 0
        operations = 0
        for i in range(n):
            answer[i] += operations
            balls += int(boxes[i]) # Count balls
            operations += balls   # Add the current number of balls to operations

        # Right-to-left pass
        balls = 0
        operations = 0
        for i in range(n - 1, -1, -1):
            answer[i] += operations
            balls += int(boxes[i]) # Count balls
            operations += balls    # Add the current number of balls to operations

        return answer
