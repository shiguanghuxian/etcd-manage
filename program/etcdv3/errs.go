package etcdv3

import "errors"

// 错误类型
var (
	ErrorPutKey      = errors.New("key is not under a directory or key is a directory or key is not empty")
	ErrorKeyNotFound = errors.New("key has not been set")
	ErrorListKey     = errors.New("can only list a directory")
)
