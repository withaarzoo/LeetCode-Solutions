#include <queue>
#include <string>
using namespace std;

class Solution
{
public:
    string longestDiverseString(int a, int b, int c)
    {
        // Priority queue to store the characters and their counts.
        priority_queue<pair<int, char>> pq;
        if (a > 0)
            pq.push({a, 'a'});
        if (b > 0)
            pq.push({b, 'b'});
        if (c > 0)
            pq.push({c, 'c'});

        string result = "";

        while (!pq.empty())
        {
            // Get the character with the most count.
            auto [count1, char1] = pq.top();
            pq.pop();

            // Check if the last two characters in result are the same.
            if (result.size() >= 2 && result.back() == char1 && result[result.size() - 2] == char1)
            {
                if (pq.empty())
                    break; // No valid characters left.

                // Get the second most character.
                auto [count2, char2] = pq.top();
                pq.pop();

                // Add the second character to avoid consecutive repetition.
                result += char2;
                count2--;

                if (count2 > 0)
                    pq.push({count2, char2});

                // Push back the most frequent character for later use.
                pq.push({count1, char1});
            }
            else
            {
                // If no repetition issue, add the most frequent character.
                result += char1;
                count1--;

                if (count1 > 0)
                    pq.push({count1, char1});
            }
        }

        return result;
    }
};