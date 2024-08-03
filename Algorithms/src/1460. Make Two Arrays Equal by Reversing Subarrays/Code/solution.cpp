#include <vector>
#include <algorithm> // For std::sort

class Solution
{
public:
    bool canBeEqual(std::vector<int> &target, std::vector<int> &arr)
    {
        // Step 1: Sort the 'target' array
        // Sorting helps in comparing the arrays element by element in a straightforward manner
        std::sort(target.begin(), target.end());

        // Step 2: Sort the 'arr' array
        // After sorting both arrays, they should have the same elements in the same order if they can be made equal
        std::sort(arr.begin(), arr.end());

        // Step 3: Compare the sorted arrays
        // If the sorted arrays are identical, it means the original arrays can be made equal by reordering
        return target == arr;
    }
};
