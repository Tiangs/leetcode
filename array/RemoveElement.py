class Solution:
    def removeElement(self, nums, val: int) -> int:
        n = len(nums)
        i = 0
        for j in range(n):
            if nums[j] != val:
                nums[i] = nums[j]
                i+=1

        return i,nums[:i]


if __name__ == "__main__":
    s = Solution()
    print(s.removeElement([2,3,3,2],2))
