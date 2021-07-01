# Definition for singly-linked list.
class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next

class Solution:
    def reverseList(self, head : ListNode) -> ListNode:
        
        prev = None
        current = head

        while current is not None:

            # create a new node to store the next of current node
            next_node = current.next

            # make this node's pointer points the previous node, which stores in prev
            current.next = prev

            # update prev to be the current node
            prev = current

            # move on the current node to the next node
            current = next_node


        return prev

class Solution_recursive:
    def reverseList(self, head : ListNode) -> ListNode:
        if head is None or head.next is None:
            return head
        
        node = reverseList(self,head.next)

        head.next.next = head
        head.next = None
        return node



if __name__ == "__main__":
    print("=== Linked list test ===")

    head = ListNode(1,None)
    node = head
    print("node.next is {}".format(node.next))
    print("head.next is {}".format(head.next))
    
    node1 = ListNode(2,None)
    node2 = ListNode(5,None)
    node1.next = node2

    node.next = node1
    print("node.next is {}".format(node.next))
    print("head.next is {}".format(head.next))
    
    node = node.next
    print("now node.next is {}".format(node.next))
    print("head.next is {}".format(head.next))