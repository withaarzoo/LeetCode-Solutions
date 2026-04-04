class Solution {
    public String decodeCiphertext(String encodedText, int rows) {
        // Edge case: empty string
        if (encodedText.length() == 0)
            return "";

        int n = encodedText.length();
        int cols = n / rows;

        StringBuilder result = new StringBuilder();

        // Start from every column in the first row
        for (int startCol = 0; startCol < cols; startCol++) {
            int row = 0;
            int col = startCol;

            // Move diagonally down-right
            while (row < rows && col < cols) {
                result.append(encodedText.charAt(row * cols + col));
                row++;
                col++;
            }
        }

        // Remove trailing spaces
        while (result.length() > 0 && result.charAt(result.length() - 1) == ' ') {
            result.deleteCharAt(result.length() - 1);
        }

        return result.toString();
    }
}