package tools

import (
	"fmt"
	"time"
)

// 生成随机名称函数
func GenerateName(perfix string) string {
	//todo 默认前缀放在配置文件里面可配置
	now := time.Now()

	// 生成毫秒级时间戳
	timestampMilliseconds := now.UnixMilli()
	// var result string = perfix + fmt.Sprintf("%d", timestampMilliseconds)
	return perfix + fmt.Sprintf("%d", timestampMilliseconds)
}
