class Solution:
    def validateCoupons(self, code, businessLine, isActive):

        priority = {
            "electronics": 0,
            "grocery": 1,
            "pharmacy": 2,
            "restaurant": 3
        }

        valid = []

        for i in range(len(code)):

            if not isActive[i]:
                continue

            if businessLine[i] not in priority:
                continue

            if not code[i]:
                continue

            if not all(c.isalnum() or c == '_' for c in code[i]):
                continue

            valid.append((priority[businessLine[i]], code[i]))

        valid.sort()
        return [c for _, c in valid]
