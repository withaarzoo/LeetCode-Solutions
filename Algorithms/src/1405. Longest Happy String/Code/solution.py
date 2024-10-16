import heapq

class Solution:
    def longestDiverseString(self, a: int, b: int, c: int) -> str:
        # Max heap to always pick the character with the highest count.
        pq = []
        if a > 0:
            heapq.heappush(pq, (-a, 'a'))
        if b > 0:
            heapq.heappush(pq, (-b, 'b'))
        if c > 0:
            heapq.heappush(pq, (-c, 'c'))

        result = []

        while pq:
            count1, char1 = heapq.heappop(pq)

            # If the last two characters are the same as char1.
            if len(result) >= 2 and result[-1] == char1 and result[-2] == char1:
                if not pq:
                    break  # No valid characters left to pick.

                count2, char2 = heapq.heappop(pq)
                result.append(char2)
                count2 += 1  # Decrease count (negated)

                if count2 < 0:
                    heapq.heappush(pq, (count2, char2))

                heapq.heappush(pq, (count1, char1))
            else:
                result.append(char1)
                count1 += 1  # Decrease count (negated)

                if count1 < 0:
                    heapq.heappush(pq, (count1, char1))

        return ''.join(result)