/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func rotateRight(head *ListNode, k int) *ListNode {
    // edge cases
    if head == nil || head.Next == nil || k == 0 {
        return head
    }
    var curr *ListNode
    curr = head
    length := 1
    for curr.Next != nil {
        length++
        curr = curr.Next
    }
    curr.Next = head
    // If K is greter then length of the list then reduce K
    if (k >= length) {
    k = k % length
    }
    k = length - k
    for(k > 0) {
        curr = curr.Next
        k--
    }
    head = curr.Next
    curr.Next = nil
    return head
}
