class Solution
{
public:
    bool rotateString(string s, string goal)
    {
        // If the lengths differ, s can't be rotated to match goal
        if (s.length() != goal.length())
            return false;

        // Concatenate s with itself to cover all possible rotations
        string doubled = s + s;

        // Check if goal is a substring of doubled
        return doubled.find(goal) != string::npos;
    }
};
