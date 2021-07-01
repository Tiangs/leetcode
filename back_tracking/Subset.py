# back tracking template
class Solution:
    def combine(self, n: int, k: int) :
        if n < k :
            return []

        res = []
        path = []
        def dfs(n,k,start,path,res):
            print("RECURSION start  = {}".format( start))
            print("path = {}".format(path))

            if len(path) == k:
                res.append(path)
                print("res is {}".format(res))
                return
            
            for i in range(start,n+1):
                print("forloop i = {}".format(i))
                path.append(i)
                dfs(n,k,i+1,path,res)
                path.pop()
        
        dfs(n,k,1,path,res)
        return res

# def test(res,n):
#     if n ==0:
#         return  
#     res.append([n])
#     test(res,n-1)

if __name__ == "__main__":
    s = Solution()
    print(s.combine(4,2))

    # res = []
    # test(res,5)
    # print(res)

