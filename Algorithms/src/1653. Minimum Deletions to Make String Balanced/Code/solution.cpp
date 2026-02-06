class Solution
{
public:
    int minimumDeletions(string s)
    {
        int countB = 0;    // number of 'b' seen so far
        int deletions = 0; // minimum deletions needed

        for (char ch : s)
        {
            if (ch == 'b')
            {
                countB++;
            }
            else
            {
                // ch == 'a'
                deletions = min(deletions + 1, countB);
            }
        }
        return deletions;
    }
};
