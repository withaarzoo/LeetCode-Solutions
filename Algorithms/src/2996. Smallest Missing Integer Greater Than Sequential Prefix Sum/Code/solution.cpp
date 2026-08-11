class Solution {
public:
    int missingInteger(vector<int>& nums) {
        // I start with the first element because nums[0] is always
        // considered part of the sequential prefix.
        int sum = nums[0];

        // I scan the array from the second element onward.
        for (int i = 1; i < nums.size(); i++) {
            // If the current value is exactly one greater than
            // the previous value, the sequential prefix continues.
            if (nums[i] == nums[i - 1] + 1) {
                sum += nums[i];
            } else {
                // The sequence breaks here, so the prefix ends.
                break;
            }
        }

        // I store every value so I can quickly check whether
        // a candidate number exists anywhere in the array.
        unordered_set<int> seen(nums.begin(), nums.end());

        // I start checking from the sum of the longest sequential prefix.
        int answer = sum;

        // If the current number exists, it cannot be the answer,
        // so I keep moving to the next integer.
        while (seen.count(answer)) {
            answer++;
        }

        // The first missing value is the required answer.
        return answer;
    }
};