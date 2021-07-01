class Solution:

    def findDisappearedNumbers(self, nums):
        l = len(nums)

        for n in nums:
            x = (n -1) % l
            nums[x] += l

        return [i+1 for i,val in enumerate(nums) if val <= l]

if __name__ == "__main__":
    s = Solution()
    print(s.findDisappearedNumbers([4,3,2,7,8,2,3,1]))
