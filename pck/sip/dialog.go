package pck

import (
	"sync"
)

// DialogState 定义会话状态
type DialogState string

const (
	DialogStateInit       DialogState = "INIT"
	DialogStateEarly      DialogState = "EARLY"
	DialogStateConfirmed  DialogState = "CONFIRMED"
	DialogStateTerminated DialogState = "TERMINATED"
)

// Dialog 表示一个 SIP 会话
type Dialog struct {
	CallID    string      // 会话的唯一标识
	LocalTag  string      // 本地 Tag
	RemoteTag string      // 远程 Tag
	State     DialogState // 会话状态
	RemoteURI string      // 远程 URI
	LocalURI  string      // 本地 URI
	RouteSet  []string    // 路由集合
	mu        sync.Mutex  // 互斥锁
}

// NewDialog 创建一个新的会话
func NewDialog(callID, localTag, remoteTag, localURI, remoteURI string) *Dialog {
	return &Dialog{
		CallID:    callID,
		LocalTag:  localTag,
		RemoteTag: remoteTag,
		LocalURI:  localURI,
		RemoteURI: remoteURI,
		State:     DialogStateInit,
		RouteSet:  make([]string, 0),
	}
}

// UpdateState 更新会话状态
func (d *Dialog) UpdateState(state DialogState) {
	d.mu.Lock()
	defer d.mu.Unlock()
	d.State = state
}

// AddRoute 添加路由到路由集合
func (d *Dialog) AddRoute(route string) {
	d.mu.Lock()
	defer d.mu.Unlock()
	d.RouteSet = append(d.RouteSet, route)
}

// Terminate 终止会话
func (d *Dialog) Terminate() {
	d.mu.Lock()
	defer d.mu.Unlock()
	d.State = DialogStateTerminated
}

// IsTerminated 检查会话是否已终止
func (d *Dialog) IsTerminated() bool {
	d.mu.Lock()
	defer d.mu.Unlock()
	return d.State == DialogStateTerminated
}
