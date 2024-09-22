class Solution
{
public:
    // Function to count how many numbers exist with the current prefix in the range from 1 to n.
    long countSteps(long curr, long n)
    {
        long steps = 0;    // Variable to accumulate the total number of valid steps
        long first = curr; // The lower bound of the range (starts from the current prefix)
        long last = curr;  // The upper bound of the range (starts from the current prefix)

        // This loop calculates how many numbers exist between `first` and `last` inclusive,
        // for each level of the tree-like structure formed by numbers with a given prefix.
        while (first <= n)
        {
            // At each level, the range of numbers with the current prefix is from `first` to `min(n + 1, last + 1)`.
            steps += min(n + 1, last + 1) - first;

            // Move to the next level by multiplying the bounds by 10.
            // This shifts the prefix further (e.g., 1 -> 10, 100, etc.).
            first *= 10;
            last = last * 10 + 9; // Expand the upper bound by appending 9 at the end (e.g., 1 -> 19, 100 -> 199).
        }

        // Return the total number of valid steps/numbers with the current prefix.
        return steps;
    }

    // Function to find the k-th smallest number in the range from 1 to n.
    int findKthNumber(int n, int k)
    {
        long curr = 1; // Start from the first prefix, which is 1.
        k--;           // Decrease k by 1 because we're starting at 1, so effectively looking for the (k-1)-th next number.

        // Continue searching until we locate the k-th number.
        while (k > 0)
        {
            // Calculate how many numbers exist in the lexicographical tree with the current prefix.
            long steps = countSteps(curr, n);

            // If the number of steps with the current prefix is less than or equal to k,
            // it means that the k-th number is not within this prefix's subtree.
            if (steps <= k)
            {
                // Move to the next sibling by incrementing `curr`.
                curr++;
                // Decrease k by the number of steps we've skipped over in the current subtree.
                k -= steps;
            }
            else
            {
                // If the k-th number lies within the current prefix's subtree, move down to the next level.
                curr *= 10; // Move to the first child by appending a '0' (e.g., 1 -> 10).
                k--;        // We've already accounted for the current number, so decrement k.
            }
        }

        // Once k reaches 0, `curr` will hold the k-th smallest number.
        return curr;
    }
};
