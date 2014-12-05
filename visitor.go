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
