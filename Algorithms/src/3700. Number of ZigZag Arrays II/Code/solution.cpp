class Solution
{
public:
    static const long long MOD = 1000000007LL;

    // Multiply two matrices
    vector<vector<long long>> multiply(
        const vector<vector<long long>> &A,
        const vector<vector<long long>> &B)
    {
        int sz = A.size();

        vector<vector<long long>> C(sz, vector<long long>(sz, 0));

        for (int i = 0; i < sz; i++)
        {
            for (int k = 0; k < sz; k++)
            {
                if (A[i][k] == 0)
                    continue;

                long long cur = A[i][k];

                for (int j = 0; j < sz; j++)
                {
                    if (B[k][j] == 0)
                        continue;

                    C[i][j] = (C[i][j] + cur * B[k][j]) % MOD;
                }
            }
        }

        return C;
    }

    int zigZagArrays(int n, int l, int r)
    {
        int m = r - l + 1;
        int sz = 2 * m;

        vector<vector<long long>> T(sz, vector<long long>(sz, 0));

        // State layout:
        // [0 ... m-1]       => up[x]
        // [m ... 2m-1]      => down[x]

        for (int x = 0; x < m; x++)
        {

            // up[x] -> down[y] for y > x
            for (int y = x + 1; y < m; y++)
            {
                T[x][m + y] = 1;
            }

            // down[x] -> up[y] for y < x
            for (int y = 0; y < x; y++)
            {
                T[m + x][y] = 1;
            }
        }

        // Identity matrix
        vector<vector<long long>> result(sz, vector<long long>(sz, 0));
        for (int i = 0; i < sz; i++)
        {
            result[i][i] = 1;
        }

        long long power = n - 1;

        // Fast exponentiation
        while (power > 0)
        {
            if (power & 1)
            {
                result = multiply(result, T);
            }

            T = multiply(T, T);
            power >>= 1;
        }

        // Initial vector is all ones
        vector<long long> initial(sz, 1);

        long long answer = 0;

        // Sum of all entries in result * initial
        for (int i = 0; i < sz; i++)
        {
            long long rowSum = 0;

            for (int j = 0; j < sz; j++)
            {
                rowSum = (rowSum + result[i][j]) % MOD;
            }

            answer = (answer + rowSum) % MOD;
        }

        return (int)answer;
    }
};