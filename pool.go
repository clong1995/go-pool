package pool

import (
	"sync/atomic"
)

// Pool 资源池,池中的资源必须是不可变的,修改资源会产生并发问题
type Pool[T any] struct {
	resources []T
	total     int64
	index     atomic.Int64
}

// NewPool 创建一个资源池
func NewPool[T any](resources []T) *Pool[T] {
	pool := &Pool[T]{}
	for _, res := range resources {
		pool.resources = append(pool.resources, res)
	}
	pool.total = int64(len(resources))
	return pool
}

// Get 获取一个资源
func (p *Pool[T]) Get() *T {
	idx := p.index.Add(1) % p.total // 轮询选择下一个
	return &p.resources[idx]
}
