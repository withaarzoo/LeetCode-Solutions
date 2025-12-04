class Solution
{
public:
    int countCollisions(string directions)
    {
        int n = directions.size();
        int i = 0, j = n - 1;

        // Skip all leading 'L' cars (they move left forever, no collision)
        while (i < n && directions[i] == 'L')
        {
            i++;
        }

        // Skip all trailing 'R' cars (they move right forever, no collision)
        while (j >= 0 && directions[j] == 'R')
        {
            j--;
        }

        int collisions = 0;
        // In the middle part, every 'L' or 'R' will collide exactly once
        for (int k = i; k <= j; k++)
        {
            if (directions[k] != 'S')
            {
                collisions++;
            }
        }

        return collisions;
    }
};
