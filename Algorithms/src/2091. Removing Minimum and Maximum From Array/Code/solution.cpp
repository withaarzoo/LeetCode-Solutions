class Solution
{
public:
    int minimumDeletions(vector<int> &nums)
    {
        // Store the total number of elements in the array.
        int n = nums.size();

        // Start by assuming the first element is both minimum and maximum.
        int minIndex = 0;
        int maxIndex = 0;

        // Find the positions of the minimum and maximum elements.
        for (int i = 1; i < n; i++)
        {
            // Update the minimum index if a smaller value is found.
            if (nums[i] < nums[minIndex])
            {
                minIndex = i;
            }

            // Update the maximum index if a larger value is found.
            if (nums[i] > nums[maxIndex])
            {
                maxIndex = i;
            }
        }

        // Remove everything from the front up to the farther special element.
        int removeFromFront = max(minIndex, maxIndex) + 1;

        // Remove everything from the back up to the farther special element.
        int removeFromBack = n - min(minIndex, maxIndex);

        // Remove one special element from the front and the other from the back.
        int removeFromBothSides = min(
            minIndex + 1 + (n - maxIndex),
            maxIndex + 1 + (n - minIndex));

        // Return the minimum deletions among all possible strategies.
        return min(removeFromFront, min(removeFromBack, removeFromBothSides));
    }
};