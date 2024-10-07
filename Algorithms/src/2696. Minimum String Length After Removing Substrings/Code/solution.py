class Solution:
    def minLength(self, s: str) -> int:
        stack = []
        
        # Traverse through each character in the string
        for ch in s:
            # If top of the stack forms "AB" or "CD" with the current character, pop the stack
            if stack and ((stack[-1] == 'A' and ch == 'B') or (stack[-1] == 'C' and ch == 'D')):
                stack.pop()  # Remove the substring
            else:
                stack.append(ch)  # Push current character onto the stack if no pair
        
        return len(stack)  # The size of the stack is the minimum length