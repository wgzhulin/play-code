package git

import (
	"github.com/play-code/src/inspiration/gitparse/gitexec"
	"log"
	"strings"
)

func GetGitUserName() string {
	result,err := gitexec.Command( "config", "--global", "user.name")
	if err != nil {
		log.Println("xxx")
		return ""
	}

	return result
}

func GetCommitContext(commitId string) string {
	result,err := gitexec.Command( "show", commitId)
	if err != nil {
		log.Println("xxx")
		return ""
	}

	return result
}

func GetLatestCommit() string {
	result,err := gitexec.Command( "log", "-1")
	if err != nil {
		log.Println("xxx")
		return ""
	}

	if !strings.Contains(result, "commit") {
		return ""
	}

	split := strings.Split(result, "\n")
	commit := strings.Split(split[0], " ")[1]

	return commit
}

func GetCommitChangedFunc(commitId string) []string {
	result,err := gitexec.Command( "show", commitId)
	if err != nil {
		return nil
	}
	context := strings.Split(result, "\n")

	var funcList []string
	for _, v := range context {
		if strings.Contains(v, "func") {
			funcList = append(funcList, v)
		}
	}

	return funcList
}

func GitShowChangedFile(commitId string) []string {
	diffFiles,_ := gitexec.Command( "show", commitId, "--stat=200")
	// fmt.Println(diffFiles)
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
	return fileList
}

