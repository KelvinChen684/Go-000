package rolling

import (
	"container/ring"
	"sync"
	"time"
)

type Bucket struct {
	Success int64
	Failure int64
	Timeout int64
	Rejection int64
}

// 滑动窗口计数器
type RollingCount struct {
	sync.RWMutex
	now int64			// 当前窗口最近的桶的秒数，每个桶1s，该数据作为索引查找桶
	buckets *ring.Ring	// 10个桶，使用环形链表
	counts *Bucket		// 窗口内4 个指标数据总计：成功量、失败量、超时量、拒绝量
}

func NewRollingCount() *RollingCount {
	r := &RollingCount{
		now:     time.Now().Unix(),
		buckets: ring.New(10),	// 10 个桶
		counts:  &Bucket{},
	}
	for i := 0; i < r.buckets.Len(); i++ {
		r.buckets.Value= &Bucket{}
		r.buckets = r.buckets.Next()
	}

	return  r
}

func (r *RollingCount) decreBucket(b *Bucket)  {
	r.counts.Success -= b.Success
	r.counts.Failure -= b.Failure
	r.counts.Timeout -= b.Timeout
	r.counts.Rejection -= b.Rejection
}

func (r *RollingCount) updateTime()  {
	now := time.Now().Unix()
	r.Lock()
	defer r.Unlock()

	if now - r.now >= int64(10) {	// 当前时间窗口已失效，清除当前窗口内数据
		for i := 0; i < r.buckets.Len(); i++ {
			r.buckets.Value= &Bucket{}
			r.buckets = r.buckets.Next()
		}
		r.counts = &Bucket{}
		r.now = now
		return
	} else if now - r.now > int64(0) {	// 清除已失效桶内的数据
		for i := int64(0); i < now - r.now; i++ {
			r.buckets = r.buckets.Prev()
			b := r.buckets.Value.(*Bucket)
			r.decreBucket(b)
			r.buckets.Value = &Bucket{}
		}
		r.now = now
		return
	}
	// 还在当前桶内
	return
}
func (r *RollingCount) GetCounts() Bucket {
	r.updateTime()

	r.RLock()
	defer r.RUnlock()

	return *r.counts	// 不能返回指针
}

func (r *RollingCount) IncreSuccess()  {
	r.updateTime()

	r.Lock()
	defer r.Unlock()

	r.counts.Success++
	r.buckets.Value.(*Bucket).Success++
}

func (r *RollingCount) IncreFailure()  {
	r.updateTime()

	r.Lock()
	defer r.Unlock()

	r.counts.Failure++
	r.buckets.Value.(*Bucket).Failure++
}

func (r *RollingCount) IncreTimeout()  {
	r.updateTime()

	r.Lock()
	defer r.Unlock()

	r.counts.Timeout++
	r.buckets.Value.(*Bucket).Timeout++
}

func (r *RollingCount) IncreRejection()  {
	r.updateTime()

	r.Lock()
	defer r.Unlock()

	r.counts.Rejection++
	r.buckets.Value.(*Bucket).Rejection++
}
