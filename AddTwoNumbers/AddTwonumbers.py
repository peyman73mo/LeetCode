
class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next

class Solution:
    def addTwoNumbers(self, l1: Optional[ListNode], l2: Optional[ListNode]) -> Optional[ListNode]:
        
        head = ListNode()
        tail = head
        while True:
            val1 = val2 = 0
            if l1 :
                val1 = l1.val
                l1 = l1.next
            
            if l2 :
                val2 = l2.val
                l2 = l2.next
            
            sum = val1 + val2 + tail.val
            tail.val = sum % 10
            if l1 == None and l2 == None and int(sum / 10) == 0:
                break
            tail.next = ListNode(int(sum / 10),None)
            tail = tail.next
            
        return head
        