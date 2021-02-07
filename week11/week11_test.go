package week11

import (
	"crypto/md5"
	"encoding/hex"
	"testing"
)
//md5算法
func encrypt(pwd string) string {
	m := md5.Sum([]byte(pwd))
	return hex.EncodeToString(m[:])
}
//2次md5 增加穷举难度
func makePwd(userId int,pwd string) string{
	return encrypt(encrypt(pwd+string(userId)))
}
//校验
func checkPW(userId int,pwd string,encryptStr string) bool  {
	s := makePwd(userId, pwd)
	return s == encryptStr
}
func TestWeek11Job(t *testing.T) {
	s := makePwd(12, "1234")
	t.Log(checkPW(12,"1234",s))
	t.Log(checkPW(12,"12345",s))
}