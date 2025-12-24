class Solution
{
public:
    int minimumBoxes(vector<int> &apple, vector<int> &capacity)
    {
        // Step 1: Calculate total apples
        int totalApples = 0;
        for (int a : apple)
        {
            totalApples += a;
        }

        // Step 2: Sort capacities in descending order
        sort(capacity.begin(), capacity.end(), greater<int>());

        // Step 3: Pick boxes until capacity is enough
        int usedCapacity = 0;
        int boxes = 0;

        for (int cap : capacity)
        {
            usedCapacity += cap;
            boxes++;
            if (usedCapacity >= totalApples)
            {
                return boxes;
            }
        }

        return boxes; // guaranteed possible as per constraints
    }
};
