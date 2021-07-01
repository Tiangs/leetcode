class Solution:
    def moveZeroes(self, nums) -> None:
        n = len(nums)
        slow = quick = 0
        while quick < n:
            if nums[quick] != 0:
                nums[slow], nums[quick] = nums[quick], nums[slow]
                slow += 1
            quick += 1

