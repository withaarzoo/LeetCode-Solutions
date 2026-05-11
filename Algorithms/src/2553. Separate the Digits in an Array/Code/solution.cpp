class Solution
{
public:
    vector<int> separateDigits(vector<int> &nums)
    {

        // Final array that will store all separated digits
        vector<int> result;

        // Traverse every number in the input array
        for (int num : nums)
        {

            // Convert number into string so digits become easy to access
            string s = to_string(num);

            // Traverse every character in the string
            for (char ch : s)
            {

                // Convert character digit into integer and store it
                result.push_back(ch - '0');
            }
        }

        // Return final separated digits array
        return result;
    }
};