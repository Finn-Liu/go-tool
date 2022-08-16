package tool

import (
	"crypto/md5"
	"fmt"
	"math/rand"
	"time"
)

func UnixToDate(timestamp int) string {
	t := time.Unix(int64(timestamp), 0)
	return t.Format("2006-01-02 15:04:05")
}

func DateToUnix(str string) int64 {
	template := "2006-01-02 15:04:05"
	t, err := time.ParseInLocation(template, str, time.Local)
	if err != nil {
		return 0
	}
	return t.Unix()
}

func GetDate() string {
	template := "2006-01-02 15:04:05"
	return time.Now().Format(template)
}

func Md5(str string) string {
	data := []byte(str)
	return fmt.Sprintf("%x\n", md5.Sum(data))
}

func Hello(in string) (out string) {
	out = in + "word"
	return
}

//随机字符串
func GetRandomString(n int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyz"
	bytes := []byte(str)
	var result []byte
	for i := 0; i < n; i++ {
		result = append(result, bytes[rand.Intn(len(bytes))])
	}
	return string(result)
}

func SplitLine() {
	fmt.Println("################# 分割辅助线 #################")
}
