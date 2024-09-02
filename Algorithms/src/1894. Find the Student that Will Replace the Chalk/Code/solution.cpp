class Solution
{
public:
    int chalkReplacer(vector<int> &chalk, int k)
    {
        // Step 1: Calculate the total amount of chalk required for one complete round.
        // This is done by summing up all the chalk used by each student.
        long long total_chalk = 0; // Initialize total_chalk as a long long to avoid overflow issues.
        for (int c : chalk)
        {
            total_chalk += c; // Accumulate the total chalk used by all students.
        }

        // Step 2: Reduce k by the total amount of chalk needed for one full round.
        // Using modulo operation ensures k is reduced to a value less than total_chalk,
        // as the number of rounds doesn't affect the outcome; only the remainder does.
        k %= total_chalk;

        // Step 3: Determine which student will be the one to replace the chalk.
        // Iterate through each student's chalk usage.
        for (int i = 0; i < chalk.size(); ++i)
        {
            if (k < chalk[i])
            {             // Check if the remaining k chalks are less than what the current student needs.
                return i; // Return the index of the student who will replace the chalk.
            }
            k -= chalk[i]; // Subtract the chalk used by the current student from k.
        }

        // Return -1 as a safety measure, though logically this point should never be reached
        // because the loop will always return an index before it completes.
        return -1;
    }
};
