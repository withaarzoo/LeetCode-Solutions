class Solution:
    def bestClosingTime(self, customers: str) -> int:
        totalY = customers.count('Y')

        openPenalty = 0
        closedPenalty = totalY
        minPenalty = closedPenalty
        answer = 0

        for i, c in enumerate(customers):
            if c == 'N':
                openPenalty += 1
            else:
                closedPenalty -= 1

            currentPenalty = openPenalty + closedPenalty
            if currentPenalty < minPenalty:
                minPenalty = currentPenalty
                answer = i + 1

        return answer
