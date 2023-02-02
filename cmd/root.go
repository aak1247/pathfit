/*
Copyright © 2023 aak1247<aak1247@126.com>
*/
package cmd

import (
	"errors"
	"fmt"
	"os"
	"runtime"
	"strings"
)

func fitPath(path string) string {
	if isInGitBash(path) {
		// 如果是 Git Bash 环境，则需要先将路径中自动添加的MSYSTEM_PREFIX路径去掉
		prefix := os.Getenv("MSYSTEM_PREFIX")
		prefix = strings.Replace(prefix, "/mingw64", "", 1)
		prefix = strings.Replace(prefix, "/mingw32", "", 1)
		path = strings.Replace(path, prefix, "", 1)
	}
	var drive = getDrive(path)
	var shortPath = strings.Replace(path, drive+":", "", 1)
	shortPath = strings.Replace(shortPath, "/"+drive, "", 1)
	shortPath = strings.Replace(shortPath, "\\", string(os.PathSeparator), -1)
	shortPath = strings.Replace(shortPath, "/", string(os.PathSeparator), -1)
	if shortPath[len(shortPath)-1] == '"' || shortPath[len(shortPath)-1] == '\'' {
		shortPath = shortPath[:len(shortPath)-1]
	}
	if isInWSL() {
		return "/mnt/" + strings.ToLower(drive) + shortPath
	}
	if isInCmd() {
		if len(drive) == 0 {
			return shortPath
		}
		if drive == "mnt" {
			return strings.ToUpper(shortPath[1:2]) + ":\\" + shortPath[2:]
		}
		return strings.ToUpper(drive) + ":\\" + shortPath
	}
	if isInPosix() {
		return string(os.PathSeparator) + drive + shortPath
	}
	return path
}

func isWindows(path string) bool {
	return (len(path) > 2 && ((path[1] == ':' && path[2] == '\\') ||
		(path[1] == ':' && path[2] == '/'))) ||
		(len(path) > 3 && path[2] == ':' && path[3] == '\\') ||
		(len(path) > 3 && path[2] == ':' && path[3] == '/')
}

func isInGitBash(path string) bool {
	return runtime.GOOS == "windows" && (os.Getenv("MSYSTEM") == "MINGW64" || os.Getenv("MSYSTEM") == "MINGW32")
}

func isUnix(path string) bool {
	return len(path) > 0 && (path[0] == '/' ||
		(len(path) > 1 && (path[0] == '"' || path[0] == '\'') && path[1] == '/'))
}

func isInCmd() bool {
	return runtime.GOOS == "windows"
}

func isInPosix() bool {
	return runtime.GOOS == "linux" || runtime.GOOS == "darwin"
}

func isInWSL() bool {
	return isInPosix() && os.Getenv("WSL_DISTRO_NAME") != ""
}

func getDrive(path string) string {
	// 如果是 Windows 系统
	if isWindows(path) {
		var drive = strings.Split(path, ":")[0]
		return strings.Replace(drive, "\"", "", -1)
	}

	// 如果是 Unix 系统，则根据分隔符获取路径的第一段字符串
	if isUnix(path) {
		segments := strings.Split(path, "/")
		for _, segment := range segments {
			if len(segment) > 0 {
				return segment
			}
		}
	}

	return ""
}

func execute() error {
	// for k, v := range os.Args {
	// 	fmt.Printf("args[%v]=[%v]\n", k, v)
	// }
	var args = os.Args[1:]
	if len(args) == 0 {
		return errors.New("Please input the path")
	}
	var path = args[0]
	// fmt.Println("args: ", args)
	// fmt.Println("path: " + path)
	if len(args) == 1 && (args[0] == "-h" || args[0] == "--help") {
		fmt.Println(`Easily translate path to fit the current runtime in all kinds of runtime env. 
Usage: pathfit [path]

eg: pathfit /d/Users/aak1247/Repos`)
	}
	// 从路径中通过正则匹配出盘符
	fmt.Print(fitPath(path))
	return nil
}

func Execute() {
	err := execute()
	if err != nil {
		os.Exit(1)
	}
}
