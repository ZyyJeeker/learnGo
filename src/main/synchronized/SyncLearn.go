package synchronized

import "sync"

func syncTest() {
	// 互斥所
	lock := new(sync.Mutex)
	lock.TryLock()
	// do  something
	lock.Unlock()

	// 读写锁
	rLock := new(sync.RWMutex)
	rLock.TryRLock()
	// read or write
	rLock.RUnlock()

	// 无论多少个线程，确保只执行一次
	once := new(sync.Once)
	once.Do(func() {
		// do something
	})
}
