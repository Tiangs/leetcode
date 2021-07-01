class Solution:
    def lengthOfLongestSubstring(self, s) -> int:

        if not s:
            return 0
        d = {}
        maxLen = 0
        left,right =0,0
        
        while right < len(s):
            
            # moving left pointer to the next character different to right pointer
            while s[right] in d.keys():

                # remove left from the dict  abcba
                d.pop(s[left])
                left +=1


            # update dict 
            d[s[right]] = right

            # update maxLen
            maxLen = max(maxLen,right-left+1)

            right+=1
            
        return maxLen


if __name__ == "__main__":
    s = Solution()
    print(s.lengthOfLongestSubstring("abba"))
    print(s.lengthOfLongestSubstring("bbbbb"))
    print(s.lengthOfLongestSubstring("abcabcbb"))
    print(s.lengthOfLongestSubstring("abcba"))