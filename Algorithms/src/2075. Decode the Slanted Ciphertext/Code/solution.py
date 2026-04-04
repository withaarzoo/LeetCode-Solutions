class Solution:
    def decodeCiphertext(self, encodedText: str, rows: int) -> str:
        # Edge case: empty string
        if not encodedText:
            return ""
        
        n = len(encodedText)
        cols = n // rows
        
        result = []
        
        # Start from every column in the first row
        for start_col in range(cols):
            row = 0
            col = start_col
            
            # Move diagonally down-right
            while row < rows and col < cols:
                result.append(encodedText[row * cols + col])
                row += 1
                col += 1
        
        # Remove trailing spaces
        return ''.join(result).rstrip()