package pck

import (
	"sync"
)

// 事务状态
type TransactionState byte

const (
	StateTrying TransactionState = iota
	StateProceeding
	StateCompleted
	StateTerminated
)

// SIP 事务
type Transaction struct {
	ID       string
	State    TransactionState
	Request  *Message
	Response *Message
	mu       sync.Mutex
}

// 创建新事务
func NewTransaction(id string, req *Message) *Transaction {
	return &Transaction{
		ID:      id,
		State:   StateTrying,
		Request: req,
	}
}

// 更新事务状态
func (t *Transaction) UpdateState(state TransactionState) {
	t.mu.Lock()
	defer t.mu.Unlock()
	t.State = state
}

// 处理响应
func (t *Transaction) HandleResponse(res *Message) {
	t.mu.Lock()
	defer t.mu.Unlock()
	t.Response = res

	switch t.State {
	case StateTrying, StateProceeding:
		if res.Headers["CSeq"] == t.Request.Headers["CSeq"] {
			t.State = StateCompleted
		}
	case StateCompleted:
		t.State = StateTerminated
	}
}

// 超时处理
func (t *Transaction) Timeout() {
	t.mu.Lock()
	defer t.mu.Unlock()
	if t.State != StateTerminated {
		t.State = StateTerminated
	}
}
