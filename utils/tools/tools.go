/*
 * @Author: jffan
 * @Date: 2024-07-31 14:18:43
 * @LastEditTime: 2024-07-31 14:42:15
 * @LastEditors: jffan
 * @FilePath: \tcas-gitee\utils\tools\tools.go
 * @Description: 🎉🎉🎉
 */
package tools

import (
	"fmt"
	"time"
)

// A function that generates a random name
func GenerateName(perfix string) string {
	now := time.Now()
	timestampMilliseconds := now.UnixMilli()
	return perfix + fmt.Sprintf("%d", timestampMilliseconds)
}
