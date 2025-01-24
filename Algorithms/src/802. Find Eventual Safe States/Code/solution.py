from collections import deque

class Solution:
    def eventualSafeNodes(self, graph: List[List[int]]) -> List[int]:
        n = len(graph)
        reversedGraph = [[] for _ in range(n)]
        inDegree = [0] * n
        
        # Reverse the graph and calculate in-degree
        for i in range(n):
            for neighbor in graph[i]:
                reversedGraph[neighbor].append(i)
                inDegree[i] += 1
        
        # Find all terminal nodes
        queue = deque([i for i in range(n) if inDegree[i] == 0])
        
        # Topological sorting to find safe nodes
        safeNodes = []
        while queue:
            node = queue.popleft()
            safeNodes.append(node)
            
            for neighbor in reversedGraph[node]:
                inDegree[neighbor] -= 1
                if inDegree[neighbor] == 0:
                    queue.append(neighbor)
        
        return sorted(safeNodes)
