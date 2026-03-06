class Solution
{
public:
    bool checkOnesSegment(string s)
    {
        // If "01" exists, it means a new segment of 1s started
        // after a zero, so return false.
        return s.find("01") == string::npos;
    }
};