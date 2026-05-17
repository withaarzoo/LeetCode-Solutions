class Solution:
    def canReach(self, arr: List[int], start: int) -> bool:
        
        # Visited array to avoid infinite loops
        visited = [False] * len(arr)

        # DFS function
        def dfs(index):

            # Invalid index
            if index < 0 or index >= len(arr):
                return False

            # Skip already visited indexes
            if visited[index]:
                return False

            # Found value 0
            if arr[index] == 0:
                return True

            # Mark current index as visited
            visited[index] = True

            # Explore both directions
            return dfs(index + arr[index]) or dfs(index - arr[index])

        # Start DFS from given index
        return dfs(start)