import java.util.Stack;

class Solution {
    public int minLength(String s) {
        Stack<Character> stack = new Stack<>();

        // Traverse through each character in the string
        for (char ch : s.toCharArray()) {
            // If the top of the stack forms "AB" or "CD" with the current character, pop
            // the stack
            if (!stack.isEmpty() && ((stack.peek() == 'A' && ch == 'B') || (stack.peek() == 'C' && ch == 'D'))) {
                stack.pop(); // Remove the substring
            } else {
                stack.push(ch); // Push character onto the stack if no pair is found
            }
        }

        return stack.size(); // The size of the stack is the minimum length of the string
    }
}