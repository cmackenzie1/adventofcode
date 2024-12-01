package days

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/cmackenzie1/adventofcode/v2022/input"
)

type INode struct {
	parent   *INode
	Name     string
	Bytes    int
	children []*INode
}

func (n *INode) String() string {
	sb := strings.Builder{}
	if n.IsDirectory() {
		sb.WriteString(fmt.Sprintf("%*s - %s (dir)\n", n.Level(), " ", n.Name))
		for _, i := range n.children {
			sb.WriteString(i.String())
		}
	} else {
		sb.WriteString(fmt.Sprintf("%*s - %s (file, size=%d)\n", n.Level(), " ", n.Name, n.Bytes))
	}
	return sb.String()
}

func (n *INode) IsDirectory() bool {
	return n.children != nil
}

func (n *INode) HasChild(s string) bool {
	for _, c := range n.children {
		if c.Name == s {
			return true
		}
	}
	return false
}

func (n *INode) Directories() []*INode {
	dirs := make([]*INode, 0)
	for _, i := range n.children {
		if i.IsDirectory() {
			t := i
			dirs = append(dirs, t)
		}
	}
	return dirs
}

func (n *INode) Files() []*INode {
	files := make([]*INode, 0)
	for _, i := range n.children {
		if !i.IsDirectory() {
			t := i
			files = append(files, t)
		}
	}
	return files
}

func (n *INode) Level() int {
	curr := n
	level := 2
	for curr.parent != nil {
		level += 2
		curr = curr.parent
	}
	return level
}

func NewFile(parent *INode, name string, size int) *INode {
	return &INode{parent: parent, Name: name, Bytes: size}
}

func NewDirectory(parent *INode, name string) *INode {
	dir := &INode{parent: parent, Name: name, children: make([]*INode, 0)}
	if parent != nil {
		parent.children = append(parent.children, dir)
	}
	return dir
}

func Day7Part1(path string) {
	lines, err := input.ReadLines(path)
	if err != nil {
		log.Fatal(err)
	}

	root := NewDirectory(nil, "/")
	var curr *INode = nil
	for i := 0; i < len(lines); i++ {
		args := strings.Split(lines[i], " ")
		log.Printf("input = %#v", args)
		switch args[0] {
		case "$":
			switch args[1] {
			case "cd":
				if args[2] == "/" {
					curr = root
					break
				} else if args[2] == ".." {
					curr = curr.parent
					break
				}
				curr, err = cd(curr, args...)
				if err != nil {
					log.Fatal(err)
				}
			case "ls":
				contents := make([][]string, 0)
				j := i + 1 // start at the next line and collect lines until the next $ cmd
				log.Printf("current line = %q, next line = %q", lines[i], lines[j])
				log.Printf("lines = %d, i = %d, j = %d", len(lines), i, j)
				for j <= len(lines)-1 && !strings.HasPrefix(lines[j], "$") {
					contents = append(contents, strings.Split(lines[j], " "))
					j += 1
				}
				i += len(contents)
				log.Printf("dir %s contents: %#v", curr.Name, contents)
				ls(curr, contents)
			}
		}
	}

	fmt.Println(root)
}

// cd changes the curr pointer to the target directory. If
// the target directory does not exist, return an error
// args are the input command line such as ["$", "cd", "/"]
func cd(curr *INode, args ...string) (*INode, error) {
	//["$", "cd", "/"]
	_, _, dst := args[0], args[1], args[2]
	for _, dir := range curr.Directories() {
		if dir.Name == dst {
			t := dir // temp clone
			curr = t
			return curr, nil
		}
	}
	return nil, fmt.Errorf("directory does not exist: %s", dst)
}

// ls will create (if not exists) any files or directories
// based on the reported contents.
// Example contents are [["dir", "a"], ["29116", "f"]]
func ls(curr *INode, contents [][]string) {
	for _, c := range contents {
		log.Printf("children = %v", curr.children)
		if curr.HasChild(c[1]) {
			continue
		}
		switch c[0] {
		case "dir":
			NewDirectory(curr, c[1])
			log.Printf("added dir %s to %s", c[1], curr.Name)
		default: // is a file
			size, err := strconv.Atoi(c[0])
			if err != nil {
				log.Fatalf("unable to parse input: %#v", c)
			}

			curr.children = append(curr.children, NewFile(curr, c[1], size))
			log.Printf("added file %s with size %d to %s", c[1], size, curr.Name)
		}
	}
}
