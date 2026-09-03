class Solution {
public:
    bool uniformArray(vector<int>& nums1) {
        // Store the smallest odd and even values found so far.
        int minOdd = INT_MAX;
        int minEven = INT_MAX;

        // Scan every number once because only the minimum values matter.
        for (int x : nums1) {
            if (x % 2 == 0) {
                // An even value may need to become odd.
                minEven = min(minEven, x);
            } else {
                // The smallest odd value is the best possible value to subtract.
                minOdd = min(minOdd, x);
            }
        }

        // If there is no odd number, every element is already even.
        if (minOdd == INT_MAX) {
            return true;
        }

        // Every even number must be larger than the smallest odd number.
        // Then subtracting minOdd makes every even number odd.
        return minOdd < minEven;
    }
};