class Solution
{
public:
    int totalWaviness(int num1, int num2)
    {
        int answer = 0;

        // Check every number in the given range
        for (int num = num1; num <= num2; num++)
        {
            string s = to_string(num);

            // Numbers with fewer than 3 digits cannot have peaks or valleys
            if (s.size() < 3)
            {
                continue;
            }

            // Check every middle digit
            for (int i = 1; i < (int)s.size() - 1; i++)
            {
                // Peak condition
                if (s[i] > s[i - 1] && s[i] > s[i + 1])
                {
                    answer++;
                }
                // Valley condition
                else if (s[i] < s[i - 1] && s[i] < s[i + 1])
                {
                    answer++;
                }
            }
        }

        return answer;
    }
};