package algorithm


type KeyType int
type Status bool
type BiTree struct {
	data 			KeyType
	left, right    *BiTree
}

func SearchBST(T *BiTree, key KeyType) (Status) {
	if T == nil {
		return false
	} else if key == T.data {
		return true
	} else if key < T.data {
		return SearchBST(T.left, key)
	} else {
		return SearchBST(T.right, key)
	}
}

func AddBST(T *BiTree, value KeyType) *BiTree {
	if T == nil {
		T = new(BiTree)
		T.data = value
		return T
	}

	if value > T.data {
		T.right = AddBST(T.right, value)
	} else {
		T.left = AddBST(T.left, value)
	}

	return T
}
