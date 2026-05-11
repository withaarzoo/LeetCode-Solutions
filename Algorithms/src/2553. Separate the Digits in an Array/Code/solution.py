class Solution:
    def separateDigits(self, nums: List[int]) -> List[int]:
        
        # Final array to store separated digits
        result = []

        # Traverse every number
        for num in nums:

            # Convert number into string
            s = str(num)

            # Traverse every character in string
            for ch in s:

                # Convert character back to integer and store it
                result.append(int(ch))

        # Return final answer
        return result