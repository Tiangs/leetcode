# Definition for singly-linked list.
class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next
class Solution:
    def addTwoNumbers(self, l1: ListNode, l2: ListNode) -> ListNode:
        curr = ListNode()
        head = curr
        carry = 0

        while l1 or l2 or carry:
            if not l1 :
                l1 = ListNode(0)
            if not l2:
                l2 = ListNode(0)
            
            ln = ListNode()
            sum = l1.val + l2.val + carry
            carry  = int(sum/10)
            val = sum%10
            ln.val = val

            l1 = l1.next
            l2 = l2.next

            curr.next = ln
            curr = curr.next


        return head.next