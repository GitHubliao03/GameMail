package dao

import (
	"GameMail/internal/model"
	"errors"
	"sync"
)
var (
	GU GlobalUser
	GM GlobalMail
)

type GlobalUser struct {
	uMutex sync.Mutex
	userCacheMap map[int]*model.UserCache
}
type GlobalMail struct {
	mMutex sync.Mutex
	mailCacheMap map[int]*model.Mail
}

func init() {
	GU.userCacheMap = make(map[int]*model.UserCache)
	GM.mailCacheMap = make(map[int]*model.Mail)
}

type IGlobal interface {
	SetGlobal(interface{}) error
	GetGlobal(int)(interface{},error)
	ClearAll()

}

// SetGlobal 设置缓存
func (g *GlobalUser) SetGlobal(userCache interface{}) error {
	g.uMutex.Lock()
	defer g.uMutex.Unlock()
	v,ok := userCache.(*model.UserCache)
	if !ok {
		return errors.New("obj is not model.UserCache")
	}
	g.userCacheMap[v.Id] = v
	return nil
}
func (g *GlobalMail) SetGlobal(mailCache interface{}) error {
	g.mMutex.Lock()
	defer g.mMutex.Unlock()
	v,ok := mailCache.(*model.Mail)
	if !ok {
		return errors.New("obj is not model.Mail")
	}
	g.mailCacheMap[v.Id] = v
	return nil
}

// GetGlobal 获取缓存
func (g *GlobalUser) GetGlobal(id int) (interface{},error){
	g.uMutex.Lock()
	defer g.uMutex.Unlock()
	if v, exist := g.userCacheMap[id]; exist {
		return v,nil
	}else {
		return nil, errors.New("has not cache")
	}
}
func (g *GlobalMail) GetGlobal(id int) (interface{},error){
	g.mMutex.Lock()
	defer g.mMutex.Unlock()
	if v, exist := g.mailCacheMap[id]; exist {
		return v,nil
	}else {
		return nil, errors.New("has not cache")
	}
}


// ClearAll 清空缓存
func (g *GlobalUser) ClearAll() {
	g.uMutex.Lock()
	defer g.uMutex.Unlock()
	g.userCacheMap = make(map[int]*model.UserCache)
}

func (g *GlobalMail) ClearAll() {
	g.mMutex.Lock()
	defer g.mMutex.Unlock()
	g.mailCacheMap = make(map[int]*model.Mail)
}
