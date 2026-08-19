class Solution
{
public:
    int maxNumberOfFamilies(int n, vector<vector<int>> &reservedSeats)
    {
        // Store a bitmask for every row that has at least one reserved seat.
        unordered_map<int, int> rows;

        // Process every reserved seat and mark that seat in its row's bitmask.
        for (const auto &seat : reservedSeats)
        {
            int row = seat[0];
            int col = seat[1];

            // Seats 2 through 9 are represented by bits 1 through 8.
            // Seats 1 and 10 do not belong to any possible group, so I ignore them.
            if (col >= 2 && col <= 9)
            {
                rows[row] |= (1 << col);
            }
        }

        // Start with the rows that have no reserved seats.
        // Every completely empty row can always fit two groups.
        int answer = 2 * (n - static_cast<int>(rows.size()));

        // These masks represent the three possible four-seat blocks.
        // left  = seats 2,3,4,5
        // middle = seats 4,5,6,7
        // right = seats 6,7,8,9
        const int left = (1 << 2) | (1 << 3) | (1 << 4) | (1 << 5);
        const int middle = (1 << 4) | (1 << 5) | (1 << 6) | (1 << 7);
        const int right = (1 << 6) | (1 << 7) | (1 << 8) | (1 << 9);

        // Check only rows that contain reserved seats.
        for (const auto &[row, mask] : rows)
        {
            // A block is available when none of its seats are reserved.
            bool canLeft = (mask & left) == 0;
            bool canMiddle = (mask & middle) == 0;
            bool canRight = (mask & right) == 0;

            // Left and right blocks do not overlap, so both can be used together.
            if (canLeft && canRight)
            {
                answer += 2;
            }
            // Otherwise, if any one block is available, I can place one group.
            else if (canLeft || canMiddle || canRight)
            {
                answer += 1;
            }
            // If no block is available, this row cannot fit a group.
        }

        // Return the maximum number of four-person groups.
        return answer;
    }
};