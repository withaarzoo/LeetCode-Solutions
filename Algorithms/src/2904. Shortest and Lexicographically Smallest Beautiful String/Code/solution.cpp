class Solution
{
public:
    string shortestBeautifulSubstring(string s, int k)
    {
        string answer = ""; // I store the best beautiful substring found so far.
        int left = 0;       // I keep the left boundary of the sliding window.
        int ones = 0;       // I count how many '1' characters are inside the window.

        // I expand the window by moving the right pointer through the string.
        for (int right = 0; right < s.size(); right++)
        {
            // I update the count when the newly added character is '1'.
            if (s[right] == '1')
            {
                ones++;
            }

            // If I have too many ones, I remove characters from the left
            // until the window contains at most k ones again.
            while (ones > k)
            {
                if (s[left] == '1')
                {
                    ones--;
                }
                left++;
            }

            // When I have exactly k ones, I remove leading zeros because
            // they only make the substring longer without adding any ones.
            while (ones == k && s[left] == '0')
            {
                left++;
            }

            // The current window is beautiful when it contains exactly k ones.
            if (ones == k)
            {
                int length = right - left + 1;

                // I extract the current shortest candidate ending at right.
                string candidate = s.substr(left, length);

                // I replace the answer if no answer exists yet, if this candidate
                // is shorter, or if equal-length candidates need lexicographical comparison.
                if (answer.empty() ||
                    candidate.size() < answer.size() ||
                    (candidate.size() == answer.size() && candidate < answer))
                {
                    answer = candidate;
                }
            }
        }

        // I return the best substring, or an empty string if none was found.
        return answer;
    }
};