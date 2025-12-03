class Solution
{
public:
    using ll = long long;

    // small offset so we can pack signed components into one int
    static constexpr int SHIFT = 3000; // |ux|,|uy|,|dx|,|dy| <= 2000

    // encode (a,b) into one 32-bit key
    int encodePair(int a, int b)
    {
        // each part fits in 13 bits after +SHIFT
        return ((a + SHIFT) << 13) ^ (b + SHIFT);
    }

    // gcd for non-negative ints
    int myGcd(int a, int b)
    {
        if (a < 0)
            a = -a;
        if (b < 0)
            b = -b;
        while (b != 0)
        {
            int t = a % b;
            a = b;
            b = t;
        }
        return a;
    }

    // generic counting function described in the approach
    ll countPairs(const unordered_map<int, unordered_map<int, int>> &mp)
    {
        ll res = 0;
        for (const auto &outer : mp)
        {
            const auto &inner = outer.second;
            ll sum = 0, sumSq = 0;
            for (auto &kv : inner)
            {
                ll c = kv.second;
                sum += c;
                sumSq += c * c;
            }
            res += (sum * sum - sumSq) / 2;
        }
        return res;
    }

    int countTrapezoids(vector<vector<int>> &points)
    {
        int n = (int)points.size();

        // slope -> (lineId -> count)
        unordered_map<int, unordered_map<int, int>> bySlope;
        // full vector -> (lineId -> count)
        unordered_map<int, unordered_map<int, int>> byVector;

        bySlope.reserve(n * n);
        byVector.reserve(n * n);

        for (int i = 0; i < n; ++i)
        {
            int x1 = points[i][0];
            int y1 = points[i][1];
            for (int j = i + 1; j < n; ++j)
            {
                int x2 = points[j][0];
                int y2 = points[j][1];

                int dx = x2 - x1;
                int dy = y2 - y1;

                // fix direction sign to be consistent
                if (dx < 0 || (dx == 0 && dy < 0))
                {
                    dx = -dx;
                    dy = -dy;
                }

                int g = myGcd(dx, dy);
                int ux = dx / g; // reduced dir
                int uy = dy / g;

                // line identifier for this slope: ux*y - uy*x
                int lineId = ux * y1 - uy * x1;

                int slopeKey = encodePair(ux, uy);
                int vectorKey = encodePair(dx, dy);

                bySlope[slopeKey][lineId] += 1;
                byVector[vectorKey][lineId] += 1;
            }
        }

        ll withParallel = countPairs(bySlope);
        ll parallelogramTwice = countPairs(byVector);

        ll ans = withParallel - parallelogramTwice / 2;
        return (int)ans;
    }
};
