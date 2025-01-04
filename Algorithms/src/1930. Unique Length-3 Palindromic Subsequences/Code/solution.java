class Solution {
    public int countPalindromicSubsequence(String s) {
        int[] first = new int[26];
        int[] last = new int[26];
        Arrays.fill(first, -1);

        int n = s.length();
        for (int i = 0; i < n; i++) {
            int index = s.charAt(i) - 'a';
            if (first[index] == -1)
                first[index] = i;
            last[index] = i;
        }

        int result = 0;
        for (int i = 0; i < 26; i++) {
            if (first[i] != -1 && last[i] > first[i]) {
                Set<Character> middleChars = new HashSet<>();
                for (int j = first[i] + 1; j < last[i]; j++) {
                    middleChars.add(s.charAt(j));
                }
                result += middleChars.size();
            }
        }

        return result;
    }
}
