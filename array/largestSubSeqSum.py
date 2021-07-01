class Solution:
    def maxSubArray(self, nums: List[int]) -> int:
        max_val = 0
        n = len(nums)

        while i < n:
            j = i + 1
            sum = i
            while j < n:
                pass 