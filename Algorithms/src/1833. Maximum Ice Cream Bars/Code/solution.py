class Solution:
    def maxIceCream(self, costs: List[int], coins: int) -> int:
        
        # Maximum possible cost according to constraints
        MAX_COST = 100000

        # Frequency array to count occurrences of each cost
        freq = [0] * (MAX_COST + 1)

        # Count every ice cream cost
        for cost in costs:
            freq[cost] += 1

        # Stores total purchased bars
        answer = 0

        # Process costs from smallest to largest
        for cost in range(1, MAX_COST + 1):

            # Skip unavailable costs
            if freq[cost] == 0:
                continue

            # Maximum bars affordable at current cost
            can_buy = min(freq[cost], coins // cost)

            # Increase purchased count
            answer += can_buy

            # Deduct spent coins
            coins -= can_buy * cost

        return answer