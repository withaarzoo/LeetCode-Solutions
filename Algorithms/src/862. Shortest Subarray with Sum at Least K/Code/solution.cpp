class Solution
{
public:
    int shortestSubarray(vector<int> &nums, int k)
    {
        int n = nums.size();
        vector<long> prefix(n + 1, 0);

        // Step 1: Compute prefix sums
        for (int i = 0; i < n; ++i)
        {
            prefix[i + 1] = prefix[i] + nums[i];
        }

        deque<int> dq; // Monotonic queue
        int minLength = INT_MAX;

        // Step 2: Process prefix sums
        for (int i = 0; i <= n; ++i)
        {
            // Remove indices from the front if the condition is met
            while (!dq.empty() && prefix[i] - prefix[dq.front()] >= k)
            {
                minLength = min(minLength, i - dq.front());
                dq.pop_front();
            }

            // Maintain monotonicity of the deque
            while (!dq.empty() && prefix[i] <= prefix[dq.back()])
            {
                dq.pop_back();
            }

            // Add the current index to the deque
            dq.push_back(i);
        }

        return minLength == INT_MAX ? -1 : minLength;
    }
};
