class Solution {
    public int countSeniors(String[] details) {
        int count = 0; // Initialize a counter for the number of seniors

        // Iterate through each detail in the input array
        for (String detail : details) {
            // Extract the age substring from the detail string
            // Age is located at index 11 and 12 (2 characters)
            String ageStr = detail.substring(11, 13);

            // Convert the age substring to an integer
            int age = Integer.parseInt(ageStr);

            // Check if the age is greater than 60
            if (age > 60) {
                count++; // Increment the counter if the person is a senior
            }
        }

        // Return the total count of seniors
        return count;
    }
}
