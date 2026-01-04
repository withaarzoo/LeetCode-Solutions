class Solution:
    def sumFourDivisors(self, nums: List[int]) -> int:
        total_sum = 0

        for num in nums:
            cnt = 0
            div_sum = 0

            d = 1
            while d * d <= num:
                if num % d == 0:
                    other = num // d

                    cnt += 1
                    div_sum += d

                    if other != d:
                        cnt += 1
                        div_sum += other

                    if cnt > 4:
                        break
                d += 1

            if cnt == 4:
                total_sum += div_sum

        return total_sum
