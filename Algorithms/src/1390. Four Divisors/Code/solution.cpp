class Solution
{
public:
    int sumFourDivisors(vector<int> &nums)
    {
        int totalSum = 0;

        for (int num : nums)
        {
            int cnt = 0;
            int sum = 0;

            for (int d = 1; d * d <= num; d++)
            {
                if (num % d == 0)
                {
                    int other = num / d;

                    cnt++;
                    sum += d;

                    if (other != d)
                    {
                        cnt++;
                        sum += other;
                    }

                    if (cnt > 4)
                        break;
                }
            }

            if (cnt == 4)
            {
                totalSum += sum;
            }
        }

        return totalSum;
    }
};
