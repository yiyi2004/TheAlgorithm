package topological

type Node struct {
	Value    int
	In       int
	Out      int
	Children []Node
	Edges    []Edge
}

type Edge struct {
	From   Node
	End    Node
	Weight int
}

type Graph struct {
	Nodes map[int]Node
	Edges []Edge
}
