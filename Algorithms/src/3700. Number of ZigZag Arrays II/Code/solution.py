class Solution:
    def zigZagArrays(self, n: int, l: int, r: int) -> int:
        MOD = 1000000007

        m = r - l + 1
        sz = 2 * m

        # Matrix multiplication
        def multiply(A, B):
            C = [[0] * sz for _ in range(sz)]

            for i in range(sz):
                for k in range(sz):
                    if A[i][k] == 0:
                        continue

                    cur = A[i][k]

                    for j in range(sz):
                        if B[k][j] == 0:
                            continue

                        C[i][j] = (C[i][j] + cur * B[k][j]) % MOD

            return C

        T = [[0] * sz for _ in range(sz)]

        for x in range(m):

            # up[x] -> down[y]
            for y in range(x + 1, m):
                T[x][m + y] = 1

            # down[x] -> up[y]
            for y in range(x):
                T[m + x][y] = 1

        result = [[0] * sz for _ in range(sz)]
        for i in range(sz):
            result[i][i] = 1

        power = n - 1

        while power:
            if power & 1:
                result = multiply(result, T)

            T = multiply(T, T)
            power >>= 1

        answer = 0

        for i in range(sz):
            row_sum = sum(result[i]) % MOD
            answer = (answer + row_sum) % MOD

        return answer