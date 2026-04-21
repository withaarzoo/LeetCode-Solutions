class Solution:
    def minimumHammingDistance(self, source: List[int], target: List[int], allowedSwaps: List[List[int]]) -> int:
        n = len(source)

        parent = list(range(n))
        rank = [0] * n

        def find(x):
            if parent[x] != x:
                parent[x] = find(parent[x])
            return parent[x]

        def union(a, b):
            pa = find(a)
            pb = find(b)

            if pa == pb:
                return

            if rank[pa] < rank[pb]:
                parent[pa] = pb
            elif rank[pb] < rank[pa]:
                parent[pb] = pa
            else:
                parent[pb] = pa
                rank[pa] += 1

        # Build connected components
        for u, v in allowedSwaps:
            union(u, v)

        from collections import defaultdict, Counter

        groups = defaultdict(list)

        # Group indices by root parent
        for i in range(n):
            groups[find(i)].append(i)

        answer = 0

        # Process each component
        for indices in groups.values():
            freq = Counter()

            # Count source values
            for idx in indices:
                freq[source[idx]] += 1

            # Match target values
            for idx in indices:
                if freq[target[idx]] > 0:
                    freq[target[idx]] -= 1
                else:
                    answer += 1

        return answer