package data_structure

import (
	"errors"
	"fmt"
	"sync"
)

var (
	ErrorNotInit   = errors.New("tree not init")
	ErrorDuplicate = errors.New("insert duplicate")
)

type BinaryTree interface {
	Insert(e int)     //插入元素
	Delete(e int)     //删除元素
	PreOrder() []int  //前序遍历
	InOrder() []int   //中序遍历
	PostOrder() []int //后序遍历
	Size() (num int)  //返回树中的元素个数
	Clear()           //清空树中所有元素
	Empty() (b bool)  //判断树是否为空
	Graph()           //通过图形化的方式打印出来树的结构
}

type AvlTree struct {
	root    *avlNode
	size    int
	RWMutex sync.RWMutex
}

func NewAvlTree() *AvlTree {
	return &AvlTree{
		root:    nil,
		size:    0,
		RWMutex: sync.RWMutex{},
	}
}
func (a *AvlTree) GetRoot() (*avlNode, int) {
	return a.root, a.root.depth
}

func (a *AvlTree) Insert(e int) error {
	if a == nil {
		return ErrorNotInit
	}
	a.RWMutex.Lock()
	defer a.RWMutex.Unlock()

	if a.root == nil {
		a.root = newAvlNode(e)
		a.size++
		return nil
	}
	newRoot, err := a.root.insert(e)
	if err != nil {
		return err
	}
	a.root = newRoot
	a.size++
	return nil
}

func (a *AvlTree) Delete(value int) (ok bool) {
	if a == nil || a.Empty() {
		return false
	}
	a.RWMutex.Lock()
	defer a.RWMutex.Unlock()
	if a.size == 1 && a.root.key == value {
		a.root = nil
		a.size--
		ok = true
	}
	a.root, ok = a.root.delete(value)
	if ok {
		a.size--
	}
	return ok
}

func (a *AvlTree) PreOrder() []int {
	a.RWMutex.RLock()
	defer a.RWMutex.RUnlock()
	var ret []int
	preOrderTraversal(a.root, &ret)
	return ret
}

func (a *AvlTree) InOrder() []int {
	a.RWMutex.RLock()
	defer a.RWMutex.RUnlock()
	var ret []int
	inOrderTraversal(a.root, &ret)
	return ret
}

func (a *AvlTree) PostOrder() []int {
	a.RWMutex.RLock()
	defer a.RWMutex.RUnlock()
	var ret []int
	postOrderTraversal(a.root, &ret)
	return ret
}

func (a *AvlTree) Size() (num int) {
	return a.size
}

func (a *AvlTree) Clear() {
	a.RWMutex.Lock()
	defer a.RWMutex.Unlock()
	a.root = nil
}

func (a *AvlTree) Empty() (b bool) {
	return a.size == 0
}

type avlNode struct {
	key   int
	depth int
	left  *avlNode
	right *avlNode
}

func newAvlNode(key int) *avlNode {
	return &avlNode{
		key:   key,
		depth: 1,
		left:  nil,
		right: nil,
	}
}

func (a *avlNode) getDepth() (depth int) {
	if a == nil {
		return 0
	}
	return a.depth
}

func (a *avlNode) insert(e int) (res *avlNode, err error) {
	if a == nil {
		return newAvlNode(e), nil
	}
	res = a
	if a.key > e {
		a.left, err = a.left.insert(e)
	} else if a.key < e {
		a.right, err = a.right.insert(e)
	} else {
		return a, ErrorDuplicate
	}
	if err == nil {
		// 插入成功，判断节点是否需要翻转
		res = a.rotate()
	}
	res.depth = MaxCompare(a.left.getDepth(), a.right.getDepth()) + 1
	return res, err

}

// 旋转后返回根节点
func (a *avlNode) rotate() *avlNode {
	diffDepth := a.left.getDepth() - a.right.getDepth()
	if -1 <= diffDepth && diffDepth <= 1 {
		return a
	}
	if diffDepth > 1 {
		//左边高
		if a.left.left != nil {
			// 左左子树，右旋（以左子树节点为中心，当前节点右旋）
			return a.rightRotate()
		} else {
			// 左右子树，左右旋（1. 以左左子树节点为中心，将左子树节点左旋 2. 以左子树节点为中心，当前节点右旋）
			return a.leftRightRotate()
		}

	} else {
		//右边高
		if a.right.right != nil {
			// 右右子树，左旋
			return a.leftRotate()
		} else {
			// 右左子树，右左旋
			return a.rightLeftRotate()
		}
	}

}

func (a *avlNode) leftRotate() *avlNode {
	pivot := a.right
	a.right = pivot.left
	pivot.left = a

	// 计算深度的时候，先计算当前节点，再计算pivot节点
	a.depth = MaxCompare(a.left.getDepth(), a.right.getDepth()) + 1
	pivot.depth = MaxCompare(pivot.getDepth(), a.left.getDepth())
	return pivot
}

func (a *avlNode) rightRotate() *avlNode {
	pivot := a.left
	a.left = pivot.right
	pivot.right = a
	// 计算深度的时候，先计算当前节点，再计算pivot节点
	a.depth = MaxCompare(a.left.getDepth(), a.right.getDepth()) + 1
	pivot.depth = MaxCompare(pivot.getDepth(), a.left.getDepth())
	return pivot
}

func (a *avlNode) leftRightRotate() *avlNode {
	a.left = a.left.leftRotate()
	res := a.rightRotate()
	return res
}

func (a *avlNode) rightLeftRotate() *avlNode {
	a.right = a.right.rightRotate()
	res := a.leftRotate()
	return res
}

func (a *avlNode) delete(value int) (node *avlNode, ok bool) {
	if a == nil {
		return nil, false
	}
	res := a
	if value < a.key {
		a.left, ok = a.left.delete(value)
	} else if value > a.key {
		a.right, ok = a.right.delete(value)
	} else {
		ok = true
		if a.left != nil && a.right != nil {
			a.key = a.getMin()
			res, ok = a.right.delete(a.key)
		} else if a.left != nil {
			res = a.left
		} else {
			// 只有右边或者两边都没有的情况
			res = a.right
		}
	}
	if res != nil {
		res.depth = max(res.left.getDepth(), res.right.getDepth()) + 1
		res = res.rotate()
	}

	return res, ok
}

func (a *avlNode) getMin() int {
	if a.left == nil {
		return a.key
	} else {
		return a.left.getMin()
	}
}

func preOrderTraversal(n *avlNode, ret *[]int) {
	if n == nil {
		return
	}
	*ret = append(*ret, n.key)
	preOrderTraversal(n.left, ret)
	preOrderTraversal(n.right, ret)
}
func inOrderTraversal(n *avlNode, ret *[]int) {
	if n == nil {
		return
	}
	inOrderTraversal(n.left, ret)
	*ret = append(*ret, n.key)
	inOrderTraversal(n.right, ret)
}
func postOrderTraversal(n *avlNode, ret *[]int) {
	if n == nil {
		return
	}
	postOrderTraversal(n.left, ret)
	postOrderTraversal(n.right, ret)
	*ret = append(*ret, n.key)
}

// 为了实现方便，该图向左旋转了90度
func (a *AvlTree) Graph() {

	if a == nil || a.root == nil {
		fmt.Println("The tree is empty or not initialized.")
		return
	}
	PrintTree(a.root, "", true)
}

// PrintTree 以图形化的方式通过中序遍历打印树的结构。
func PrintTree(node *avlNode, prefix string, isTail bool) {
	if node != nil {
		PrintTree(node.right, prefix+ifThenElse(isTail, "│   ", "    "), false)      // 先打印右子树，保证左子节点在前，右子节点在后。
		fmt.Printf("%s%s─ %d\n", prefix, ifThenElse(isTail, "└──", "├──"), node.key) // 根据 isTail 决定连接符。
		PrintTree(node.left, prefix+ifThenElse(isTail, "    ", "│   "), true)        // 再打印左子树。
	} else { // 如果节点为空，则打印一个空格占位符。
		fmt.Printf("%s%s\n", prefix, ifThenElse(isTail, "│", "")) // 根据 isTail 决定连接符。
	}
}

// ifThenElse 辅助函数，根据条件选择不同的字符串。
func ifThenElse(condition bool, trueVal string, falseVal string) string {
	if condition {
		return trueVal
	} else {
		return falseVal
	}
}
