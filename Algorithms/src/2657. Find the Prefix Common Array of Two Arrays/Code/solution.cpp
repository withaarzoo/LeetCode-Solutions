class Solution
{
public:
    vector<int> findThePrefixCommonArray(vector<int> &A, vector<int> &B)
    {

        int n = A.size();

        // Frequency array to track how many times
        // each number has appeared so far
        vector<int> freq(n + 1, 0);

        // Result array
        vector<int> ans(n);

        // Stores current count of common numbers
        int common = 0;

        for (int i = 0; i < n; i++)
        {

            // Add current element from A
            freq[A[i]]++;

            // If frequency becomes 2,
            // it means number appeared in both arrays
            if (freq[A[i]] == 2)
            {
                common++;
            }

            // Add current element from B
            freq[B[i]]++;

            // Same check for B
            if (freq[B[i]] == 2)
            {
                common++;
            }

            // Store answer for current prefix
            ans[i] = common;
        }

        return ans;
    }
};