class Solution
{
public:
    int maximumProduct(vector<int> &nums)
    {
        // Store the three largest numbers
        int max1 = INT_MIN, max2 = INT_MIN, max3 = INT_MIN;

        // Store the two smallest numbers
        int min1 = INT_MAX, min2 = INT_MAX;

        // Traverse the array once
        for (int num : nums)
        {

            // Update the three largest numbers
            if (num >= max1)
            {
                max3 = max2;
                max2 = max1;
                max1 = num;
            }
            else if (num >= max2)
            {
                max3 = max2;
                max2 = num;
            }
            else if (num >= max3)
            {
                max3 = num;
            }

            // Update the two smallest numbers
            if (num <= min1)
            {
                min2 = min1;
                min1 = num;
            }
            else if (num <= min2)
            {
                min2 = num;
            }
        }

        // Product of the three largest numbers
        int product1 = max1 * max2 * max3;

        // Product of two smallest and the largest number
        int product2 = min1 * min2 * max1;

        // Return the larger product
        return max(product1, product2);
    }
};