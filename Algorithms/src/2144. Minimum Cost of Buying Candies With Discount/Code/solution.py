class Solution:
    def minimumCost(self, cost: List[int]) -> int:
        # Sort candies from highest cost to lowest cost
        cost.sort(reverse=True)

        total = 0

        # Every third candy becomes free
        for i in range(len(cost)):
            if i % 3 == 2:
                continue

            total += cost[i]

        return total