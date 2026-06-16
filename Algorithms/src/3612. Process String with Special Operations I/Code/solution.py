class Solution:
    def processStr(self, s: str) -> str:
        # Stores the current result being built
        result = ""

        for c in s:
            # Lowercase letter -> append to result
            if 'a' <= c <= 'z':
                result += c

            # Remove last character if it exists
            elif c == '*':
                if result:
                    result = result[:-1]

            # Duplicate current result
            elif c == '#':
                result += result

            # Reverse current result
            elif c == '%':
                result = result[::-1]

        return result