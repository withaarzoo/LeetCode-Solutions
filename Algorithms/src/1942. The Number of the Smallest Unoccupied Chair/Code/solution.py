import heapq

class Solution:
    def smallestChair(self, times: List[List[int]], targetFriend: int) -> int:
        n = len(times)
        
        # Create a list of arrivals with friend index
        arrivals = [(times[i][0], i) for i in range(n)]
        
        # Sort friends by arrival time
        arrivals.sort()
        
        # Min-Heap to track available chairs
        availableChairs = list(range(n))
        heapq.heapify(availableChairs)

        # Priority queue to track when chairs are freed
        leavingQueue = []
        
        # Iterate through each friend based on arrival
        for arrivalTime, friendIndex in arrivals:
            # Free chairs that are vacated before the current arrival time
            while leavingQueue and leavingQueue[0][0] <= arrivalTime:
                heapq.heappush(availableChairs, heapq.heappop(leavingQueue)[1])
            
            # Assign the smallest available chair
            chair = heapq.heappop(availableChairs)
            
            # If this is the target friend, return their chair number
            if friendIndex == targetFriend:
                return chair
            
            # Mark the chair as being used until the friend's leave time
            heapq.heappush(leavingQueue, (times[friendIndex][1], chair))
        
        return -1  # Should never reach here