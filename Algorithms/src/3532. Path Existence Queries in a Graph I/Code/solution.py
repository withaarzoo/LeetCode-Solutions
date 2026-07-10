class Solution:
    def pathExistenceQueries(self, n: int, nums: List[int], maxDiff: int, queries: List[List[int]]) -> List[bool]:
        # component[i] stores which connected component node i belongs to.
        component = [0] * n

        # Start with component 0 for the first node.
        component_id = 0

        # Check every gap between two consecutive sorted values.
        for i in range(1, n):
            # A gap larger than maxDiff separates the graph into two parts.
            if nums[i] - nums[i - 1] > maxDiff:
                component_id += 1

            # Store the component of the current node.
            component[i] = component_id

        # Two nodes have a path exactly when their component IDs are equal.
        answer = []
        for u, v in queries:
            answer.append(component[u] == component[v])

        return answer