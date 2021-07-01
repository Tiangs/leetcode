 
def monoQ(nums):
    stack =[]
    for i in range(len(nums)):
        while stack and nums[i] < stack[-1]:
            stack.pop()

        stack.append(nums[i])

    return stack 
if __name__ == "__main__":
    nums = [1,5,3,7,2,1,2,9]
    print(monoQ(nums))
