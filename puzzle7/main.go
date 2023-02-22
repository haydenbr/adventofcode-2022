package main

import (
	"bufio"
	"fmt"
	"haydenbr/adventofcode-2022/util"
	"os"
	"strconv"
	"strings"
)

/*
$ - begins a command line
cd - does as expected, possibilities: "/", "..", or <dir_name>
ls - lists everything in cwd
	- "dir <dir_name>" is a subdirectory of cwd
	- "<file_size> <filename>" is a file in cwd
*/

type File struct {
	Name string
	Size int
}

func NewFile(fileName string, fileSize int) *File {
	newFile := new(File)

	newFile.Name = fileName
	newFile.Size = fileSize

	return newFile
}

type Dir struct {
	Name     string
	Parent   *Dir
	Children []*Dir
	Files    []*File
}

type FileSystem struct {
	Root *Dir
	Cwd  *Dir
}

func NewDir(name string, parent *Dir) *Dir {
	newDir := new(Dir)

	newDir.Name = name
	newDir.Parent = parent
	newDir.Children = make([]*Dir, 0, 1)
	newDir.Files = make([]*File, 0, 1)

	return newDir
}

func (dir *Dir) Size() int {
	return util.SumWith(dir.Files, func(file *File) int {
		return file.Size
	})
}

func (dir *Dir) SizeR() int {
	return dir.Size() + util.SumWith(dir.Children, func(d *Dir) int {
		return d.SizeR()
	})
}

func (dir *Dir) GetSubDir(name string) *Dir {
	var subDir *Dir = nil

	for _, child := range dir.Children {
		if child.Name == name {
			subDir = child
			break
		}
	}

	return subDir
}

func (dir *Dir) HasSubDir(name string) bool {
	return dir.GetSubDir(name) != nil
}

func (fileSystem *FileSystem) ChangeDirectory(newDir string) {
	switch newDir {
	case "":
		panic("cant change directory to empty string")
	case "/":
		fileSystem.Cwd = fileSystem.Root
	case "..":
		fileSystem.Cwd = fileSystem.Cwd.Parent
	default:
		fileSystem.Cwd = fileSystem.Cwd.GetSubDir(newDir)
	}
}

func (fileSystem *FileSystem) MkDir(newDir string) {
	if fileSystem.Cwd.HasSubDir(newDir) {
		panic("directory already exists: " + newDir)
	}

	(*fileSystem).Cwd.Children = append((*fileSystem).Cwd.Children, NewDir(newDir, fileSystem.Cwd))
}

func (fileSystem *FileSystem) AddFile(fileName string, fileSize string) {
	_fileSize, parseErr := strconv.Atoi(fileSize)

	if parseErr != nil {
		panic(parseErr)
	}

	(*fileSystem).Cwd.Files = append((*fileSystem).Cwd.Files, NewFile(fileName, _fileSize))
}

func createFileSystem() *FileSystem {
	inputFile, err := os.Open("input.txt")
	defer inputFile.Close()

	if err != nil {
		panic(err)
	}

	inputScanner := bufio.NewScanner(inputFile)
	inputScanner.Split(bufio.ScanLines)

	fileSystem := new(FileSystem)
	fileSystem.Root = NewDir("/", nil)

	for inputScanner.Scan() {
		nextLine := inputScanner.Text()

		if util.IsWhiteSpace(nextLine) {
			break
		}

		tokens := strings.Split(nextLine, " ")
		firstToken := tokens[0]
		isCommand := firstToken == "$"

		if isCommand {
			command := tokens[1]

			if command == "ls" {
				continue
			}

			if command == "cd" {
				fileSystem.ChangeDirectory(tokens[2])
			}
		} else {
			if firstToken == "dir" {
				if !fileSystem.Cwd.HasSubDir(firstToken) {
					fileSystem.MkDir(firstToken)
				}
			} else {
				fileSystem.AddFile(tokens[1], firstToken)
			}
		}
	}

	return fileSystem
}

func main() {
	fileSystem := createFileSystem()

	fileSystem.ChangeDirectory("/")
	fmt.Println(fileSystem.Root.Children)
}
