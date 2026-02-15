class Solution {
public:
    string addBinary(string a, string b) {
        int i = a.size() - 1;   // pointer for string a
        int j = b.size() - 1;   // pointer for string b
        int carry = 0;          // carry for binary addition
        
        string result = "";
        
        // Loop until both strings and carry are processed
        while (i >= 0 || j >= 0 || carry) {
            int sum = carry;
            
            // Add digit from a if available
            if (i >= 0) {
                sum += a[i] - '0';
                i--;
            }
            
            // Add digit from b if available
            if (j >= 0) {
                sum += b[j] - '0';
                j--;
            }
            
            // Current digit is sum % 2
            result += (sum % 2) + '0';
            
            // Update carry
            carry = sum / 2;
        }
        
        // Reverse result because we built it backwards
        reverse(result.begin(), result.end());
        
        return result;
    }
};
