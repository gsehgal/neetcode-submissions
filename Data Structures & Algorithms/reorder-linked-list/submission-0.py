# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next

class Solution:
    def reorderList(self, head: Optional[ListNode]) -> None:

        def traverse(left,right ):
            if right is None:
                return left
            left=traverse(left,right.next)
            if left is None or right is None:
             return left
            temp = left.next
            left.next=right
            if temp == right:
                temp=None
            right.next=temp
            return temp
        length = 0
        left=head
        while left is not None:
            left=left.next
            length+=1
        length=length/2
        right = head
        left = head
        while length>0:
            right=right.next
            length-=1
        left=traverse(left,right)
        if left is not None:
            left.next=None

    