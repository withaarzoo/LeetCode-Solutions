class Solution {
    public long countCommas(long n) {
        // Start at 1000 because numbers smaller than 1000 have no commas.
        long start = 1000;

        // Numbers in the first comma group contain one comma.
        long commas = 1;

        // This variable stores the total number of commas.
        long answer = 0;

        // Process each comma group until the start goes beyond n.
        while (start <= n) {
            // If start * 1000 would go beyond n, use n as the group end.
            // The division check also avoids unnecessary large multiplication.
            long end = (start > n / 1000)
                    ? n
                    : start * 1000 - 1;

            // Calculate how many numbers belong to this group.
            long count = end - start + 1;

            // Add the contribution of this whole group.
            answer += count * commas;

            // Move to the next group, where the comma count increases by one.
            start *= 1000;
            ++commas;
        }

        // Return the total comma count.
        return answer;
    }
}