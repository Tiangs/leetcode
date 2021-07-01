class Solution:
    def threeSum(self, nums ): 

        if not nums:
            return []

        n = len(nums)
        nums.sort()
        res = []

        for i in range(n-2):
            
            # cutting branch
            if nums[i] + nums[i+1] + nums[i+2] > 0:
                break
            
            # move i to the next element to find a bigger one 
            if nums[i] + nums[n-1] + nums[n-2] < 0:
                continue

            # Avoid repeating result
            if i > 0 and nums[i] == nums[i-1]: 
                continue

            # starting from two endpoint
            l, r = i+1, n-1 

            while l < r:
                tmp = nums[i] + nums[l] + nums[r]

                if tmp == 0:
                    # save this result 
                    res.append([nums[i],nums[l],nums[r]])

                    # Continue moving l to find another different value, avoid repeating answer
                    while nums[l] == nums[l+1] and l+1 < r:
                        l += 1
                    l+=1

                    # same for r
                    while nums[r] == nums[r-1] and l<r-1:
                        r -= 1
                    r-=1

                elif tmp < 0:
                    l += 1
                else:
                    r -= 1
                
        return res

if __name__ == "__main__":
    s  = Solution()
    print(s.threeSum([-1,0,1,2,-1,-4]))