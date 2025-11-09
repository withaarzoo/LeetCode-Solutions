class Solution
{
public:
    int countOperations(int num1, int num2)
    {
        long long a = num1, b = num2; // use long long to be extra safe, though ints are fine here
        int ops = 0;
        while (a > 0 && b > 0)
        {
            if (a < b)
                std::swap(a, b); // ensure a >= b
            ops += a / b;        // how many times we'd subtract b from a
            a %= b;              // remaining part after removing (a/b) copies of b
        }
        return ops;
    }
};
