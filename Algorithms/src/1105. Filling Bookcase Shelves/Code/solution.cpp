#include <vector>
#include <algorithm>
#include <climits>

class Solution
{
public:
    int minHeightShelves(std::vector<std::vector<int>> &books, int shelfWidth)
    {
        int n = books.size();

        // dp[i] represents the minimum height of the bookshelf that can store the first i books.
        std::vector<int> dp(n + 1, INT_MAX);
        dp[0] = 0; // Base case: No books placed means the height is 0

        // Iterate over each book to determine the minimum shelf height required
        for (int i = 1; i <= n; ++i)
        {
            int width = 0;  // Current shelf's width
            int height = 0; // Current shelf's height

            // Try placing books[i-1], books[i-2], ..., books[j-1] on the current shelf
            for (int j = i; j > 0; --j)
            {
                // Add the width of the current book (books[j-1])
                width += books[j - 1][0];

                // If adding this book exceeds the shelf width, break the loop
                if (width > shelfWidth)
                {
                    break;
                }

                // Update the shelf height as the maximum height of the books placed so far
                height = std::max(height, books[j - 1][1]);

                // Update dp[i] to the minimum of its current value and
                // the height obtained by placing books[j-1] to books[i-1] on the same shelf
                dp[i] = std::min(dp[i], dp[j - 1] + height);
            }
        }

        // The result is the minimum height of the bookshelf that can store all books
        return dp[n];
    }
};
