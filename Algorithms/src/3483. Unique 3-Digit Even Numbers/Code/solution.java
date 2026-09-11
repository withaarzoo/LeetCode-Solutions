class Solution {
    public int totalNumbers(int[] digits) {
        // I store how many times each digit appears.
        int[] freq = new int[10];

        // I build the frequency table to correctly handle duplicate digits.
        for (int digit : digits) {
            freq[digit]++;
        }

        // I store the number of distinct valid 3-digit even numbers.
        int answer = 0;

        // I choose the hundreds digit from 1 to 9 because zero cannot be the first
        // digit.
        for (int first = 1; first <= 9; first++) {
            // I choose the tens digit from 0 to 9.
            for (int second = 0; second <= 9; second++) {
                // I choose only even digits for the last position.
                for (int third = 0; third <= 8; third += 2) {
                    // I skip the number if any required digit does not exist.
                    if (freq[first] == 0 || freq[second] == 0 || freq[third] == 0) {
                        continue;
                    }

                    // Three equal digits need three copies of that digit.
                    if (first == second && second == third && freq[first] < 3) {
                        continue;
                    }

                    // The first and second positions need two copies when they are equal.
                    if (first == second && freq[first] < 2) {
                        continue;
                    }

                    // The first and third positions need two copies when they are equal.
                    if (first == third && freq[first] < 2) {
                        continue;
                    }

                    // The second and third positions need two copies when they are equal.
                    if (second == third && freq[second] < 2) {
                        continue;
                    }

                    // This combination forms one distinct valid number.
                    answer++;
                }
            }
        }

        // I return the final count.
        return answer;
    }
}