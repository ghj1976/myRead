package weibo

import (
	"github.com/kardianos/osext"
	"io/ioutil"
	"log"
	"path"
)

// 给指定的文件写内容
func WriteFile(filename, context string) error {
	return ioutil.WriteFile(filename, []byte(context), 0700)
}

// 在当前执行目录下些文件
func WriteFile2ExecutableFolder(shortFilename, context string) error {
	dir, err := osext.ExecutableFolder()
	if err != nil {
		log.Println("osext.ExecutableFolder()", err)
	}

	filename := path.Join(dir, shortFilename)

	return WriteFile(filename, context)
}

// 在当前执行目录下些文件
func WriteFile2ExecutableFolder2(shortFilename string, text []byte) error {
	dir, err := osext.ExecutableFolder()
	if err != nil {
		log.Println("osext.ExecutableFolder()", err)
	}

	filename := path.Join(dir, shortFilename)

	return ioutil.WriteFile(filename, text, 0700)
}
