package libs

import (
	"crypto/md5"
	"fmt"
	"math/rand"
	"time"
)

func Md5(buf []byte) string  {
	hash := md5.New()
	hash.Write(buf)
	fmt.Printf("%x",hash.Sum(nil))
	return fmt.Sprintf("%x", hash.Sum(nil))
}

func Password(len int, pwdO string) (pwd string, salt string) {
	salt = GetRandomString(4)
	defaultPwd := "123456"
	if pwdO != "" {
		defaultPwd = pwdO
	}
	pwd = Md5([]byte(defaultPwd + salt))
	return pwd, salt
}

//生成随机字符串
func GetRandomString(lens int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < lens; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}
