# Такси сервис

Результат работы теста Apache Benchmark


ab -c 100 -n 100000 http://localhost:8000/request
```
This is ApacheBench, Version 2.3 <$Revision: 1826891 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking localhost (be patient)
Completed 10000 requests
Completed 20000 requests
Completed 30000 requests
Completed 40000 requests
Completed 50000 requests
Completed 60000 requests
Completed 70000 requests
Completed 80000 requests
Completed 90000 requests
Completed 100000 requests
Finished 100000 requests


Server Software:
Server Hostname:        localhost
Server Port:            8000

Document Path:          /request
Document Length:        2 bytes

Concurrency Level:      100
Time taken for tests:   58.781 seconds
Complete requests:      100000
Failed requests:        0
Total transferred:      11800000 bytes
HTML transferred:       200000 bytes
Requests per second:    1701.22 [#/sec] (mean)
Time per request:       58.781 [ms] (mean)
Time per request:       0.588 [ms] (mean, across all concurrent requests)
Transfer rate:          196.04 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0   27  64.5     18     918
Processing:     1   32  69.5     21     922
Waiting:        0   21  54.2     15     895
Total:         10   58  93.8     39     936

Percentage of the requests served within a certain time (ms)
  50%     39
  66%     44
  75%     46
  80%     48
  90%     52
  95%     60
  98%    515
  99%    551
 100%    936 (longest request)

```

ab -c 200 -n 100000 http://localhost:8000/request
```
This is ApacheBench, Version 2.3 <$Revision: 1826891 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking localhost (be patient)
Completed 10000 requests
Completed 20000 requests
Completed 30000 requests
Completed 40000 requests
Completed 50000 requests
Completed 60000 requests
Completed 70000 requests
Completed 80000 requests
Completed 90000 requests
Completed 100000 requests
Finished 100000 requests


Server Software:
Server Hostname:        localhost
Server Port:            8000

Document Path:          /request
Document Length:        2 bytes

Concurrency Level:      200
Time taken for tests:   64.850 seconds
Complete requests:      100000
Failed requests:        0
Total transferred:      11800000 bytes
HTML transferred:       200000 bytes
Requests per second:    1542.03 [#/sec] (mean)
Time per request:       129.699 [ms] (mean)
Time per request:       0.648 [ms] (mean, across all concurrent requests)
Transfer rate:          177.70 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0   59 123.5     38    3585
Processing:     4   70 168.4     43    3587
Waiting:        0   45 106.4     29    3531
Total:         19  129 209.5     81    3837

Percentage of the requests served within a certain time (ms)
  50%     81
  66%     91
  75%     97
  80%    102
  90%    191
  95%    552
  98%    607
  99%    650
 100%   3837 (longest request)
```

ab -c 300 -n 100000 http://localhost:8000/request
```
This is ApacheBench, Version 2.3 <$Revision: 1826891 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking localhost (be patient)
Completed 10000 requests
Completed 20000 requests
Completed 30000 requests
Completed 40000 requests
Completed 50000 requests
Completed 60000 requests
Completed 70000 requests
Completed 80000 requests
Completed 90000 requests
Completed 100000 requests
Finished 100000 requests


Server Software:
Server Hostname:        localhost
Server Port:            8000

Document Path:          /request
Document Length:        2 bytes

Concurrency Level:      300
Time taken for tests:   61.529 seconds
Complete requests:      100000
Failed requests:        0
Total transferred:      11800000 bytes
HTML transferred:       200000 bytes
Requests per second:    1625.25 [#/sec] (mean)
Time per request:       184.587 [ms] (mean)
Time per request:       0.615 [ms] (mean, across all concurrent requests)
Transfer rate:          187.28 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0   87 123.5     57     754
Processing:     5   97 129.1     65     755
Waiting:        0   65 105.0     43     698
Total:         31  184 172.5    124     829

Percentage of the requests served within a certain time (ms)
  50%    124
  66%    135
  75%    144
  80%    152
  90%    593
  95%    648
  98%    709
  99%    747
 100%    829 (longest request)

```


ab -c 400 -n 500000 http://localhost:8000/request
```
This is ApacheBench, Version 2.3 <$Revision: 1826891 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking localhost (be patient)
Completed 50000 requests
Completed 100000 requests
Completed 150000 requests
Completed 200000 requests
Completed 250000 requests
Completed 300000 requests
Completed 350000 requests
Completed 400000 requests
Completed 450000 requests
Completed 500000 requests
Finished 500000 requests


Server Software:
Server Hostname:        localhost
Server Port:            8000

Document Path:          /request
Document Length:        2 bytes

Concurrency Level:      400
Time taken for tests:   314.586 seconds
Complete requests:      500000
Failed requests:        0
Total transferred:      59000000 bytes
HTML transferred:       1000000 bytes
Requests per second:    1589.39 [#/sec] (mean)
Time per request:       251.669 [ms] (mean)
Time per request:       0.629 [ms] (mean, across all concurrent requests)
Transfer rate:          183.15 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0  120 185.4     75    4255
Processing:     6  132 197.9     83    4267
Waiting:        1   87 164.0     54    4259
Total:         42  251 265.3    162    4656

Percentage of the requests served within a certain time (ms)
  50%    162
  66%    180
  75%    194
  80%    214
  90%    676
  95%    726
  98%    771
  99%    806
 100%   4656 (longest request)

```