package main

import (
	"os"
	"path"
	"sort"
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
