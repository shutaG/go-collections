package data_structure

import "sync"

type AvlTree struct {
	root    *avlNode
	size    int
	RWMutex sync.RWMutex
}

type BinaryTree interface {
	Insert(e int)     //插入元素
	Delete(e int)     //删除元素
	PreOrder() []int  //前序遍历
	InOrder() []int   //中序遍历
	PostOrder() []int //后序遍历
	Size() (num int)  //返回树中的元素个数
	Clear()           //清空树中所有元素
	Empty() (b bool)  //判断树是否为空
}

type avlNode struct {
	key    int
	height int
	left   *avlNode
	right  *avlNode
}

func (a AvlTree) Insert(e int) {
	//TODO implement me
	panic("implement me")
}

func (a AvlTree) Delete(e int) {
	//TODO implement me
	panic("implement me")
}

func (a AvlTree) PreOrder() []int {
	//TODO implement me
	panic("implement me")
}

func (a AvlTree) InOrder() []int {
	//TODO implement me
	panic("implement me")
}

func (a AvlTree) PostOrder() []int {
	//TODO implement me
	panic("implement me")
}

func (a AvlTree) Size() (num int) {
	//TODO implement me
	panic("implement me")
}

func (a AvlTree) Clear() {
	//TODO implement me
	panic("implement me")
}

func (a AvlTree) Empty() (b bool) {
	//TODO implement me
	panic("implement me")
}
