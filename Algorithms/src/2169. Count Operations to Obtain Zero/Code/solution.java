class Solution {
    public int countOperations(int num1, int num2) {
        long a = num1, b = num2; // widen to avoid edge worries
        int ops = 0;
        while (a > 0 && b > 0) {
            if (a < b) {
                long tmp = a;
                a = b;
                b = tmp;
            } // ensure a >= b
            ops += (int) (a / b); // batch count of subtractions
            a %= b; // remainder left in a
        }
        return ops;
    }
}
