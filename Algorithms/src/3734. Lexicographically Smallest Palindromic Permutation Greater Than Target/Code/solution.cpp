class Solution {
public:
    // Builds the complete palindrome from its left half and optional middle character.
    string buildPalindrome(const string& half, char middle) {
        string result = half; // Start with the chosen left half.

        if (middle != 0) {
            result += middle; // Add the fixed middle character for odd-length strings.
        }

        // Mirror the left half in reverse order to complete the palindrome.
        for (int i = (int)half.size() - 1; i >= 0; --i) {
            result += half[i];
        }

        return result;
    }

    // Finds the lexicographically smallest permutation of the multiset
    // that is greater than or equal to targetHalf.
    string smallestGreaterOrEqual(vector<int> count, const string& targetHalf) {
        int k = targetHalf.size(); // Number of characters in the palindrome's first half.
        int matched = 0; // Number of target characters matched exactly so far.

        // Try to match targetHalf from left to right for as long as possible.
        while (matched < k && count[targetHalf[matched] - 'a'] > 0) {
            --count[targetHalf[matched] - 'a']; // Use this exact character.
            ++matched; // Move to the next position.
        }

        // If every position matched, targetHalf itself is a valid permutation.
        if (matched == k) {
            return targetHalf;
        }

        // Backtrack to find the rightmost position that can be increased.
        for (int pos = matched; pos >= 0; --pos) {
            // When moving left, restore the character that was previously matched.
            if (pos < matched) {
                ++count[targetHalf[pos] - 'a'];
            }

            // Choose the smallest available character strictly greater than targetHalf[pos].
            for (int c = targetHalf[pos] - 'a' + 1; c < 26; ++c) {
                if (count[c] == 0) continue; // This character is not available.

                string result = targetHalf.substr(0, pos); // Keep the prefix unchanged.
                result += char('a' + c); // Increase this position by the smallest possible amount.
                --count[c]; // Consume the chosen larger character.

                // Fill the remaining positions in ascending order for the smallest result.
                for (int ch = 0; ch < 26; ++ch) {
                    result.append(count[ch], char('a' + ch));
                }

                return result;
            }
        }

        return ""; // No permutation can be greater than or equal to targetHalf.
    }

    // Returns the next lexicographical permutation of the first half.
    bool nextPermutation(string& half) {
        int n = half.size();
        int pivot = n - 2; // Start by searching for the rightmost increasing position.

        // Find the rightmost position where half[pivot] < half[pivot + 1].
        while (pivot >= 0 && half[pivot] >= half[pivot + 1]) {
            --pivot;
        }

        // The whole sequence is non-increasing, so no larger permutation exists.
        if (pivot < 0) {
            return false;
        }

        int swapPos = n - 1; // Find the smallest larger character from the suffix.

        // Because the suffix is non-increasing, the first valid character from the right is correct.
        while (half[swapPos] <= half[pivot]) {
            --swapPos;
        }

        swap(half[pivot], half[swapPos]); // Increase the pivot position.

        // Reverse the suffix to make it as small as possible.
        reverse(half.begin() + pivot + 1, half.end());

        return true;
    }

    string lexPalindromicPermutation(string s, string target) {
        vector<int> frequency(26, 0); // Count every character in s.

        for (char ch : s) {
            ++frequency[ch - 'a'];
        }

        char middle = 0; // Stores the unique odd-frequency character, if needed.
        int oddCount = 0; // Counts how many characters have odd frequency.

        for (int c = 0; c < 26; ++c) {
            if (frequency[c] % 2 == 1) {
                ++oddCount;
                middle = char('a' + c);
            }
        }

        // A palindrome can have at most one odd-frequency character.
        if (oddCount > 1) {
            return "";
        }

        vector<int> halfCount(26, 0); // Frequency multiset for the first half.

        for (int c = 0; c < 26; ++c) {
            halfCount[c] = frequency[c] / 2; // Only half of each pair goes to the left side.
        }

        int k = s.size() / 2; // Length of the first half.
        string targetHalf = target.substr(0, k); // Only this prefix controls the first comparison.

        // Find the smallest possible first half that is at least targetHalf.
        string half = smallestGreaterOrEqual(halfCount, targetHalf);

        if (half.empty() && k > 0) {
            return ""; // No first-half permutation can reach targetHalf.
        }

        // Build the palindrome corresponding to this smallest valid first half.
        string candidate = buildPalindrome(half, middle);

        // If it is already strictly greater, it is the required smallest answer.
        if (candidate > target) {
            return candidate;
        }

        // Otherwise, move to the next larger first-half permutation.
        if (!nextPermutation(half)) {
            return "";
        }

        // The next first half produces the smallest possible larger palindrome.
        return buildPalindrome(half, middle);
    }
};