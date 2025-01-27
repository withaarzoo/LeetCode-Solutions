class Solution:
    def checkIfPrerequisite(self, numCourses: int, prerequisites: List[List[int]], queries: List[List[int]]) -> List[bool]:
        # Initialize the graph
        graph = [[False] * numCourses for _ in range(numCourses)]

        # Build the direct edges from prerequisites
        for u, v in prerequisites:
            graph[u][v] = True

        # Floyd-Warshall to compute transitive closure
        for k in range(numCourses):
            for i in range(numCourses):
                for j in range(numCourses):
                    if graph[i][k] and graph[k][j]:
                        graph[i][j] = True

        # Answer the queries
        return [graph[u][v] for u, v in queries]
