class Solution
{
public:
    int earliestFinishTime(vector<int> &landStartTime, vector<int> &landDuration,
                           vector<int> &waterStartTime, vector<int> &waterDuration)
    {

        // Store the minimum finishing time found so far
        int ans = INT_MAX;

        int n = landStartTime.size();
        int m = waterStartTime.size();

        // Try every land ride with every water ride
        for (int i = 0; i < n; i++)
        {
            for (int j = 0; j < m; j++)
            {

                // ----------------------------
                // Option 1: Land -> Water
                // ----------------------------

                // Time when land ride finishes
                int landFinish = landStartTime[i] + landDuration[i];

                // Water ride can start only after both:
                // 1. land ride is finished
                // 2. water ride is open
                int waterStart = max(landFinish, waterStartTime[j]);

                // Final finishing time for this order
                int finish1 = waterStart + waterDuration[j];

                // ----------------------------
                // Option 2: Water -> Land
                // ----------------------------

                // Time when water ride finishes
                int waterFinish = waterStartTime[j] + waterDuration[j];

                // Land ride can start only after both:
                // 1. water ride is finished
                // 2. land ride is open
                int landStart = max(waterFinish, landStartTime[i]);

                // Final finishing time for this order
                int finish2 = landStart + landDuration[i];

                // Update answer with the better option
                ans = min(ans, min(finish1, finish2));
            }
        }

        return ans;
    }
};