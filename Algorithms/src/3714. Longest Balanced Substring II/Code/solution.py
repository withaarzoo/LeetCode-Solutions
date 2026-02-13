class Solution:
    def longestBalanced(self, s: str) -> int:
        n = len(s)
        a = b = c = 0
        ans = 0

        # longest single-char run
        run = 0
        prev = ''
        for i, ch in enumerate(s):
            if i == 0 or ch != prev:
                run = 1
            else:
                run += 1
            prev = ch
            ans = max(ans, run)

        # maps: using tuple keys
        map3 = { (0,0) : 0 }         # (b-a, c-a)
        map_ab_c = { (0,0) : 0 }     # (b-a, c)
        map_ac_b = { (0,0) : 0 }     # (c-a, b)
        map_bc_a = { (0,0) : 0 }     # (c-b, a)

        for p in range(1, n+1):
            ch = s[p-1]
            if ch == 'a':
                a += 1
            elif ch == 'b':
                b += 1
            else:
                c += 1

            key3 = (b - a, c - a)
            if key3 in map3:
                ans = max(ans, p - map3[key3])
            else:
                map3[key3] = p

            key_ab_c = (b - a, c)
            if key_ab_c in map_ab_c:
                ans = max(ans, p - map_ab_c[key_ab_c])
            else:
                map_ab_c[key_ab_c] = p

            key_ac_b = (c - a, b)
            if key_ac_b in map_ac_b:
                ans = max(ans, p - map_ac_b[key_ac_b])
            else:
                map_ac_b[key_ac_b] = p

            key_bc_a = (c - b, a)
            if key_bc_a in map_bc_a:
                ans = max(ans, p - map_bc_a[key_bc_a])
            else:
                map_bc_a[key_bc_a] = p

        return ans
