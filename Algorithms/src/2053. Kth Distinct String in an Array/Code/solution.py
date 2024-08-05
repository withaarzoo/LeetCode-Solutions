class Solution:
    def kthDistinct(self, arr: List[str], k: int) -> str:
        count = {}  # Dictionary to store the frequency of each string
        distinct = []  # List to store distinct strings

        # Iterate through each string in the array to count occurrences
        for str in arr:
            count[str] = count.get(str, 0) + 1  # Increment the count for each string

        # Iterate through the array again to collect distinct strings
        for str in arr:
            if count[str] == 1:  # Check if the string is distinct (appears only once)
                distinct.append(str)  # Add distinct string to the list

        # Check if the k-th distinct string exists
        if k <= len(distinct):
            return distinct[k-1]  # Return the k-th distinct string (1-based index)
        else:
            return ""  # Return an empty string if k is out of range
