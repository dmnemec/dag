package main

import "fmt"

func main() {
	d1 := []string{"A:", "B: A", "C: B", "D:", "E: D", "F: C, E", "G: F", "H: G"}
	dag1, err := NewDAG(d1)
	if err != nil {
		fmt.Print(err)
	}
	fmt.Println(dag1.GetAncestors("C"))
	fmt.Println(dag1.GetLeaves())
	fmt.Println(dag1.GetBisectors())
}
