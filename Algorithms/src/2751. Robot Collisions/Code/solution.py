class Solution:
    def survivedRobotsHealths(self, positions: List[int], healths: List[int], directions: str) -> List[int]:
        n = len(positions)

        # Store robot indices
        indices = list(range(n))

        # Sort indices based on positions
        indices.sort(key=lambda i: positions[i])

        # Stack stores indices of robots moving right
        stack = []

        for idx in indices:
            # Robot moving right
            if directions[idx] == 'R':
                stack.append(idx)
            else:
                # Robot moving left
                while stack and healths[idx] > 0:
                    top_idx = stack[-1]

                    if healths[top_idx] < healths[idx]:
                        stack.pop()
                        healths[idx] -= 1
                        healths[top_idx] = 0

                    elif healths[top_idx] == healths[idx]:
                        stack.pop()
                        healths[top_idx] = 0
                        healths[idx] = 0

                    else:
                        healths[top_idx] -= 1
                        healths[idx] = 0

        # Collect surviving robots in original order
        result = []

        for health in healths:
            if health > 0:
                result.append(health)

        return result