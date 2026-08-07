class Solution:
    def smallestNumber(self, num: str, t: int) -> str:
        req2 = req3 = req5 = req7 = 0
        temp = t
        # Remove primary factors to detect impossible structures natively
        while temp % 2 == 0:
            temp //= 2
            req2 += 1
        while temp % 3 == 0:
            temp //= 3
            req3 += 1
        while temp % 5 == 0:
            temp //= 5
            req5 += 1
        while temp % 7 == 0:
            temp //= 7
            req7 += 1
        if temp > 1: return "-1"

        dp = [[float('inf')] * 40 for _ in range(60)]
        dp[0][0] = 0
        
        # Track shortest path generating required factors explicitly via base digits
        trans = [(1, 0), (0, 1), (2, 0), (1, 1), (3, 0), (0, 2)]
        for i in range(60):
            for j in range(40):
                if dp[i][j] == float('inf'):
                    continue
                for d2, d3 in trans:
                    ni = min(59, i + d2)
                    nj = min(39, j + d3)
                    dp[ni][nj] = min(dp[ni][nj], dp[i][j] + 1)
                    
        # Retroactive updates enforce property solving "minimum digits for AT LEAST i, j factors"
        for i in range(59, -1, -1):
            for j in range(39, -1, -1):
                if i < 59:
                    dp[i][j] = min(dp[i][j], dp[i + 1][j])
                if j < 39:
                    dp[i][j] = min(dp[i][j], dp[i][j + 1])

        F2 = [0, 0, 1, 0, 2, 0, 1, 0, 3, 0]
        F3 = [0, 0, 0, 1, 0, 0, 1, 0, 0, 2]
        F5 = [0, 0, 0, 0, 0, 1, 0, 0, 0, 0]
        F7 = [0, 0, 0, 0, 0, 0, 0, 1, 0, 0]

        n = len(num)
        has_zero = False
        first_zero = n
        for idx, char in enumerate(num):
            if char == '0':
                has_zero = True
                first_zero = idx
                break

        # If zero-free initially, confirm requirement checks directly
        if not has_zero:
            r2, r3, r5, r7 = req2, req3, req5, req7
            for char in num:
                d = int(char)
                r2 = max(0, r2 - F2[d])
                r3 = max(0, r3 - F3[d])
                r5 = max(0, r5 - F5[d])
                r7 = max(0, r7 - F7[d])
            if r2 == 0 and r3 == 0 and r5 == 0 and r7 == 0:
                return num

        # Cap iteration cleanly evaluating prefixes
        limit = min(n - 1, first_zero)
        p2 = p3 = p5 = p7 = 0
        for i in range(limit):
            d = int(num[i])
            p2 += F2[d]
            p3 += F3[d]
            p5 += F5[d]
            p7 += F7[d]

        # Scan inverse array trying incrementally superior rightmost numbers safely
        for i in range(limit, -1, -1):
            start_d = int(num[i]) + 1
            for d in range(start_d, 10):
                n2 = max(0, req2 - p2 - F2[d])
                n3 = max(0, req3 - p3 - F3[d])
                n5 = max(0, req5 - p5 - F5[d])
                n7 = max(0, req7 - p7 - F7[d])
                L = n - 1 - i
                
                # Check spatial fit mapping remainder components seamlessly
                if n7 + n5 + dp[n2][n3] <= L:
                    ans_list = list(num[:i]) + [str(d)]
                    rem2, rem3, rem5, rem7 = n2, n3, n5, n7
                    # Insert minimum variables dynamically left-to-right filling constraints
                    for pos in range(L):
                        for x in range(1, 10):
                            nn2 = max(0, rem2 - F2[x])
                            nn3 = max(0, rem3 - F3[x])
                            nn5 = max(0, rem5 - F5[x])
                            nn7 = max(0, rem7 - F7[x])
                            if nn7 + nn5 + dp[nn2][nn3] <= L - 1 - pos:
                                ans_list.append(str(x))
                                rem2, rem3, rem5, rem7 = nn2, nn3, nn5, nn7
                                break
                    return "".join(ans_list)
            
            if i > 0:
                d = int(num[i - 1])
                p2 -= F2[d]
                p3 -= F3[d]
                p5 -= F5[d]
                p7 -= F7[d]

        # Push dimension length when base string fundamentally lacks sufficient scale
        min_len_needed = req7 + req5 + dp[req2][req3]
        M = max(n + 1, min_len_needed)
        ans_list = []
        rem2, rem3, rem5, rem7 = req2, req3, req5, req7
        
        for pos in range(M):
            for x in range(1, 10):
                nn2 = max(0, rem2 - F2[x])
                nn3 = max(0, rem3 - F3[x])
                nn5 = max(0, rem5 - F5[x])
                nn7 = max(0, rem7 - F7[x])
                if nn7 + nn5 + dp[nn2][nn3] <= M - 1 - pos:
                    ans_list.append(str(x))
                    rem2, rem3, rem5, rem7 = nn2, nn3, nn5, nn7
                    break
        return "".join(ans_list)