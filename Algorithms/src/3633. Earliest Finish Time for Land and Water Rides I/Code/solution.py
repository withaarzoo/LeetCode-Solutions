class Solution:
    def earliestFinishTime(self, landStartTime: List[int], landDuration: List[int],
                           waterStartTime: List[int], waterDuration: List[int]) -> int:

        # Store minimum finishing time
        ans = float('inf')

        n = len(landStartTime)
        m = len(waterStartTime)

        # Try every land ride with every water ride
        for i in range(n):
            for j in range(m):

                # Land -> Water
                land_finish = landStartTime[i] + landDuration[i]

                # Water starts when both conditions are satisfied:
                # land ride finished and water ride opened
                water_start = max(land_finish, waterStartTime[j])

                finish1 = water_start + waterDuration[j]

                # Water -> Land
                water_finish = waterStartTime[j] + waterDuration[j]

                # Land starts when both conditions are satisfied:
                # water ride finished and land ride opened
                land_start = max(water_finish, landStartTime[i])

                finish2 = land_start + landDuration[i]

                # Keep the best answer
                ans = min(ans, finish1, finish2)

        return ans