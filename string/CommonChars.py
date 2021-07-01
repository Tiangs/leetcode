class Solution:
    def commonChars(self, A)  :
        res =[]
        min_freq = [float("inf")] *26

        for _,w in enumerate(A):
            freq = [float(0)] * 26
            for c in w:
                freq[ord(c) - ord('a')] +=1
                
            # select the min of freq and minfreq
            for i in range(26):
                min_freq[i] = min(min_freq[i],freq[i])
            
        # outputing the minfreq to result
        for i,v in enumerate(min_freq):
            res += [chr(97+i)]*int(v) 

        return res


if __name__ == "__main__":
    s = Solution()
    print(s.commonChars(["bella","label","roller"]))