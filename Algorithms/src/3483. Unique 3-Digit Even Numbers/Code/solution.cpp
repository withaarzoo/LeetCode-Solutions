class Solution
{
public:
    int totalNumbers(vector<int> &digits)
    {
        // I store how many times each digit appears.
        int freq[10] = {};

        // I build the frequency table so I can handle duplicate digits correctly.
        for (int digit : digits)
        {
            freq[digit]++;
        }

        // I store the total number of distinct valid 3-digit even numbers.
        int answer = 0;

        // I choose the hundreds digit from 1 to 9 because leading zero is not allowed.
        for (int first = 1; first <= 9; first++)
        {
            // I choose the tens digit from 0 to 9.
            for (int second = 0; second <= 9; second++)
            {
                // I choose only even digits for the ones position.
                for (int third = 0; third <= 8; third += 2)
                {
                    // I check whether all three required digits are available.
                    if (freq[first] == 0 || freq[second] == 0 || freq[third] == 0)
                    {
                        // If any required digit is missing, this number cannot be formed.
                        continue;
                    }

                    // If the same digit is used in multiple positions, I need enough copies.
                    if (first == second && second == third && freq[first] < 3)
                    {
                        // Three equal digits require at least three copies.
                        continue;
                    }

                    // If the first two digits are equal, I need at least two copies.
                    if (first == second && freq[first] < 2)
                    {
                        // There are not enough copies of the first digit.
                        continue;
                    }

                    // If the first and third digits are equal, I need at least two copies.
                    if (first == third && freq[first] < 2)
                    {
                        // There are not enough copies of the first digit.
                        continue;
                    }

                    // If the second and third digits are equal, I need at least two copies.
                    if (second == third && freq[second] < 2)
                    {
                        // There are not enough copies of the second digit.
                        continue;
                    }

                    // All three digits are available, so this distinct number is valid.
                    answer++;
                }
            }
        }

        // I return the number of distinct valid 3-digit even numbers.
        return answer;
    }
};