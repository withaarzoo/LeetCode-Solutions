class Solution {
public:
    string decodeCiphertext(string encodedText, int rows) {
        // Edge case: empty string
        if (encodedText.empty()) return "";
        
        int n = encodedText.size();
        int cols = n / rows;
        
        string result;
        
        // Start from every column in the first row
        for (int startCol = 0; startCol < cols; startCol++) {
            int row = 0;
            int col = startCol;
            
            // Move diagonally down-right
            while (row < rows && col < cols) {
                result += encodedText[row * cols + col];
                row++;
                col++;
            }
        }
        
        // Remove trailing spaces
        while (!result.empty() && result.back() == ' ') {
            result.pop_back();
        }
        
        return result;
    }
};