package main

//定义结点
type Node struct {
	Next *Node
	data int
}

//定义链表
type LinkTable struct {
	pHead *Node
	pTail *Node
	lenth int
}

//创建链表
func CreatLinkTable() *LinkTable {
	var pLinkTable *LinkTable = new(LinkTable)
	if pLinkTable == nil {
		return nil
	}
	pLinkTable.pHead = nil
	pLinkTable.pTail = nil
	pLinkTable.lenth = 0
	return pLinkTable
}

//删除链表
func DelLinkTable(pLinktable *LinkTable) int {
	if pLinktable == nil {
		return -1
	}
	for pLinktable.pHead != nil {
		var pNode *Node = pLinktable.pHead
		pLinktable.pHead = pNode.Next
		pLinktable.lenth--
	}
	pLinktable.pHead = nil
	pLinktable.pTail = nil
	pLinktable.lenth = 0
	return 0
}

//增加结点,尾插法
func AddNode(pLinktable *LinkTable, pNode *Node) int {
	if pLinktable == nil || pNode == nil {
		return -1
	}
	pNode.Next = nil
	if pLinktable.pHead == nil {
		pLinktable.pHead = pNode
	}
	if pLinktable.pTail == nil {
		pLinktable.pTail = pNode
	} else {
		pLinktable.pTail.Next = pNode
		pLinktable.pTail = pNode
	}
	pLinktable.lenth++
	return 0
}

//删除结点
func DelNode(pLinktable *LinkTable, pNode *Node) int {
	if pLinktable == nil || pNode == nil {
		return -1
	}
	if pLinktable.pHead == pNode {
		pLinktable.pHead = pLinktable.pHead.Next
		pLinktable.lenth--
		if pLinktable.lenth == 0 {
			pLinktable.pTail = nil
		}
		return 0
	}
	var qNode *Node = pLinktable.pHead
	for qNode != nil {
		if qNode.Next == pNode {
			qNode.Next = qNode.Next.Next
			pLinktable.lenth--
			if pLinktable.lenth == 0 {
				pLinktable.pTail = nil
			}
			return 0
		}
		qNode = qNode.Next
	}
	return -1
}

//查找结点
func SearchNode(pLinktable *LinkTable, pNode *Node) *Node {
	if pLinktable == nil || pNode == nil {
		return nil
	}
	var qNode *Node = pLinktable.pHead
	for qNode != nil {
		if qNode == pNode {
			return qNode
		}
		qNode = qNode.Next
	}
	return nil
}

//获得链表头结点
func GetLinkTableHead(pLinktable *LinkTable) *Node {
	if pLinktable == nil {
		return nil
	}
	return pLinktable.pHead
}

//获得给定结点的下一结点
func GetNextLinkTableNode(pLinktable *LinkTable, pNode *Node) *Node {
	if pLinktable == nil || pNode == nil {
		return nil
	}
	var qNode *Node = pLinktable.pHead
	for qNode != nil {
		if qNode == pNode {
			return qNode.Next
		}
		qNode = qNode.Next
	}
	return nil
}
