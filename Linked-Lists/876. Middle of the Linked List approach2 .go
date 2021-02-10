/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func middleNode(head *ListNode) *ListNode {
    output := []*ListNode{}
    for head != nil {
        output = append(output, head)
        head = head.Next
    }
    return output[len(output)/2]
}
