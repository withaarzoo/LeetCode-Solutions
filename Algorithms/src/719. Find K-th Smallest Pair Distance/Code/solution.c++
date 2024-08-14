#include <vector>
#include <algorithm>

class Solution
{
public:
    // This function counts the number of pairs (i, j) such that the difference between nums[j] and nums[i] is less than or equal to 'mid'.
    int countPairs(const std::vector<int> &nums, int mid)
    {
        int count = 0; // Initialize the count of pairs to zero
        int j = 0;     // Initialize the second pointer j to the beginning

        // Loop through each element with the first pointer i
        for (int i = 0; i < nums.size(); ++i)
        {
            // Move the second pointer j as long as the difference between nums[j] and nums[i] is less than or equal to 'mid'
            while (j < nums.size() && nums[j] - nums[i] <= mid)
            {
                ++j;
            }
            // For the current i, all elements between i and j-1 have a difference <= 'mid' with nums[i]
            // Hence, we add (j - i - 1) to the count, since j points to the first element that exceeds 'mid' difference
            count += j - i - 1;
        }

        return count; // Return the total count of such pairs
    }

    // This function finds the k-th smallest distance pair in the sorted array 'nums'
    int smallestDistancePair(std::vector<int> &nums, int k)
    {
        std::sort(nums.begin(), nums.end()); // Sort the array in non-decreasing order

        // Set the initial search range for binary search
        int low = 0;                           // The smallest possible distance
        int high = nums.back() - nums.front(); // The largest possible distance

        // Perform binary search to find the smallest distance that has at least 'k' pairs
        while (low < high)
        {
            int mid = (low + high) / 2; // Calculate the middle distance

            // Check how many pairs have a distance <= 'mid'
            if (countPairs(nums, mid) >= k)
            {
                high = mid; // If there are at least 'k' pairs, search in the lower half
            }
            else
            {
                low = mid + 1; // If there are fewer than 'k' pairs, search in the upper half
            }
        }

        return low; // 'low' will be the smallest distance that has at least 'k' pairs
    }
};
