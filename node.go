package mm

// Node represents an item that can exist in an object hierearchy.
// A node is the most basic of objects in the marshmallow library
type Node interface {
    Add(Node)
    Remove(Node) bool

    GetParent() Node
    GetChildren() []Node

    setParent(Node)
}

func IterateDepth(n Node) <-chan Node {

    ch := make(chan Node)

    go func() {

        yieldChildrenDepth(n, ch)
        close(ch)

    }()

    return ch

}

func yieldChildrenDepth(n Node, ch chan Node) {

    ch <- n

    for _, child := range n.GetChildren() {

        yieldChildrenDepth(child, ch)

    }

}
