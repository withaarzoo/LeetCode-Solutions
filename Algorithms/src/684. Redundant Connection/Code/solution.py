from typing import List

class Solution:
    def findRedundantConnection(self, edges: List[List[int]]) -> List[int]:
        parent = list(range(len(edges) + 1))
        rank = [0] * (len(edges) + 1)

        def find(node):
            if parent[node] != node:
                parent[node] = find(parent[node])  # Path compression
            return parent[node]

        def union(u, v):
            rootU, rootV = find(u), find(v)
            if rootU == rootV:
                return False  # Cycle detected
            if rank[rootU] > rank[rootV]:
                parent[rootV] = rootU
            elif rank[rootU] < rank[rootV]:
                parent[rootU] = rootV
            else:
                parent[rootV] = rootU
                rank[rootU] += 1
            return True

        for u, v in edges:
            if not union(u, v):
                return [u, v]
        return []
