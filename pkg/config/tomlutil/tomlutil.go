// Package tomlutil provides helper functions for https://github.com/pelletier/go-toml.
package tomlutil

import (
	"sort"

	"github.com/pelletier/go-toml"
)

// VisitLeaves walks a TOML tree and executes the process function on each leaf.
// If tree is the root tree, keys must be an empty slice. If tree is a subtree,
// keys is its current path.
func VisitLeaves(tree *toml.Tree, keys []string, process func(leaf interface{}, keys []string)) {
	treeKeys := tree.Keys()
	sort.Strings(treeKeys)

	for _, key := range treeKeys {
		nodeKeys := append(keys, key)
		v := tree.Get(key)
		switch value := v.(type) {
		case *toml.Tree:
			// the node is a subtree, walk deeper
			VisitLeaves(value, nodeKeys, process)
		default:
			// this is a leaf, process it
			process(value, nodeKeys)
		}
	}
}
