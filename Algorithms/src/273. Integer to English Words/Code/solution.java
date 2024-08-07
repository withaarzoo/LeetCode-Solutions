class Solution {
    // Arrays to hold the words for numbers below 20, tens, and thousands
    String[] below_20 = { "", "One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Eleven",
            "Twelve", "Thirteen", "Fourteen", "Fifteen", "Sixteen", "Seventeen", "Eighteen", "Nineteen" };
    String[] tens = { "", "", "Twenty", "Thirty", "Forty", "Fifty", "Sixty", "Seventy", "Eighty", "Ninety" };
    String[] thousands = { "", "Thousand", "Million", "Billion" };

    public String numberToWords(int num) {
        // If the number is zero, directly return "Zero"
        if (num == 0)
            return "Zero";

        // Initialize an empty result string
        String result = "";
        // Variable to keep track of the thousand place (units, thousands, millions,
        // etc.)
        int i = 0;

        // Loop until the entire number has been processed
        while (num > 0) {
            // Process each 3-digit segment of the number
            if (num % 1000 != 0) {
                // Convert the 3-digit segment to words and add the corresponding thousand place
                // word
                result = helper(num % 1000) + thousands[i] + " " + result;
            }
            // Move to the next 3-digit segment
            num /= 1000;
            // Increment the thousand place index
            i++;
        }

        // Return the final result, trimmed of any leading or trailing spaces
        return result.trim();
    }

    // Helper function to convert a number less than 1000 to words
    private String helper(int num) {
        // If the number is zero, return an empty string
        if (num == 0)
            return "";

        // If the number is less than 20, return the corresponding word from below_20
        // array
        else if (num < 20)
            return below_20[num] + " ";

        // If the number is less than 100, process the tens and units place
        else if (num < 100)
            return tens[num / 10] + " " + helper(num % 10);

        // If the number is 100 or greater, process the hundreds place and the remaining
        // part
        else
            return below_20[num / 100] + " Hundred " + helper(num % 100);
    }
}
