class Solution
{
public:
    // Function to find GCD using the Euclidean Algorithm
    int gcd(int a, int b)
    {
        // Keep reducing until the remainder becomes 0
        while (b != 0)
        {
            int temp = b; // Store current value of b
            b = a % b;    // Update b with the remainder
            a = temp;     // Move previous b into a
        }

        // a now stores the GCD
        return a;
    }

    int findGCD(vector<int> &nums)
    {
        // Find the smallest element
        int mn = *min_element(nums.begin(), nums.end());

        // Find the largest element
        int mx = *max_element(nums.begin(), nums.end());

        // Return the GCD of the smallest and largest numbers
        return gcd(mn, mx);
    }
};