class Solution {
    public String fractionToDecimal(int numerator, int denominator) {
        if (numerator == 0) return "0";

        StringBuilder res = new StringBuilder();
        // sign
        if ((numerator < 0) ^ (denominator < 0)) res.append('-');

        // use long to avoid overflow
        long n = Math.abs((long) numerator);
        long d = Math.abs((long) denominator);

        // integer part
        res.append(n / d);
        long rem = n % d;
        if (rem == 0) return res.toString();

        res.append('.');
        Map<Long, Integer> seen = new HashMap<>(); // remainder -> index in res

        while (rem != 0) {
            if (seen.containsKey(rem)) {
                int pos = seen.get(rem);
                res.insert(pos, "(");
                res.append(')');
                break;
            }
            seen.put(rem, res.length());
            rem *= 10;
            res.append(rem / d);
            rem = rem % d;
        }
        return res.toString();
    }
}
