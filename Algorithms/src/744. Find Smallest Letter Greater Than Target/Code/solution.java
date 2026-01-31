class Solution {
    public char nextGreatestLetter(char[] letters, char target) {
        int left = 0;
        int right = letters.length - 1;
        char answer = letters[0]; // wrap-around default

        while (left <= right) {
            int mid = left + (right - left) / 2;

            if (letters[mid] > target) {
                answer = letters[mid];
                right = mid - 1;
            } else {
                left = mid + 1;
            }
        }

        return answer;
    }
}
