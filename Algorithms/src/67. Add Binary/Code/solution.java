class Solution {
    public String addBinary(String a, String b) {
        int i = a.length() - 1;  // pointer for a
        int j = b.length() - 1;  // pointer for b
        int carry = 0;           // carry
        
        StringBuilder result = new StringBuilder();
        
        while (i >= 0 || j >= 0 || carry != 0) {
            int sum = carry;
            
            // Add digit from a
            if (i >= 0) {
                sum += a.charAt(i) - '0';
                i--;
            }
            
            // Add digit from b
            if (j >= 0) {
                sum += b.charAt(j) - '0';
                j--;
            }
            
            // Append current bit
            result.append(sum % 2);
            
            // Update carry
            carry = sum / 2;
        }
        
        // Reverse because digits were added from LSB to MSB
        return result.reverse().toString();
    }
}
