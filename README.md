# go-csv-read-benchmarks

Calling external requests 1ms 
        Junior       Senior 
10k  |   13.1s  |      0.343s
100k |   2m6s   |      0.483s
1M   |   21m    |      2.68s
-----------------------------

Calling external requests that take 100ms
        Junior       Senior 
10k  |  17m6s   |     2.51s
100k |  2h53m   |     21.2s
1M   |  28h58m  |     4m20s
-----------------------------


Calling external requests that take 250ms
        Junior       Senior 
10k  |  42m4s   |     5.39s
100k |  7h5m    |     1m6s
1M   |  70h42m  |     10m6s 
-----------------------------