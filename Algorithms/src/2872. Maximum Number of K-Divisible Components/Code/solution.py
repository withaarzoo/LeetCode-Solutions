from typing import List
import sys

class Solution:
    def maxKDivisibleComponents(self, n: int, edges: List[List[int]], values: List[int], k: int) -> int:
        # Increase recursion limit for deep trees
        sys.setrecursionlimit(10**6)
        
        # Build adjacency list
        adj = [[] for _ in range(n)]
        for u, v in edges:
            adj[u].append(v)
            adj[v].append(u)
        
        self.ans = 0
        
        def dfs(u: int, parent: int) -> int:
            """
            Returns: subtree sum % k for node u
            """
            # Start with the current node's value modulo k
            total = values[u] % k
            
            # Process all children
            for v in adj[u]:
                if v == parent:
                    continue      # don't go back up the tree
                child_rem = dfs(v, u)
                total = (total + child_rem) % k
            
            # If this subtree sum is divisible by k, it forms a component
            if total % k == 0:
                self.ans += 1
                return 0         # this subtree is cut off
            return total         # pass remainder up
        
        dfs(0, -1)  # root the tree at node 0
        return self.ans
