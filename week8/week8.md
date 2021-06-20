#### 第八周：分布式缓存 & 分布式事务

1、使用 redis benchmark 工具, 测试 10 20 50 100 200 1k 5k 字节 value 大小，redis get set 性能。

2、写入一定量的 kv 数据, 根据数据大小 1w-50w 自己评估, 结合写入前后的 info memory 信息  , 分析上述不同 value 大小下，平均每个 key 的占用内存空间。

---
1、

10 字节
```
...> redis-benchmark -t set,get -d 10 -q
====== SET ======
  100000 requests completed in 2.70 seconds
  50 parallel clients
  10 bytes payload
  keep alive: 1

20.43% <= 1 milliseconds
98.99% <= 2 milliseconds
99.85% <= 3 milliseconds
99.93% <= 4 milliseconds
99.99% <= 5 milliseconds
99.99% <= 6 milliseconds
100.00% <= 55 milliseconds
100.00% <= 55 milliseconds
37050.76 requests per second

====== GET ======
  100000 requests completed in 2.67 seconds
  50 parallel clients
  10 bytes payload
  keep alive: 1

25.33% <= 1 milliseconds
98.71% <= 2 milliseconds
99.62% <= 3 milliseconds
99.78% <= 4 milliseconds
99.86% <= 5 milliseconds
99.88% <= 6 milliseconds
99.90% <= 7 milliseconds
99.90% <= 12 milliseconds
99.90% <= 15 milliseconds
99.90% <= 17 milliseconds
99.90% <= 20 milliseconds
99.95% <= 36 milliseconds
99.95% <= 37 milliseconds
99.96% <= 38 milliseconds
99.96% <= 39 milliseconds
99.96% <= 40 milliseconds
99.97% <= 41 milliseconds
99.97% <= 42 milliseconds
99.97% <= 43 milliseconds
99.97% <= 44 milliseconds
99.98% <= 45 milliseconds
99.98% <= 46 milliseconds
99.98% <= 47 milliseconds
99.98% <= 48 milliseconds
99.98% <= 49 milliseconds
99.99% <= 50 milliseconds
99.99% <= 51 milliseconds
99.99% <= 52 milliseconds
99.99% <= 53 milliseconds
99.99% <= 54 milliseconds
100.00% <= 55 milliseconds
100.00% <= 56 milliseconds
100.00% <= 57 milliseconds
100.00% <= 57 milliseconds
37495.31 requests per second
```

20 字节
```
...> redis-benchmark -t set,get -d 20
====== SET ======
  100000 requests completed in 3.09 seconds
  50 parallel clients
  20 bytes payload
  keep alive: 1

9.63% <= 1 milliseconds
93.17% <= 2 milliseconds
97.45% <= 3 milliseconds
98.52% <= 4 milliseconds
98.99% <= 5 milliseconds
99.27% <= 6 milliseconds
99.51% <= 7 milliseconds
99.62% <= 8 milliseconds
99.71% <= 9 milliseconds
99.79% <= 10 milliseconds
99.82% <= 11 milliseconds
99.82% <= 12 milliseconds
99.83% <= 13 milliseconds
99.83% <= 14 milliseconds
99.84% <= 15 milliseconds
99.85% <= 16 milliseconds
99.86% <= 17 milliseconds
99.89% <= 18 milliseconds
99.93% <= 19 milliseconds
99.97% <= 20 milliseconds
99.97% <= 21 milliseconds
99.98% <= 23 milliseconds
99.99% <= 24 milliseconds
99.99% <= 25 milliseconds
100.00% <= 25 milliseconds
32414.91 requests per second

====== GET ======
  100000 requests completed in 3.09 seconds
  50 parallel clients
  20 bytes payload
  keep alive: 1

9.89% <= 1 milliseconds
91.97% <= 2 milliseconds
97.83% <= 3 milliseconds
99.09% <= 4 milliseconds
99.53% <= 5 milliseconds
99.72% <= 6 milliseconds
99.79% <= 7 milliseconds
99.82% <= 8 milliseconds
99.86% <= 9 milliseconds
99.89% <= 10 milliseconds
99.90% <= 11 milliseconds
99.93% <= 12 milliseconds
99.93% <= 13 milliseconds
99.94% <= 14 milliseconds
99.96% <= 15 milliseconds
99.99% <= 16 milliseconds
99.99% <= 17 milliseconds
99.99% <= 18 milliseconds
100.00% <= 19 milliseconds
100.00% <= 20 milliseconds
32372.94 requests per second
```

50 字节
```
...> redis-benchmark -t set,get -d 50
====== SET ======
  100000 requests completed in 2.68 seconds
  50 parallel clients
  50 bytes payload
  keep alive: 1

12.73% <= 1 milliseconds
98.32% <= 2 milliseconds
99.84% <= 3 milliseconds
99.94% <= 4 milliseconds
99.98% <= 5 milliseconds
100.00% <= 5 milliseconds
37327.36 requests per second

====== GET ======
  100000 requests completed in 2.68 seconds
  50 parallel clients
  50 bytes payload
  keep alive: 1

21.28% <= 1 milliseconds
98.12% <= 2 milliseconds
99.33% <= 3 milliseconds
99.55% <= 4 milliseconds
99.70% <= 5 milliseconds
99.82% <= 6 milliseconds
99.90% <= 7 milliseconds
99.93% <= 8 milliseconds
99.95% <= 11 milliseconds
99.96% <= 12 milliseconds
99.97% <= 13 milliseconds
100.00% <= 14 milliseconds
100.00% <= 14 milliseconds
37327.36 requests per second
```
100 字节
```
...> redis-benchmark -t set,get -d 100
====== SET ======
  100000 requests completed in 2.89 seconds
  50 parallel clients
  100 bytes payload
  keep alive: 1

6.86% <= 1 milliseconds
96.87% <= 2 milliseconds
99.50% <= 3 milliseconds
99.73% <= 4 milliseconds
99.85% <= 5 milliseconds
99.89% <= 6 milliseconds
99.90% <= 12 milliseconds
99.91% <= 13 milliseconds
99.92% <= 14 milliseconds
99.93% <= 15 milliseconds
99.93% <= 16 milliseconds
99.95% <= 17 milliseconds
99.98% <= 18 milliseconds
99.99% <= 19 milliseconds
100.00% <= 19 milliseconds
34614.05 requests per second

====== GET ======
  100000 requests completed in 2.51 seconds
  50 parallel clients
  100 bytes payload
  keep alive: 1

17.47% <= 1 milliseconds
99.44% <= 2 milliseconds
99.91% <= 3 milliseconds
99.93% <= 4 milliseconds
99.96% <= 5 milliseconds
100.00% <= 5 milliseconds
39808.91 requests per second
```
200 字节
```
...> redis-benchmark -t set,get -d 200
====== SET ======
  100000 requests completed in 2.64 seconds
  50 parallel clients
  200 bytes payload
  keep alive: 1

13.51% <= 1 milliseconds
98.50% <= 2 milliseconds
99.96% <= 3 milliseconds
99.98% <= 4 milliseconds
100.00% <= 4 milliseconds
37821.48 requests per second

====== GET ======
  100000 requests completed in 2.64 seconds
  50 parallel clients
  200 bytes payload
  keep alive: 1

20.17% <= 1 milliseconds
99.20% <= 2 milliseconds
99.74% <= 3 milliseconds
99.82% <= 4 milliseconds
99.85% <= 5 milliseconds
99.86% <= 6 milliseconds
99.87% <= 7 milliseconds
99.88% <= 8 milliseconds
99.89% <= 9 milliseconds
99.90% <= 15 milliseconds
99.90% <= 16 milliseconds
99.90% <= 18 milliseconds
99.93% <= 19 milliseconds
99.95% <= 55 milliseconds
99.96% <= 56 milliseconds
99.96% <= 57 milliseconds
99.96% <= 58 milliseconds
99.97% <= 59 milliseconds
99.97% <= 60 milliseconds
99.98% <= 61 milliseconds
99.98% <= 62 milliseconds
99.98% <= 63 milliseconds
99.98% <= 64 milliseconds
99.99% <= 65 milliseconds
99.99% <= 66 milliseconds
99.99% <= 67 milliseconds
100.00% <= 68 milliseconds
100.00% <= 69 milliseconds
100.00% <= 69 milliseconds
37821.48 requests per second
```
1k 
```
...> redis-benchmark -t set,get -d 1024
====== SET ======
  100000 requests completed in 2.67 seconds
  50 parallel clients
  1024 bytes payload
  keep alive: 1

8.09% <= 1 milliseconds
98.99% <= 2 milliseconds
99.83% <= 3 milliseconds
99.86% <= 4 milliseconds
99.87% <= 5 milliseconds
99.90% <= 6 milliseconds
99.93% <= 7 milliseconds
99.95% <= 23 milliseconds
99.96% <= 24 milliseconds
99.98% <= 25 milliseconds
99.99% <= 26 milliseconds
100.00% <= 27 milliseconds
100.00% <= 28 milliseconds
100.00% <= 29 milliseconds
100.00% <= 30 milliseconds
37425.15 requests per second

====== GET ======
  100000 requests completed in 2.72 seconds
  50 parallel clients
  1024 bytes payload
  keep alive: 1

13.05% <= 1 milliseconds
97.95% <= 2 milliseconds
99.24% <= 3 milliseconds
99.54% <= 4 milliseconds
99.71% <= 5 milliseconds
99.80% <= 6 milliseconds
99.85% <= 7 milliseconds
99.88% <= 8 milliseconds
99.93% <= 9 milliseconds
99.95% <= 12 milliseconds
99.96% <= 13 milliseconds
99.97% <= 14 milliseconds
100.00% <= 15 milliseconds
100.00% <= 15 milliseconds
36764.71 requests per second
```

5k 
```
...> redis-benchmark -t set,get -d 5120
====== SET ======
  100000 requests completed in 2.99 seconds
  50 parallel clients
  5120 bytes payload
  keep alive: 1

2.34% <= 1 milliseconds
96.16% <= 2 milliseconds
99.44% <= 3 milliseconds
99.68% <= 4 milliseconds
99.82% <= 5 milliseconds
99.90% <= 6 milliseconds
99.94% <= 7 milliseconds
99.94% <= 8 milliseconds
99.95% <= 14 milliseconds
99.96% <= 15 milliseconds
99.97% <= 16 milliseconds
99.98% <= 17 milliseconds
100.00% <= 18 milliseconds
100.00% <= 18 milliseconds
33467.20 requests per second

====== GET ======
  100000 requests completed in 3.05 seconds
  50 parallel clients
  5120 bytes payload
  keep alive: 1

24.81% <= 1 milliseconds
98.27% <= 2 milliseconds
99.82% <= 3 milliseconds
99.88% <= 4 milliseconds
99.92% <= 5 milliseconds
99.95% <= 11 milliseconds
99.96% <= 12 milliseconds
99.97% <= 13 milliseconds
99.99% <= 14 milliseconds
100.00% <= 14 milliseconds
32765.40 requests per second
```

---
2、

数据大小：5w

value大小：10 字节

写入前：
used_memory:724856

写入后：
used_memory:3514728

平均每个key占用内存空间：((3514728 - 724856) - 50000 * 10 ) / 50000 = 45.797 字节

---

数据大小：5w

value大小：20 字节

写入前：
used_memory:725832

写入后：
used_memory:3764632

平均每个key占用内存空间：((3764632 - 725832) - 50000 * 20 ) / 50000 = 40.776 字节

---

数据大小：5w

value大小：50 字节

写入前：
used_memory:726864

写入后：
used_memory:5057480

平均每个key占用内存空间：((5057480 - 726864) - 50000 * 50 ) / 50000 = 36.612 字节

---
### // TODO 数据错误，待检查
---

数据大小：5w

value大小：100 字节

写入前：
used_memory:727536

写入后：
used_memory:6541536

平均每个key占用内存空间：((6541536 - 727536) - 50000 * 100 ) / 50000 =  字节

---

数据大小：5w

value大小：200 字节

写入前：
used_memory:728328

写入后：
used_memory:10362056

平均每个key占用内存空间：((10362056 - 728328) - 50000 * 200 ) / 50000 =  字节

---

数据大小：5w

value大小：1k 

写入前：
used_memory:728664

写入后：
used_memory:43430008

平均每个key占用内存空间：((43430008 - 728664) - 50000 * 1024 ) / 50000 =  字节

---

数据大小：5w

value大小：2k 

写入前：
used_memory:730048

写入后：
used_memory:84032352

平均每个key占用内存空间：((84032352 - 730048) - 50000 * 2048 ) / 50000 =  字节

命令：
```
cli> redis-benchmark -t set -r 50000 -n 50000 -d 10
```



redis-benchmark 使用说明
```
Usage: redis-benchmark [-h <host>] [-p <port>] [-c <clients>] [-n <requests>] [-k <boolean>]

 -h <hostname>      Server hostname (default 127.0.0.1)
 -p <port>          Server port (default 6379)
 -s <socket>        Server socket (overrides host and port)
 -a <password>      Password for Redis Auth
 -c <clients>       Number of parallel connections (default 50)
 -n <requests>      Total number of requests (default 100000)
 -d <size>          Data size of SET/GET value in bytes (default 3)
 --dbnum <db>       SELECT the specified db number (default 0)
 -k <boolean>       1=keep alive 0=reconnect (default 1)
 -r <keyspacelen>   Use random keys for SET/GET/INCR, random values for SADD
  Using this option the benchmark will expand the string __rand_int__
  inside an argument with a 12 digits number in the specified range
  from 0 to keyspacelen-1. The substitution changes every time a command
  is executed. Default tests use this to hit random keys in the
  specified range.
 -P <numreq>        Pipeline <numreq> requests. Default 1 (no pipeline).
 -e                 If server replies with errors, show them on stdout.
                    (no more than 1 error per second is displayed)
 -q                 Quiet. Just show query/sec values
 --csv              Output in CSV format
 -l                 Loop. Run the tests forever
 -t <tests>         Only run the comma separated list of tests. The test
                    names are the same as the ones produced as output.
 -I                 Idle mode. Just open N idle connections and wait.
```

