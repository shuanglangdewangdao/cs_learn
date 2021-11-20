# raft

## 本身具有复制系统

脑裂问题时指，程序正常运行但本来互相复制的两台机器无法同步

1. 大量资金创建无瑕的网络
2. 人类介入，随时查看是否宕机

## 现有方法

少数服从多数的投票方法（包括宕机的服务器）

Paxos，VSR

### log

1. 复制状态机要按相同的顺序执行命令
2. leader保证副本都有相同的日志
3. 宕机后重新载入按日志执行

#### 重载

#### 全宕机

1. 重新选举leader
2. 通信确定执行过程

#### 一致

一段时间后会强制同步

#### 最长log的服务器成为leader？

不可以

#### 选举限制（election restriction）

1. 有最后的term或者最后相同的term
2. 且>=log长度

### raft的接口

1. start(command)：客户端命令放入log，返回（index,term)
2. applechannle()：是否提交

## Raft架构的代码

C->LS(applacation)->LS(raft)->RS(raft)(投票结束后)->LS(raft)LS(applacation)->C

是否接收在吓一跳消息中通知

需要其他指令表明备机的执行经度

## leader

一个任期（term）只有=一个leader

### 选举定时器

每个选举周期发起一次选举,发现没有leader

term++->ReuestVote、

只有leader允许发起AppendEntries

#### 瓜分选票

随机化选举定时器

至少定时器间隔时心跳时间间隔的2倍，最长时间考虑恢复时间和故障频率

每个服务器的超时间隔要大于选举完成时间

## 备份方案

### 快速备份（Fast Bakup）

#### ex1

S1:4 5 5

S2:4 6 6 6(leader)

### ex2

S1:4 4 4

S2:4 6 6 6 

### ex3

S1:4 4 4

S2:4 6 6 6 



## 实现中的常见bug

1. 死锁：多重持锁

2. 程序停止时不完全停止，导致执行错误

## debug

### 死锁
util.go:DPrintf()

Ctrl+\

### 共享变量

-race

### 并行测试

