class Solution:
    def minJumps(self, nums: List[int]) -> int:

        n = len(nums)

        # Already at destination
        if n == 1:
            return 0

        mx = max(nums)

        # Smallest prime factor array
        spf = list(range(mx + 1))

        # Build sieve
        for i in range(2, int(mx ** 0.5) + 1):

            if spf[i] == i:

                for j in range(i * i, mx + 1, i):

                    if spf[j] == j:
                        spf[j] = i

        # Prime factor -> indices mapping
        mp = {}

        for i, val in enumerate(nums):

            x = val

            used = set()

            # Extract unique prime factors
            while x > 1:

                p = spf[x]

                if p not in used:

                    if p not in mp:
                        mp[p] = []

                    mp[p].append(i)

                    used.add(p)

                x //= p

        # BFS queue
        q = deque([0])

        # Distance array
        dist = [-1] * n

        dist[0] = 0

        while q:

            i = q.popleft()

            steps = dist[i]

            # Reached destination
            if i == n - 1:
                return steps

            # Move left
            if i - 1 >= 0 and dist[i - 1] == -1:
                dist[i - 1] = steps + 1
                q.append(i - 1)

            # Move right
            if i + 1 < n and dist[i + 1] == -1:
                dist[i + 1] = steps + 1
                q.append(i + 1)

            val = nums[i]

            # Teleport possible only if current value is prime
            if val > 1 and spf[val] == val:

                for nxt in mp.get(val, []):

                    if dist[nxt] == -1:
                        dist[nxt] = steps + 1
                        q.append(nxt)

                # Clear to avoid repeated work
                mp[val] = []

        return -1