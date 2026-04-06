class Solution:
    def robotSim(self, commands: List[int], obstacles: List[List[int]]) -> int:
        # Store all obstacles in a set for O(1) lookup
        obstacle_set = set()

        for x, y in obstacles:
            obstacle_set.add((x, y))

        # Directions: North, East, South, West
        dx = [0, 1, 0, -1]
        dy = [1, 0, -1, 0]

        direction = 0  # Start facing North
        x, y = 0, 0
        max_distance = 0

        for command in commands:
            # Turn right
            if command == -1:
                direction = (direction + 1) % 4

            # Turn left
            elif command == -2:
                direction = (direction + 3) % 4

            # Move forward
            else:
                for _ in range(command):
                    next_x = x + dx[direction]
                    next_y = y + dy[direction]

                    # Stop if obstacle exists
                    if (next_x, next_y) in obstacle_set:
                        break

                    x, y = next_x, next_y

                    max_distance = max(max_distance, x * x + y * y)

        return max_distance