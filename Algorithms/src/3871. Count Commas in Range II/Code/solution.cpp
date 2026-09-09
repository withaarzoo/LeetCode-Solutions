class Solution
{
public:
    long long countCommas(long long n)
    {
        // Start at 1000 because numbers below 1000 contain no commas.
        long long start = 1000;

        // Every number from 1000 to 999999 contains exactly one comma.
        long long commas = 1;

        // Store the total number of commas found across all numbers.
        long long answer = 0;

        // Process one comma group at a time.
        while (start <= n)
        {
            // The group normally ends just before start * 1000.
            // If that exceeds n, the group ends at n instead.
            long long end = (start > n / 1000)
                                ? n
                                : start * 1000 - 1;

            // Count how many numbers are present in this group.
            long long count = end - start + 1;

            // Every number in this group has the same number of commas,
            // so I can add their total contribution at once.
            answer += count * commas;

            // Move to the next comma group.
            start *= 1000;

            // The next group has one additional comma.
            ++commas;
        }

        // Return the total number of commas used.
        return answer;
    }
};