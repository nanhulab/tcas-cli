/*
 * @Author: jffan
 * @Date: 2024-07-30 14:23:59
 * @LastEditTime: 2024-08-15 09:24:19
 * @LastEditors: jffan
 * @FilePath: \gitee-tcas\utils\file\file.go
 * @Description:
 */
package file

import (
	"encoding/base64"
	"encoding/json"
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

// Determine whether the directory exists, and if it does not exist, create it
func EnsureDirExists(dirPath string) error {
	_, err := os.Stat(dirPath)
	if os.IsNotExist(err) {
		// Create a directory if it doesn't exist
		err = os.MkdirAll(dirPath, 0755)
		if err != nil {
			return fmt.Errorf("error creating directory: %v", err)
		}
	} else if err != nil {
		return fmt.Errorf("error checking directory: %v", err)
	}
	return nil
}

// read json file
func ReadJSONFile(filePath string) (json.RawMessage, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("opening file failed: %w", err)
	}
	defer file.Close()
	fileData, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("reading file failed: %w", err)
	}
	var jsonRawData json.RawMessage
	err = json.Unmarshal(fileData, &jsonRawData)
	if err != nil {
		return nil, fmt.Errorf("unmarshalling json failed: %w", err)
	}
	return fileData, nil
}
