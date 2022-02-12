package binary_search_tree

// 配列を受け取り、binaryTreeを返す
func BinarySearchTree(n int, input []int) []int {
	tree := make([]int, 0)
	depth := 0
	for i := 0; i < n; i++ {
		tree, depth = appendTree(depth, tree, input[i])
	}
	return tree
}

// treeがあって、それに数値を追加
func appendTree(depth int, tree []int, appended int) ([]int, int) {
	if len(tree) == 0 {
		return []int{appended}, 1
	}

	cmpInd := 0
	cmp := tree[cmpInd]
	cur_dep := 1

	tmp := tree
	tree = make([]int, kaizyo(2)+1)
	for i := 0; i < len(tree); i++ {
		if i < len(tmp) {
			tree[i] = tmp[i]
		} else {
			tree[i] = -1
		}
	}

	for {
		if cmp >= appended {
			cmpInd = cmpInd*2 + 1
		} else {
			cmpInd = cmpInd*2 + 2
		}
		if tree[cmpInd] == -1 {
			tree[cmpInd] = appended
			return tree, cur_dep
		}
		// また要素があた時
		cur_dep++
		if cur_dep > depth {
			tmp := tree
			tree = make([]int, kaizyo(cur_dep+1)+1)
			for i := 0; i < len(tmp); i++ {
				tree[i] = tmp[i]
			}
			for i := len(tmp); i < len(tree); i++ {
				tree[i] = -1
			}
		}
	}
}

func kaizyo(n int) int {
	if n == 0 {
		return 1
	}
	return kaizyo(n-1) * n
}
