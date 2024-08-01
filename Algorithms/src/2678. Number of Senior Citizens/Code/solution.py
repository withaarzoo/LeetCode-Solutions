from typing import List

class Solution:
    def countSeniors(self, details: List[str]) -> int:
        # Initialize a counter to keep track of the number of seniors
        count = 0

        # Iterate through each detail string in the list
        for detail in details:
            # Extract the substring that represents the age, located at indices 11 and 12
            age_str = detail[11:13]

            # Convert the extracted substring to an integer to get the age
            age = int(age_str)

            # Check if the extracted age is greater than 60
            if age > 60:
                # If the age is greater than 60, increment the senior count
                count += 1

        # Return the total count of seniors
        return count
