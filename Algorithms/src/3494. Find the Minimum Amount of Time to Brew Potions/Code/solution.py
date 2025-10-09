class Solution:
    def minTime(self, skill: List[int], mana: List[int]) -> int:
        n = len(skill)
        m = len(mana)
        if m == 0:
            return 0

        # prefix sums of skills (Python int is arbitrary precision)
        pref = [0] * n
        for i in range(n):
            pref[i] = skill[i] + (pref[i-1] if i > 0 else 0)

        S = 0  # S_j (start time of current potion on wizard 0)
        # For each transition from potion j-1 -> j compute the max expression
        for j in range(1, m):
            prev = mana[j-1]
            cur  = mana[j]
            best = -10**30  # sufficiently small
            # local variable binding for speed
            pref_local = pref
            for i in range(n):
                prev_pref = pref_local[i-1] if i > 0 else 0
                cand = pref_local[i] * prev - prev_pref * cur
                if cand > best:
                    best = cand
            S += best

        ans = S + pref[-1] * mana[-1]
        return ans
