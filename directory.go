package main

import (
	"fmt"
	"os"
	"path"
	"sort"
	"unicode/utf8"
)

type Directories []Directory

type Directory struct {
	name     string
	size     Size
	children Directories
}

func (d *Directory) addChild(dir Directory) {
	d.children = append(d.children, dir)
	d.size += dir.size
}

func New(directoryPath string) Directory {
	return newDir(directoryPath, directoryPath)
}

func newDir(name string, directoryPath string) Directory {
	dir := Directory{name, 0, nil}

	f, err := os.Open(directoryPath)
	if err != nil {
		return dir
	}

	fileInfos, err := f.Readdir(-1)
	f.Close()
	if err != nil {
		return dir
	}

	for _, f := range fileInfos {
		dir.size += Size(f.Size())
		if f.IsDir() {
			childDir := newDir(f.Name(), path.Join(directoryPath, f.Name()))
			dir.addChild(childDir)
		}
	}

	sort.Sort(dir.children)
	return dir
}

func (d Directories) Len() int {
	return len(d)
}

func (d Directories) Less(i, j int) bool {
	return d[i].size > d[j].size
}

func (d Directories) Swap(i, j int) {
	d[i], d[j] = d[j], d[i]
}

func (d Directory) Traverse(visitor Visitor) {
	visitor.Visit(d)
	for _, sd := range d.children {
		sd.Traverse(visitor)
	}
}

func (d Directory) Printing(maxDepth int, minSize Size) {
	p := printing{"", 0, maxDepth, minSize}
	d.TraverseTree(&p)
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
		visitor.GoUp(d, isLast)
	}
}

type printing struct {
	praefix  string
	depth    int
	maxDepth int
	minSize  Size
}

func (p *printing) Visit(d Directory, isLast bool) bool {
	if p.depth > 0 {
		fmt.Print(p.praefix)
		if isLast {
			fmt.Print("┣ ")
		} else {
			fmt.Print("┗ ")
		}
	}
	fmt.Println(d.name, d.size)
	return p.depth < p.maxDepth && d.size > p.minSize
}

func (p *printing) GoDown(d Directory, isLast bool) {
	if p.depth > 0 {
		if isLast {
			p.praefix += "  "
		} else {
			p.praefix += "┃ "
		}
	}
	p.depth++
}
func (p *printing) GoUp(d Directory, isLast bool) {
	for i := 0; i < 3; i++ {
		_, l := utf8.DecodeLastRuneInString(p.praefix)
		p.praefix = p.praefix[:len(p.praefix)-l]
	}
	p.depth--
}
