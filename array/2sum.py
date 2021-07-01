 # target int
 # output list
class Solution:
    def twoSum(self, nums, target: int) :

        if not nums:
            return None
        
        # key: the remaining value 
        # value: current index
        m={}

        for i in range(len(nums)):
            
            # query if the current value in m
            if nums[i] in m.keys():
                return [i,m[nums[i]]]

            # save this current num
            m[target - nums[i]] = i
        
        return None

            
if __name__ == "__main__":
    s = Solution()
    print(s.twoSum([1,2,3,4],7))
