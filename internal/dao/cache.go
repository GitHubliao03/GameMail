package dao

import (
	"GameMail/internal/model"
	"errors"
	"sync"
)
var G Global

type Global struct {
	sync.Mutex
	UserCacheMap map[int]*model.UserCache
}

func init() {
	G.UserCacheMap = make(map[int]*model.UserCache)
}

// SetGlobal 设置缓存
func (g *Global) SetGlobal(userCache *model.UserCache) {
	g.Lock()
	defer g.Unlock()
	g.UserCacheMap[userCache.Id] = userCache
}

// GetGlobal 获取缓存
func (g *Global) GetGlobal(id int) (*model.UserCache,error){
	g.Lock()
	defer g.Unlock()
	if v, exist := g.UserCacheMap[id]; exist {
		return v,nil
	}else {
		return nil, errors.New("has not cache")
	}
}
// ClearAll 清空缓存
func (g *Global) ClearAll() {
	g.Lock()
	defer g.Unlock()
	g.UserCacheMap = make(map[int]*model.UserCache)
}
