class Solution
{
public:
    vector<int> sortByBits(vector<int> &arr)
    {

        // Sort using custom comparator
        sort(arr.begin(), arr.end(), [](int a, int b)
             {
            
            // Count number of 1 bits
            int bitsA = __builtin_popcount(a);
            int bitsB = __builtin_popcount(b);
            
            // First sort by number of 1 bits
            if (bitsA != bitsB)
                return bitsA < bitsB;
            
            // If equal, sort by actual value
            return a < b; });

        return arr;
    }
};