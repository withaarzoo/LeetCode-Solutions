class Solution:
    def sortableIntegers(self, nums: list[int]) -> int:
        n = len(nums)
        sorted_nums = sorted(nums)
        divs = []
        i = 1
        while i * i <= n:
            if n % i == 0:
                divs.append(i)
                if i != n // i:
                    divs.append(n // i)
            i += 1
        ans = 0
        for k in divs:
            ok = True
            for s in range(0, n, k):
                desc = 0
                pos = -1
                for j in range(k):
                    a = nums[s + j]
                    b = nums[s + (j + 1) % k]
                    if a > b:
                        desc += 1
                        pos = j
                        if desc > 1:
                            break
                if desc > 1:
                    ok = False
                    break
                r = 0 if desc == 0 else (pos + 1) % k
                for j in range(k):
                    if nums[s + (r + j) % k] != sorted_nums[s + j]:
                        ok = False
                        break
                if not ok:
                    break
            if ok:
                ans += k
        return ans