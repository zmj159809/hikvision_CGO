package hikvision_CGO

import (
	"sync"
)

type ObjectId int32

var refs struct {
	sync.Mutex
	objs map[ObjectId]IF_fMEssCallBack
	next ObjectId
}

func init() {
	refs.Lock()
	defer refs.Unlock()

	refs.objs = make(map[ObjectId]IF_fMEssCallBack)
	refs.next = 1000
}

// NewObjectId 本质是个id分配器,
// 此处的id并非设备返回的用户id，而是传入到回调函数里面的自定义id,用于注册对应回调函数时，调用的不同id使用不同的函数方法
// obj 是 IF_fMEssCallBack 接口的实现
func NewObjectId(obj IF_fMEssCallBack) ObjectId {
	refs.Lock()
	defer refs.Unlock()

	id := refs.next
	refs.next++

	refs.objs[id] = obj
	return id
}

func (id *ObjectId) IsNil() bool {
	return *id == 0
}

func (id *ObjectId) Get() IF_fMEssCallBack {
	refs.Lock()
	defer refs.Unlock()

	return refs.objs[*id]
}

func (id *ObjectId) Free() IF_fMEssCallBack {
	refs.Lock()
	defer refs.Unlock()

	obj := refs.objs[*id]
	delete(refs.objs, *id)
	*id = 0

	return obj
}
