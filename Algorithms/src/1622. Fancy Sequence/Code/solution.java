class Fancy {

    static final long MOD = 1000000007;

    ArrayList<Long> seq = new ArrayList<>();
    long mul = 1;
    long add = 0;

    private long modPow(long a, long b) {
        long res = 1;
        while (b > 0) {
            if ((b & 1) == 1)
                res = (res * a) % MOD;
            a = (a * a) % MOD;
            b >>= 1;
        }
        return res;
    }

    public Fancy() {
    }

    public void append(int val) {
        long inv = modPow(mul, MOD - 2);
        long stored = ((val - add + MOD) % MOD * inv) % MOD;
        seq.add(stored);
    }

    public void addAll(int inc) {
        add = (add + inc) % MOD;
    }

    public void multAll(int m) {
        mul = (mul * m) % MOD;
        add = (add * m) % MOD;
    }

    public int getIndex(int idx) {
        if (idx >= seq.size())
            return -1;
        return (int) ((seq.get(idx) * mul % MOD + add) % MOD);
    }
}