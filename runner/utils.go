package runner

import (
	"os"
	"path/filepath"
	"strings"
)

func initFolders() {
	runnerLog("InitFolders")
	path := tmpPath()
	runnerLog("mkdir %s", path)
	err := os.Mkdir(path, 0755)
	if err != nil {
		runnerLog(err.Error())
	}
}

func isTmpDir(path string) bool {
	absolutePath, _ := filepath.Abs(path)
	absoluteTmpPath, _ := filepath.Abs(tmpPath())

	return absolutePath == absoluteTmpPath
}

func isIgnoredFolder(path string) bool {
	paths := strings.Split(path, "/")
	if len(paths) <= 0 {
		return false
	}

	for _, e := range strings.Split(settings["ignored"], ",") {
		if strings.TrimSpace(e) == paths[0] {
			return true
		}
	}
	return false
}

func isWatchedFile(path string) bool {
	absolutePath, _ := filepath.Abs(path)
	absoluteTmpPath, _ := filepath.Abs(tmpPath())

	if strings.HasPrefix(absolutePath, absoluteTmpPath) {
		return false
	}

	ext := filepath.Ext(path)

	for _, e := range strings.Split(settings["watch_ext"], ",") {
		if strings.TrimSpace(e) == ext {
			return true
		}
	}

	return false
}

func shouldRebuild(eventName string) bool {
	if eventName == "/" {
		return true
	}

	for _, e := range strings.Split(settings["rebuild_ext"], ",") {
		e = strings.TrimSpace(e)
		fileName := strings.Replace(strings.Split(eventName, ":")[0], `"`, "", -1)
		if strings.HasSuffix(fileName, e) {
			return true
		}
	}

	return false
}

func shouldGoGenerate(eventName string) bool {
	if eventName == "/" {
		return true
	}

	for _, e := range strings.Split(settings["gogenerate_ext"], ",") {
		e = strings.TrimSpace(e)
		fileName := strings.Replace(strings.Split(eventName, ":")[0], `"`, "", -1)
		if strings.HasSuffix(fileName, e) {
			return true
		}
	}

	return false
}

func shouldGQLGenerate(eventName string) bool {
	if eventName == "/" {
		return true
	}

	for _, e := range strings.Split(settings["gqlgenerate_ext"], ",") {
		e = strings.TrimSpace(e)
		fileName := strings.Replace(strings.Split(eventName, ":")[0], `"`, "", -1)
		if strings.HasSuffix(fileName, e) {
			return true
		}
	}

	return false
}

func createBuildErrorsLog(message string) bool {
	file, err := os.Create(buildErrorsFilePath())
	if err != nil {
		return false
	}

	_, err = file.WriteString(message)
	if err != nil {
		return false
	}

	return true
}

func removeBuildErrorsLog() error {
	err := os.Remove(buildErrorsFilePath())

	return err
}
