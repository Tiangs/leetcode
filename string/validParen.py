class Solution:
    def isValid(self, s) :

        match = {
            '(':')',
            '[':']',
            '{':'}'
        }

        stack = []
        if s[0] in ['}',')',']']:
            return False
        for i in range(len(s)):

            if s[i] in match.keys():
                stack.append(s[i])
            else:
                if len(stack) ==0:
                    return False
                if match[stack.pop()] != s[i]:
                    return False
                

        return len(stack)==0 