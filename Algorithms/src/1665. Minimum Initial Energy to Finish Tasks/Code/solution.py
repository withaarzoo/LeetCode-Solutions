class Solution:
    def minimumEffort(self, tasks: List[List[int]]) -> int:

        # Sort by (minimum - actual) in descending order
        tasks.sort(key=lambda x: (x[1] - x[0]), reverse=True)

        answer = 0  # Minimum initial energy
        energy = 0  # Current available energy

        # Process every task
        for actual, minimum in tasks:

            # If current energy is less than required,
            # add the missing amount
            if energy < minimum:

                need = minimum - energy

                answer += need
                energy += need

            # Finish the task
            energy -= actual

        return answer