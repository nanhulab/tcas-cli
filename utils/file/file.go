/*
 * @Author: jffan
 * @Date: 2024-07-30 14:23:59
 * @LastEditTime: 2024-07-30 14:29:31
 * @LastEditors: jffan
 * @FilePath: \tcas-cli\utils\file\file.go
 * @Description: 🎉🎉🎉
 */
package file

import (
	"encoding/base64"
	"fmt"
	"io"
	"os"
)

// judge file exists
func IsExists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		return os.IsExist(err)
	}
	return true
}

// judge path is Dir
func IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

// judge path is fiel
func IsFile(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return !s.IsDir()
}

// file to base64
func FileToBase64(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", fmt.Errorf("error opening file: %w", err)
	}
	defer file.Close()
	fileBytes, err := io.ReadAll(file)
	if err != nil {
		return "", fmt.Errorf("error reading file: %w", err)
	}
	base64Encoded := base64.StdEncoding.EncodeToString(fileBytes)
	return base64Encoded, nil
}
