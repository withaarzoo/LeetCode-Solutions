class Solution
{
public:
    // Arrays to store the English words for numbers below 20, tens, and thousands
    vector<string> below_20 = {"", "One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten",
                               "Eleven", "Twelve", "Thirteen", "Fourteen", "Fifteen", "Sixteen", "Seventeen",
                               "Eighteen", "Nineteen"};
    vector<string> tens = {"", "", "Twenty", "Thirty", "Forty", "Fifty", "Sixty", "Seventy", "Eighty", "Ninety"};
    vector<string> thousands = {"", "Thousand", "Million", "Billion"};

    // Main function to convert number to words
    string numberToWords(int num)
    {
        // Special case for zero
        if (num == 0)
            return "Zero";

        // Variable to store the final result
        string result;
        // Index to track the position (thousands, millions, billions)
        int i = 0;

        // Process each group of three digits
        while (num > 0)
        {
            // If the current group of three digits is not zero
            if (num % 1000 != 0)
            {
                // Convert the current group to words and prepend to the result
                result = helper(num % 1000) + thousands[i] + " " + result;
            }
            // Move to the next group of three digits
            num /= 1000;
            // Increment the position index
            i++;
        }

        // Remove any trailing spaces from the result
        while (result.back() == ' ')
            result.pop_back();

        // Return the final result
        return result;
    }

    // Helper function to convert a number less than 1000 to words
    string helper(int num)
    {
        // If the number is zero, return an empty string
        if (num == 0)
            return "";
        // If the number is less than 20, return the corresponding word from below_20 array
        else if (num < 20)
            return below_20[num] + " ";
        // If the number is less than 100, return the word for the tens place and recursively call helper for the ones place
        else if (num < 100)
            return tens[num / 10] + " " + helper(num % 10);
        // If the number is 100 or more, return the word for the hundreds place and recursively call helper for the rest
        else
            return below_20[num / 100] + " Hundred " + helper(num % 100);
    }
};
