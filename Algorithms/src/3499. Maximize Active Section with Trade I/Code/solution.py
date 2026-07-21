class Solution:
    def maxActiveSectionsAfterTrade(self, s: str) -> int:
        # Count active sections in the original string
        ones = s.count("1")

        # Add virtual boundaries
        t = "1" + s + "1"

        runs = []

        # Run-length encoding
        for ch in t:
            if not runs or runs[-1][0] != ch:
                runs.append([ch, 1])
            else:
                runs[-1][1] += 1

        best = 0

        # Check every internal 1-block
        for i in range(1, len(runs) - 1):
            if (
                runs[i][0] == "1"
                and runs[i - 1][0] == "0"
                and runs[i + 1][0] == "0"
            ):
                best = max(best, runs[i - 1][1] + runs[i + 1][1])

        return ones + best