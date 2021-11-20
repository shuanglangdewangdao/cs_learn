# 并发原语

## 闭包

## 共享变量必须持有锁（mutex）

## 利用锁实现原子操作

## 条件变量
***

cond:=sync.NewCond(&mu)
mu.Lock()
...
cond.Broadcast() //唤醒所有进程，让出锁
mu.Unlock()

mu.Lock()
...
cond.Wait() //让当前进程休眠，让出锁
...
mu.Unlock()

***
、

不建议使用 cond.Signal()
## Channel

类似queue但有较大差别，本质市线程之间发送接收数据，用于发送者接收者之间的同步

bufferChannel：初学者不建议使用会造成予以冲突

## 同步原语

WaitGroup()

1. Add()
2. Done()
3. Wait()



