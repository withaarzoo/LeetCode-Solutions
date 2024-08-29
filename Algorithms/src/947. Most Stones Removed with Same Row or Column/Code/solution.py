class Solution:
    # Depth-First Search (DFS) function to explore all connected nodes
    def dfs(self, node, adj, visited):
        # Mark the current node as visited
        visited.add(node)
        
        # Explore all neighbors (connected nodes) of the current node
        for neighbor in adj[node]:
            # If the neighbor hasn't been visited yet, recursively visit it
            if neighbor not in visited:
                self.dfs(neighbor, adj, visited)
    
    def removeStones(self, stones: List[List[int]]) -> int:
        # Number of stones (nodes) in the input
        n = len(stones)
        
        # Create an adjacency list to represent the graph
        adj = [[] for _ in range(n)]
        
        # Build the graph by connecting stones that are in the same row or column
        for i in range(n):
            for j in range(i + 1, n):
                # If stones[i] and stones[j] share the same row or column
                if stones[i][0] == stones[j][0] or stones[i][1] == stones[j][1]:
                    # Add an edge between node i and node j
                    adj[i].append(j)
                    adj[j].append(i)
        
        # A set to keep track of visited nodes
        visited = set()
        
        # Initialize the count of connected components
        numComponents = 0

        # Traverse each stone (node) to count connected components
        for i in range(n):
            # If the stone (node) has not been visited, it's a new component
            if i not in visited:
                # Perform DFS starting from this node
                self.dfs(i, adj, visited)
                # Increment the number of connected components
                numComponents += 1

        # The result is the total number of stones minus the number of components
        # This gives the maximum number of stones that can be removed
        return n - numComponents
