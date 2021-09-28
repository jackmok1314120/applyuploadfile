package utils

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

// CreateDateDir
// basePath 是固定目录路径
func CreateDateDir(basePath string) (dirPath string) {

	folderPath := filepath.Join(basePath)
	if _, err := os.Stat(folderPath); os.IsNotExist(err) {
		// 必须分成两步
		// 先创建文件夹
		os.Mkdir(folderPath, 0777)
		// 再修改权限
		os.Chmod(folderPath, 0777)
	}
	return folderPath
}

// Ext
// 获取文件的扩展名
func Ext(path string) string {
	for i := len(path) - 1; i >= 0 && path[i] != '/'; i-- {
		if path[i] == '.' {
			return path[i:]
		}
	}
	return ""
}

// Exists 判断所给路径文件/文件夹是否存在
func Exists(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

// IsDir 判断所给路径是否为文件夹
func IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

// IsFile 判断所给路径是否为文件
func IsFile(path string) bool {
	return !IsDir(path)
}

func MD5File(filename string) (string, error) {
	pFile, err := os.Open(filename)
	if err != nil {
		fmt.Printf("打开文件失败，filename=%v, err=%v", filename, err)
		return "", err
	}
	defer pFile.Close()
	md5h := md5.New()
	io.Copy(md5h, pFile)
	strMd5 := hex.EncodeToString(md5h.Sum(nil))
	return strMd5, err
}

func VailDataFileMd5(path, filename string) bool {
	cah := CacheConf.CacheUtil
	cah.DeleteExpired()
	md5h, err := MD5File(path + "/" + filename)
	if err != nil {
		fmt.Printf("失败，filename=%v, err=%v", filename, err)
		return false
	}
	_, ok := cah.Get(md5h)
	if ok {
		return false
	} else {
		cah.Set(md5h, path, 30*time.Second)
		return true
	}
}

func GetAllFile(pathname string) ([]string, error) {

	var s []string
	rd, err := ioutil.ReadDir(pathname)
	if err != nil {
		fmt.Println("read dir fail:", err)
		return s, err
	}
	for _, fi := range rd {
		if !fi.IsDir() {
			fullName := pathname + "/" + fi.Name()
			s = append(s, fullName)
		}
	}

	return s, nil
}
