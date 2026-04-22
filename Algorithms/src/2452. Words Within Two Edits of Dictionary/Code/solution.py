class Solution:
    def twoEditWords(self, queries: List[str], dictionary: List[str]) -> List[str]:
        result = []

        # Check every query word
        for query in queries:

            # Compare with every dictionary word
            for word in dictionary:
                diff = 0

                # Count character differences
                for i in range(len(query)):
                    if query[i] != word[i]:
                        diff += 1

                    # Stop early if more than 2 edits are needed
                    if diff > 2:
                        break

                # If query matches within 2 edits
                if diff <= 2:
                    result.append(query)
                    break

        return result