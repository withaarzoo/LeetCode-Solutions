class Solution:
    def findAllPeople(self, n, meetings, firstPerson):
        meetings.sort(key=lambda x: x[2])

        parent = list(range(n))
        knows = [False] * n
        knows[0] = knows[firstPerson] = True

        def find(x):
            if parent[x] != x:
                parent[x] = find(parent[x])
            return parent[x]

        def union(x, y):
            x, y = find(x), find(y)
            if x != y:
                parent[y] = x

        i = 0
        while i < len(meetings):
            time = meetings[i][2]
            people = []

            j = i
            while j < len(meetings) and meetings[j][2] == time:
                union(meetings[j][0], meetings[j][1])
                people.extend([meetings[j][0], meetings[j][1]])
                j += 1

            good = set(find(p) for p in people if knows[p])

            for p in people:
                if find(p) in good:
                    knows[p] = True
                else:
                    parent[p] = p

            i = j

        return [i for i in range(n) if knows[i]]
