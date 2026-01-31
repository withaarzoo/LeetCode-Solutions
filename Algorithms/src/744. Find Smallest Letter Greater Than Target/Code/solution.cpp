class Solution
{
public:
    char nextGreatestLetter(vector<char> &letters, char target)
    {
        int left = 0;
        int right = letters.size() - 1;
        char answer = letters[0]; // default for wrap-around case

        while (left <= right)
        {
            int mid = left + (right - left) / 2;

            if (letters[mid] > target)
            {
                answer = letters[mid]; // possible answer
                right = mid - 1;       // try to find smaller valid one
            }
            else
            {
                left = mid + 1; // move right
            }
        }

        return answer;
    }
};
