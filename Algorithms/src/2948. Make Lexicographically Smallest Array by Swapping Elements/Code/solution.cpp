class Solution {
public:
    vector<int> lexicographicallySmallestArray(vector<int>& nums, int limit) {
        int n = nums.size(); // Store the size so I can use it throughout the solution.
        
        // Store every value together with its original position.
        vector<pair<int, int>> elements;
        for (int i = 0; i < n; i++) {
            elements.push_back({nums[i], i});
        }
        
        // Sort by value so connected values appear next to each other.
        sort(elements.begin(), elements.end());
        
        // This will store the lexicographically smallest possible result.
        vector<int> answer(n);
        
        int start = 0; // Marks the first element of the current connected group.
        
        while (start < n) {
            int end = start; // Expand this group as far as valid swaps allow.
            
            // Consecutive sorted values belong to the same group if their
            // difference is at most limit, which means they are connected.
            while (end + 1 < n &&
                   (long long)elements[end + 1].first - elements[end].first <= limit) {
                end++;
            }
            
            vector<int> indices; // Store the original positions of this group.
            
            // Collect every original index that belongs to the current group.
            for (int i = start; i <= end; i++) {
                indices.push_back(elements[i].second);
            }
            
            // Sort positions so smaller values can be placed earlier.
            sort(indices.begin(), indices.end());
            
            // Values are already sorted in elements[start...end].
            // Assign them to the sorted original indices.
            for (int i = 0; i < (int)indices.size(); i++) {
                answer[indices[i]] = elements[start + i].first;
            }
            
            // Start processing the next connected group.
            start = end + 1;
        }
        
        return answer; // Return the lexicographically smallest arrangement.
    }
};