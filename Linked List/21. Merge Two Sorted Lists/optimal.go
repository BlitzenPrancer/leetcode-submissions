/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
    if l1 == nil {
        return l2
    }
    if l2 == nil {
        return l1
    }
    if l1.Val > l2.Val {
        var temp *ListNode
        temp = l1
        l1 = l2
        l2 = temp
    }
    var temp *ListNode
    var res *ListNode
    res = l1  // point res to l1
    for l1 != nil && l2 != nil {
        for l1 != nil && l1.Val <= l2.Val {
            temp = l1
            l1 = l1.Next
        }
        temp.Next = l2
        // swap l1 and l2 because l1 should always be smallest and l2 should always be biggest
        var tmp *ListNode
        tmp = l1
        l1 = l2
        l2 = tmp
    }
    return res
}
