package main

import (
	"fmt"
	"strings"
)

// Node is a node in a Directed Acyclic Graph (DAG)
type Node struct {
	name      string
	parents   []string
	ancestors map[string]*Node
}

// DAG is a Directed Acyclic Graph
type DAG struct {
	nodes       map[string]*Node
	leaves      map[string]*Node
	bisectors   map[string]*Node
	biscetorMin int
}

// NewDAG returns a pointer to a new DAG structure. Returns an error if the provided
//   input is invalid
func NewDAG(nodes []string) (*DAG, error) {
	dag := &DAG{}
	dag.nodes = make(map[string]*Node)
	dag.leaves = make(map[string]*Node)
	dag.bisectors = make(map[string]*Node)
	for _, v := range nodes {
		node := newNode(v)
		err := dag.addNode(node)
		if err != nil {
			return nil, err
		}
	}
	dag.biscetorMin = len(dag.nodes)
	dag.findBisectors()
	return dag, nil
}

func newNode(s string) *Node {
	node := &Node{}
	splits := strings.Split(s, ":")
	node.name = strings.TrimSpace(splits[0])
	if len(splits) > 1 && splits[1] != "" {
		parents := strings.Split(splits[1], ",")
		for _, parent := range parents {
			node.parents = append(node.parents, strings.TrimSpace(parent))
		}
	}
	node.ancestors = make(map[string]*Node)
	node.ancestors[node.name] = node
	return node
}

func (n *Node) getName() string {
	return n.name
}

func (d *DAG) addNode(n *Node) error {
	d.nodes[n.getName()] = n
	d.leaves[n.getName()] = n

	if len(n.parents) == 0 {
		return nil
	}
	for _, parent := range n.parents {
		//validate parents
		if _, present := d.nodes[parent]; present {
			// append node as ancestor to parent nodes
			d.nodes[parent].ancestors[parent] = n
		} else {
			// prevent cyclic references and non-existent node references
			return fmt.Errorf("Invalid Map, node %s contains parent (%s) not in map yet", n.getName(), parent)
		}
		//trim leaves
		if _, present := d.leaves[parent]; present {
			delete(d.leaves, parent)
		}
	}
	return nil
}

// GetAncestors returns a string of the Ancestors for the named node. Returns an error
//   if the node is not present in the DAG
func (d DAG) GetAncestors(node string) (string, error) {
	if n, ok := d.nodes[node]; ok {
		keys := make([]string, len(n.ancestors))
		i := 0
		for k := range n.ancestors {
			keys[i] = k
			i++
		}
		return "Ancestors of " + d.nodes[node].name + ": " + strings.Join(keys, ", "), nil
	}
	return "", fmt.Errorf("Node %s not present in DAG", node)
}

// GetLeaves returns a string with the list of Leaves in the DAG stucture
func (d DAG) GetLeaves() string {
	result := make([]string, len(d.leaves))
	i := 0
	for k := range d.leaves {
		result[i] = k
		i++
	}
	return "Leaves: " + strings.Join(result, ", ")
}

// GetBisectors returns a string with the list of Bisectors in the DAG structure
func (d DAG) GetBisectors() string {
	result := make([]string, len(d.bisectors))
	i := 0
	for k := range d.bisectors {
		result[i] = k
		i++
	}
	return "Bisectors: " + strings.Join(result, ", ")
}

func (d *DAG) findBisectors() {
	N := len(d.nodes)
	for _, v := range d.nodes {
		A := len(v.ancestors)
		Min := min(A, N-A)
		if Min < d.biscetorMin {
			d.biscetorMin = Min
			d.bisectors = map[string]*Node{v.getName(): v}
		} else if Min == d.biscetorMin {
			d.bisectors[v.getName()] = v
		}
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
