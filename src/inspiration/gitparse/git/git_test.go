package git

import (
	"fmt"
	"github.com/play-code/src/inspiration/gitparse/gitexec"
	"strings"
	"testing"
)

func TestGitCommonGetUserName(t *testing.T) {
	name := GetGitUserName()
	fmt.Println(name)
}

func TestGitCommonDiff(t *testing.T) {
	fmt.Println(gitexec.Command("diff"))
}

func TestGitShowCommandId(t *testing.T) {
	fmt.Println(GetCommitContext("1391204bf3cf38dd946ff2d8d67a22407052eb88"))
}

func TestGetLatestCommit(t *testing.T) {
	fmt.Println(GetLatestCommit())
}

func TestGitDiffChangedFile(t *testing.T) {
	commit := "a9eb79aaf47b9e05af55fab066d274d520de2090"
	s := GitShowChangedFile(commit)
	fmt.Println("1", s)
	diffFiles,_ := gitexec.Command( "show", commit, "--stat=200")
	//fmt.Println(diffFiles)
	context := strings.Split(diffFiles, "\n")
	var fileList []string
	for _, v := range context {
		if strings.Contains(v, "test.go") {
			continue
		}
		if strings.Contains(v, ".go") {
			index := strings.LastIndex(v, ".go")
			v = v[:index+3]
			v = strings.Trim(v, " ")
			fileList = append(fileList, v)
		}
	}
	fmt.Println(fileList)
}

func TestGetChangedFuncs(t *testing.T) {
	commit := "1391204bf3cf38dd946ff2d8d67a22407052eb88"
	context := GetCommitChangedFunc(commit)
	fmt.Println(context)
}