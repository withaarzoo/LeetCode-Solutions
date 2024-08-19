class Solution:
    def minSteps(self, n: int) -> int:
        # Initialize the variable to count the minimum number of operations needed
        operations = 0
        
        # Start with the smallest prime factor, which is 2
        i = 2
        
        # Continue looping until i exceeds n
        while i <= n:
            # If n is divisible by i, it means i is a factor
            # We keep dividing n by i as long as it's divisible to factorize n completely
            while n % i == 0:
                # Add the factor i to the operation count
                # Each division operation corresponds to a sequence of 'Copy All' and 'Paste' operations
                operations += i
                
                # Update n by dividing it by i to remove the factor i
                n //= i
            
            # Move to the next possible factor
            i += 1
        
        # Return the total number of operations required to reach the number n starting from 'A'
        return operations
