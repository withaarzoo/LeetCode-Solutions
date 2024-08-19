class Solution
{
public:
    int minSteps(int n)
    {
        // Initialize the number of operations required to 0
        int operations = 0;

        // Start from 2 and iterate up to n
        for (int i = 2; i <= n; i++)
        {
            // Check if the current value of 'n' is divisible by 'i'
            while (n % i == 0)
            {
                // If divisible, add 'i' to the operations count
                operations += i;

                // Reduce 'n' by dividing it by 'i'
                n /= i;
            }
        }

        // Return the total number of operations needed
        return operations;
    }
};
