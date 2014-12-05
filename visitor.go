package main

type Visitor interface {
	Visit(d Directory)
}

type VisitorFunc func(d Directory)

func (vf VisitorFunc) Visit(d Directory) {
	vf(d)
}

func (d Directory) TraversFunc(f VisitorFunc) {
	d.Traverse(f)
}

type TreeVisitor interface {
	Visit(d Directory, isLast bool) bool
	GoDown(d Directory, isLast bool)
	GoUp(d Directory, isLast bool)
}

func (d Directory) TraverseTree(visitor TreeVisitor) {
	d.traverseTree(visitor, true)
}

func (d Directory) traverseTree(visitor TreeVisitor, isLast bool) {
	goDown := visitor.Visit(d, isLast)
	if goDown && len(d.children) > 0 {
		visitor.GoDown(d, isLast)
		for i, sd := range d.children {
			sd.traverseTree(visitor, i == len(d.children)-1)
		}
		visitor.GoDown(d, isLast)
	}
}
