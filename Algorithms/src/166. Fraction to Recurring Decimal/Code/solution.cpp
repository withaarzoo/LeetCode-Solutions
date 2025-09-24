class Solution {
public:
    string fractionToDecimal(int numerator, int denominator) {
        // special case
        if (numerator == 0) return "0";
        string res;
        // determine sign
        if ((numerator < 0) ^ (denominator < 0)) res.push_back('-');

        // use long long to avoid overflow for abs(INT_MIN)
        long long n = llabs((long long)numerator);
        long long d = llabs((long long)denominator);

        // integer part
        res += to_string(n / d);
        long long rem = n % d;
        if (rem == 0) return res; // no fractional part

        res.push_back('.');
        unordered_map<long long, int> seen; // remainder -> index in result string

        // simulate long division
        while (rem != 0) {
            // if remainder seen, we found repeating part
            if (seen.find(rem) != seen.end()) {
                int pos = seen[rem];
                res.insert(pos, "(");  // insert '(' at first index where this remainder appeared
                res.push_back(')');    // close parentheses at end
                break;
            }
            // record position of this remainder (position of next fractional digit)
            seen[rem] = res.size();
            rem *= 10;
            int digit = rem / d;
            res.push_back(char('0' + digit));
            rem = rem % d;
        }
        return res;
    }
};
