# Sub category : sliding window
class Solution:
    def lengthOfLongestSubstring(self, s) -> int:
        if not s:
            return 0

        n = len(s)
        maxLen = 0
        i = 0
        j = 0
        d = {}

        while j < n:

            while s[j] in d.keys():
                d.pop(s[i])
                i+=1

            d[s[j]] = j
            maxLen  = max(maxLen,j-i+1)
            j +=1

        return maxLen

if __name__ == "__main__":
    s = Solution()
    print(s.lengthOfLongestSubstring("pwwkew"))
    print(s.lengthOfLongestSubstring("abcbad"))
    print(s.lengthOfLongestSubstring("abba"))
    print(s.lengthOfLongestSubstring("abab"))