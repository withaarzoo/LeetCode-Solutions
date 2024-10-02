class Solution {
public:
    vector<int> arrayRankTransform(vector<int>& arr) {
        if (arr.empty()) return {};
        
        // Step 1: Copy and sort the array
        vector<int> sorted_arr = arr;
        sort(sorted_arr.begin(), sorted_arr.end());
        
        // Step 2: Create a map to assign ranks
        unordered_map<int, int> rank_map;
        int rank = 1;
        
        // Step 3: Assign ranks to sorted elements
        for (int num : sorted_arr) {
            if (rank_map.find(num) == rank_map.end()) {
                rank_map[num] = rank++;
            }
        }
        
        // Step 4: Replace each element in the original array with its rank
        for (int i = 0; i < arr.size(); ++i) {
            arr[i] = rank_map[arr[i]];
        }
        
        return arr;
    }
};