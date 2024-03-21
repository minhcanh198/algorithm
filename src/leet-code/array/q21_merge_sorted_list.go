package array

import (
	"encoding/json"
	"os"
)

// Merge
//21. Merge Two Sorted Lists
//Easy
//Topics
//Companies
//You are given the heads of two sorted linked lists list1 and list2.
//
//Merge the two lists into one sorted list. The list should be made by splicing together the nodes of the first two lists.
//
//Return the head of the merged linked list.

//
//fmt.Println("mica course dau xanh ==>>>>>>>>>>>>>>>>>>>>")
//
//list1 := &array.ListNode{
//Val: 1,
//Next: &array.ListNode{
//Val: 2,
//Next: &array.ListNode{
//Val:  4,
//Next: nil,
//},
//},
//}
//
//list2 := &array.ListNode{
//Val: 1,
//Next: &array.ListNode{
//Val: 3,
//Next: &array.ListNode{
//Val:  4,
//Next: nil,
//},
//},
//}
//
//array.MergeTwoLists(list1, list2)
type ListNode struct {
	Val  int
	Next *ListNode
}

func JsonPrint(myObject interface{}) {

	// Encode the object
	data, err := json.Marshal(myObject)
	if err != nil {
		panic(err)
	}

	// Print to standard output
	os.Stdout.Write(data)
	os.Stdout.WriteString("\n")

}

func MergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	//JsonPrint(list1)
	//JsonPrint(list2)

	mergeList := &ListNode{}
	current := mergeList
	for list1.Next != nil && list2.Next != nil {
		if list1.Val < list2.Val {
			current.Next = list1
			list1 = list1.Next
		} else {
			current.Next = list2
			list2 = list2.Next
		}

		current = current.Next
		//fmt.Println("current")
		//JsonPrint( current)
	}

	if list1.Next == nil {
		current.Next = list2
	} else {
		current.Next = list1
	}

	//JsonPrint(mergeList.Next)
	return mergeList
}
