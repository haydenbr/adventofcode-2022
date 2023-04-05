package main

import (
	"bufio"
	"fmt"
	"haydenbr/adventofcode-2022/util"
	"math"
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
	Name    string
	Size    int
	Parent  *Dir
	SubDirs []*Dir
	Files   []*File
}

type FileSystem struct {
	Root    *Dir
	Cwd     *Dir
	DirRefs []*Dir
}

func NewDir(name string, parent *Dir) *Dir {
	newDir := new(Dir)

	newDir.Name = name
	newDir.Parent = parent
	newDir.SubDirs = make([]*Dir, 0)
	newDir.Files = make([]*File, 0)

	return newDir
}

func sumFileSizes(files []*File) int {
	return util.SumWith(files, func(file *File) int {
		return file.Size
	})
}

func sumDirSizes(dirs []*Dir) int {
	return util.SumWith(dirs, func(dir *Dir) int {
		return dir.Size
	})
}

// func (dir *Dir) SumDirsBelowSize(size int) int {
// 	if !dir.HasSubDirs() {
// 		if dir.Size <= size {
// 			return dir.Size
// 		} else {
// 			return 0
// 		}
// 	}

// }
func (dir *Dir) GetSubDir(name string) *Dir {
	var subDir *Dir = nil

	for _, child := range dir.SubDirs {
		if child.Name == name {
			subDir = child
			break
		}
	}

	return subDir
}

func (dir *Dir) HasSubDirs() bool {
	return len(dir.SubDirs) > 0
}

func (dir *Dir) HasSubDir(name string) bool {
	return dir.GetSubDir(name) != nil
}

func (fs *FileSystem) ChangeDirectory(newDir string) {
	switch newDir {
	case "":
		panic("cant change directory to empty string")
	case "/":
		fs.Cwd = fs.Root
	case "..":
		fs.Cwd = fs.Cwd.Parent
	default:
		fs.Cwd = fs.Cwd.GetSubDir(newDir)
	}
}

func (fs *FileSystem) MkDir(dirName string) {
	if fs.Cwd.HasSubDir(dirName) {
		panic("directory already exists: " + dirName)
	}

	newDir := NewDir(dirName, fs.Cwd)
	fs.Cwd.SubDirs = append(fs.Cwd.SubDirs, newDir)
	fs.DirRefs = append(fs.DirRefs, newDir)
}

func (fs *FileSystem) AddFile(fileName string, fileSize string) {
	_fileSize, parseErr := strconv.Atoi(fileSize)

	if parseErr != nil {
		panic(parseErr)
	}

	newFile := NewFile(fileName, _fileSize)
	fs.Cwd.Files = append(fs.Cwd.Files, newFile)

	for dir := fs.Cwd; dir != nil; dir = dir.Parent {
		dir.Size += newFile.Size
	}
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
	fileSystem.DirRefs = make([]*Dir, 0)
	fileSystem.DirRefs = append(fileSystem.DirRefs, fileSystem.Root)

	for inputScanner.Scan() {
		nextLine := inputScanner.Text()

		if util.IsWhiteSpace(nextLine) {
			break
		}

		tokens := strings.Split(nextLine, " ")
		firstToken := tokens[0]
		secondToken := tokens[1]
		isCommand := firstToken == "$"

		if isCommand {
			command := secondToken

			if command == "ls" {
				continue
			}

			if command == "cd" {
				fileSystem.ChangeDirectory(tokens[2])
			}
		} else {
			target := secondToken
			if firstToken == "dir" {
				if !fileSystem.Cwd.HasSubDir(firstToken) {
					fileSystem.MkDir(target)
				}
			} else {
				fileSystem.AddFile(target, firstToken)
			}
		}
	}

	return fileSystem
}

func main() {
	fileSystem := createFileSystem()
	par1Dirs := util.Filter(fileSystem.DirRefs, func(dir *Dir) bool { return dir.Size <= 100_000 })
	part1 := util.SumWith(par1Dirs, func(dir *Dir) int { return dir.Size })
	fmt.Println("part 1:", part1)

	spaceRemaining := 70_000_000 - fileSystem.Root.Size
	capacityRequired := 30_000_000 - spaceRemaining
	part2 := math.MaxInt

	for _, dir := range fileSystem.DirRefs {
		if dir.Size < part2 && dir.Size >= capacityRequired {
			part2 = dir.Size
		}
	}

	fmt.Println("part 2:", part2)
}
