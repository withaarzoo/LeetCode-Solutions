class Solution
{
public:
    bool canBeEqual(string s1, string s2)
    {
        // Store even indexed characters
        string even1 = "";
        string even2 = "";

        // Store odd indexed characters
        string odd1 = "";
        string odd2 = "";

        for (int i = 0; i < 4; i++)
        {
            if (i % 2 == 0)
            {
                even1 += s1[i];
                even2 += s2[i];
            }
            else
            {
                odd1 += s1[i];
                odd2 += s2[i];
            }
        }

        // Sort both even groups
        sort(even1.begin(), even1.end());
        sort(even2.begin(), even2.end());

        // Sort both odd groups
        sort(odd1.begin(), odd1.end());
        sort(odd2.begin(), odd2.end());

        // Both groups must match
        return even1 == even2 && odd1 == odd2;
    }
};