class Solution:
    def compareVersion(self, version1: str, version2: str) -> int:
        i, j = 0, 0
        n, m = len(version1), len(version2)

        while i < n or j < m:
            num1 = 0
            num2 = 0
            # parse next revision from version1
            while i < n and version1[i] != '.':
                num1 = num1 * 10 + (ord(version1[i]) - ord('0'))
                i += 1
            if i < n and version1[i] == '.':
                i += 1

            # parse next revision from version2
            while j < m and version2[j] != '.':
                num2 = num2 * 10 + (ord(version2[j]) - ord('0'))
                j += 1
            if j < m and version2[j] == '.':
                j += 1

            if num1 < num2:
                return -1
            if num1 > num2:
                return 1

        return 0
