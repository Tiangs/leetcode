class Solution:
    def missingNumber(self, nums) :
        count = 0
        for i in range(len(nums)):
            if count != nums[i]:
                return count
            count +=1
        return count
        
if __name__ =="__main__":
    s = Solution()
    print(s.missingNumber([0,1,2]))