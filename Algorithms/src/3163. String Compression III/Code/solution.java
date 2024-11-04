class Solution {
    public String compressedString(String word) {
        StringBuilder comp = new StringBuilder();
        int count = 1;

        for (int i = 1; i <= word.length(); i++) {
            if (i == word.length() || word.charAt(i) != word.charAt(i - 1) || count == 9) {
                comp.append(count).append(word.charAt(i - 1));
                count = 1;
            } else {
                count++;
            }
        }

        return comp.toString();
    }
}
