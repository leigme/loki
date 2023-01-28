package file

import "io/fs"

type ListWithTime []fs.DirEntry

func (l ListWithTime) Len() int {
	return len(l)
}

func (l ListWithTime) Less(i, j int) bool {
	fsi, _ := l[i].Info()
	fsj, _ := l[j].Info()
	return fsi.ModTime().After(fsj.ModTime())
}

func (l ListWithTime) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

type ListWithName []fs.DirEntry

func (l ListWithName) Len() int {
	return len(l)
}

func (l ListWithName) Less(i, j int) bool {
	return l[i].Name() < l[j].Name()
}

func (l ListWithName) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}
