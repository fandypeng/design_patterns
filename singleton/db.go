package singleton

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"sync"
	"sync/atomic"
)

var singletondb *sql.DB

// lazy loading, not thread safe
func GetDbInstanceNTS() *sql.DB {
	mu.Lock()
	defer mu.Unlock()
	if singletondb == nil {
		db, _ := sql.Open("mysql", GetGameConf().GetDsn())
		singletondb = db
	}
	return singletondb
}

//硬件支持的原子标记加锁的单例模式，thread safe
var initialized uint32 = 0
var mu sync.Mutex

func GetDbInstanceWithLock() *sql.DB {
	if atomic.LoadUint32(&initialized) == 1 {
		return singletondb
	}
	mu.Lock()
	defer mu.Unlock()
	if initialized == 0 {
		db, _ := sql.Open("mysql", GetGameConf().GetDsn())
		singletondb = db
		atomic.StoreUint32(&initialized, 1)
	}
	return singletondb
}

var once sync.Once

//简化单例模式 thread safe
func GetDbInstance() *sql.DB {
	once.Do(func() {
		db, _ := sql.Open("mysql", GetGameConf().GetDsn())
		singletondb = db
	})
	return singletondb
}
