class Fancy
{
private:
    const long long MOD = 1e9 + 7;

    vector<long long> seq;
    long long mul = 1;
    long long add = 0;

    // Fast exponentiation to compute modular inverse
    long long modPow(long long a, long long b)
    {
        long long res = 1;
        while (b)
        {
            if (b & 1)
                res = (res * a) % MOD;
            a = (a * a) % MOD;
            b >>= 1;
        }
        return res;
    }

public:
    Fancy() {}

    void append(int val)
    {
        long long inv = modPow(mul, MOD - 2); // modular inverse
        long long stored = ((val - add + MOD) % MOD * inv) % MOD;
        seq.push_back(stored);
    }

    void addAll(int inc)
    {
        add = (add + inc) % MOD;
    }

    void multAll(int m)
    {
        mul = (mul * m) % MOD;
        add = (add * m) % MOD;
    }

    int getIndex(int idx)
    {
        if (idx >= seq.size())
            return -1;
        return (seq[idx] * mul % MOD + add) % MOD;
    }
};