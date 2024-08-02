#include <vector>
#include <algorithm>

using namespace std;

class Solution
{
public:
    int minSwaps(vector<int> &nums)
    {
        // Step 1: Count the total number of 1's in the array.
        int totalOnes = count(nums.begin(), nums.end(), 1);

        // Step 2: If there are no 1's in the array, no swaps are needed.
        if (totalOnes == 0)
            return 0;

        // Step 3: Initialize necessary variables.
        int n = nums.size();         // Size of the array
        int maxOnesInWindow = 0;     // Maximum number of 1's in any window of size 'totalOnes'
        int currentOnesInWindow = 0; // Number of 1's in the current window

        // Step 4: Calculate the number of 1's in the first window of size 'totalOnes'.
        for (int i = 0; i < totalOnes; i++)
        {
            currentOnesInWindow += nums[i];
        }

        // Step 5: Initialize maxOnesInWindow with the number of 1's in the first window.
        maxOnesInWindow = currentOnesInWindow;

        // Step 6: Use a sliding window to check every possible window of size 'totalOnes'.
        for (int i = 1; i < n; i++)
        {
            // Subtract the element that is sliding out of the window from the left.
            currentOnesInWindow -= nums[i - 1];

            // Add the new element that is sliding into the window from the right.
            currentOnesInWindow += nums[(i + totalOnes - 1) % n];

            // Update maxOnesInWindow with the maximum number of 1's found in any window.
            maxOnesInWindow = max(maxOnesInWindow, currentOnesInWindow);
        }

        // Step 7: The minimum swaps needed is the total number of 1's minus the maximum number
        // of 1's found in any window. This gives us the number of 0's that need to be swapped.
        return totalOnes - maxOnesInWindow;
    }
};
