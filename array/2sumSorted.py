class Solution:
    def twoSum(self, numbers, target) -> int:
        n = len(numbers)
        i = 0
        j = n-1
        while i < j:
            temp = numbers[i] + numbers[j]
            if temp == target:
                return [i+1,j+1]
            elif temp < target:
                i +=1
            else:
                j-=1
                
        return None

        
if __name__ == "__main__":
    s = Solution()
    print(s.twoSum( [2,7,11,15],9))
    print(s.twoSum( [2,3,4],6))
