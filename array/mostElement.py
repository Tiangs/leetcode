class Solution:
    def majorityElement(self, nums: List[int]) -> int:
        occurence = {}
        for i in range(len(nums)):
            if nums[i] not in occurence.keys():
                occurence[nums[i]] = 1
            else:
                occurence[nums[i]] += 1
            
        for i in occurence.keys():
            if occurence[i] >= len(nums)/2:
                return i
            

        

