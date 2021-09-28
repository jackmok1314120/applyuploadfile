package utils

import (
	"github.com/patrickmn/go-cache"
	"time"
)

var CacheConf CacheConn

type CacheConn struct {
	CacheUtil *cache.Cache
}

func init() {
	NewCache()
}

func NewCache() *CacheConn {

	if CacheConf.CacheUtil == nil {
		// 创建一个cache对象，默认ttl 5分钟，每10分钟对过期数据进行一次清理
		CacheConf = CacheConn{
			CacheUtil: cache.New(30*time.Minute, 30*time.Minute),
		}
	}
	return &CacheConf
}
