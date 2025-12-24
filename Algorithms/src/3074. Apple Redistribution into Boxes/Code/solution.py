class Solution:
    def minimumBoxes(self, apple: List[int], capacity: List[int]) -> int:
        # Step 1: Total apples
        total_apples = sum(apple)

        # Step 2: Sort capacities in descending order
        capacity.sort(reverse=True)

        # Step 3: Pick boxes greedily
        used_capacity = 0
        boxes = 0

        for cap in capacity:
            used_capacity += cap
            boxes += 1
            if used_capacity >= total_apples:
                return boxes

        return boxes
