# nums1 = [4,1,2], nums2 = [1,3,4,2].
class Solution:
    def nextGreaterElement(self, nums1, nums2 )   :
        stack =[]
        m = {}
        for i in range(len(nums2)):
            while stack and nums2[i] > nums2[stack[-1]]:
                index = stack.pop()
                m[nums2[index]] = nums2[i]
            stack.append(i)
        res = [m[i]if i in m.keys() else -1 for i in nums1]
        return res

if __name__ == "__main__":
    s = Solution()
    print(s.nextGreaterElement(nums1=[4,1,2],nums2=[1,3,4,2]))

