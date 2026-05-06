class Solution:
    def rotateTheBox(self, boxGrid: List[List[str]]) -> List[List[str]]:

        m = len(boxGrid)
        n = len(boxGrid[0])

        # Process every row
        for row in range(m):

            # Rightmost empty position
            empty_col = n - 1

            # Traverse from right to left
            for col in range(n - 1, -1, -1):

                # Obstacle blocks stones
                if boxGrid[row][col] == '*':

                    # Reset valid position
                    empty_col = col - 1

                # Stone found
                elif boxGrid[row][col] == '#':

                    # Remove stone from current position
                    boxGrid[row][col] = '.'

                    # Move stone to valid position
                    boxGrid[row][empty_col] = '#'

                    # Next stone goes left
                    empty_col -= 1

        # Create rotated matrix
        rotated = [['.'] * m for _ in range(n)]

        # Rotate clockwise
        for i in range(m):
            for j in range(n):

                rotated[j][m - 1 - i] = boxGrid[i][j]

        return rotated