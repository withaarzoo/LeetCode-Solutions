class Solution
{
public:
    int findTheWinner(int n, int k)
    {
        // The problem is equivalent to the Josephus problem.
        // We need to convert the 0-indexed result from the Josephus function to 1-indexed.
        return josephus(n, k) + 1;
    }

    int josephus(int n, int k)
    {
        // Base case: when there is only one person, they are the winner.
        if (n == 1)
        {
            return 0;
        }
        // Recursive case:
        // josephus(n - 1, k) finds the position of the winner in the reduced problem (n-1 persons)
        // Adding k accounts for the step size in the problem.
        // Taking modulo n wraps around the position if it exceeds the number of people.
        return (josephus(n - 1, k) + k) % n;
    }
};
