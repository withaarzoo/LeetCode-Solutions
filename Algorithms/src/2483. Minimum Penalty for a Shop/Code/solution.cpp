class Solution
{
public:
    int bestClosingTime(string customers)
    {
        int totalY = 0;
        for (char c : customers)
        {
            if (c == 'Y')
                totalY++;
        }

        int openPenalty = 0;
        int closedPenalty = totalY;
        int minPenalty = closedPenalty;
        int answer = 0;

        for (int i = 0; i < customers.size(); i++)
        {
            if (customers[i] == 'N')
            {
                openPenalty++;
            }
            else
            {
                closedPenalty--;
            }

            int currentPenalty = openPenalty + closedPenalty;
            if (currentPenalty < minPenalty)
            {
                minPenalty = currentPenalty;
                answer = i + 1;
            }
        }

        return answer;
    }
};
