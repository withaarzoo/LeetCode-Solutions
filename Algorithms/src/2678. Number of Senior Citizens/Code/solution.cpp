#include <vector>
#include <string>
using namespace std;

class Solution
{
public:
    int countSeniors(vector<string> &details)
    {
        int count = 0; // Initialize a counter to keep track of seniors

        // Iterate through each detail in the details vector
        for (const string &detail : details)
        {
            // Extract the substring representing the age, assuming the age is always at the same position
            string age_str = detail.substr(11, 2);
            // Convert the extracted substring to an integer to get the age
            int age = stoi(age_str);

            // Check if the extracted age is greater than 60
            if (age > 60)
            {
                count++; // Increment the counter if the person is a senior
            }
        }

        return count; // Return the total count of seniors
    }
};
