class Solution {
public:
    vector<int> resultArray(vector<int>& nums) {
        // I create two arrays because the problem asks me to distribute nums into them.
        vector<int> arr1, arr2;

        // I reserve space to avoid unnecessary resizing while elements are being added.
        arr1.reserve(nums.size());
        arr2.reserve(nums.size());

        // The first element must always go into arr1.
        arr1.push_back(nums[0]);

        // The second element must always go into arr2.
        arr2.push_back(nums[1]);

        // I process every remaining element according to the last values of both arrays.
        for (int i = 2; i < nums.size(); i++) {
            // If arr1 ends with a larger value, the current number goes into arr1.
            if (arr1.back() > arr2.back()) {
                arr1.push_back(nums[i]);
            } else {
                // Otherwise, including when both values are equal, it goes into arr2.
                arr2.push_back(nums[i]);
            }
        }

        // I create the final result with enough space for all n elements.
        vector<int> result;
        result.reserve(nums.size());

        // I append arr1 first because the required result starts with arr1.
        result.insert(result.end(), arr1.begin(), arr1.end());

        // I append arr2 after arr1 to complete the required concatenation.
        result.insert(result.end(), arr2.begin(), arr2.end());

        // I return the final distributed array.
        return result;
    }
};