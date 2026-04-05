class Solution:
    def judgeCircle(self, moves: str) -> bool:
        # x represents left/right position
        # y represents up/down position
        x, y = 0, 0

        # Traverse every move
        for move in moves:
            if move == 'U':
                y += 1  # Move up
            elif move == 'D':
                y -= 1  # Move down
            elif move == 'R':
                x += 1  # Move right
            elif move == 'L':
                x -= 1  # Move left

        # Robot returns to origin only if both coordinates are 0
        return x == 0 and y == 0