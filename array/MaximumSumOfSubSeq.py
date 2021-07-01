class Solution:
    # Greedy solution
    def maxSubArray(self, nums: List[int]) -> int:
        max_sum = nums[0]
        cur_sum = 0

        for i in range(len(nums)):
            cur_sum = max(nums[i],cur_sum + nums[i])
            max_sum = max(cur_sum,max_sum)
        
        return max_sum