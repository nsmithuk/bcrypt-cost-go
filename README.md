# bcrypt-cost-go

Simple test to measure the speed of bcrypt in Go for each of the supported costs.

Install
```bash
go get -u github.com/NSmithUK/bcrypt-cost-go
```

Run
```bash
$GOPATH/bin/bcrypt-cost-go
```

## Example output
Run on the cheapest 'High CPU' Digital Ocean virtual machine.

| Cost  | Time to hash password |
| ------------- | ------------- |
| 4  | 3.193709ms  |
| 5  | 6.03618ms  |
| 6  | 11.258473ms  |
| 7  | 20.273336ms  |
| 8  | 34.605597ms  |
| 9  | 56.568637ms  |
| 10 | 86.246206ms  |
| 11 | 144.570113ms  |
| 12 | 289.117209ms  |
| 13 | 577.527173ms  |
| 14 | 1.155819422s  |
| 15 | 2.311880633s  |
| 16 | 4.625168556s  |
| 17 | 9.251611191s  |
| 18 | 18.4832305s  |
| 19 | 37.10841037s  |
| 20 | 1m 14.077770059s  |
| 21 | 2m 28.111918885s  |
| 22 | 4m 56.326662123s  |
| 23 | 9m 51.933949103s  |
| 24 | 19m 45.156559968s  |
| 25 | 39m 29.474770423s  |
| 26 | 1h 20m2 4.538693859s  |
| 27 | 2h 39m 41.19676742s  |
| 28 | 5h 9m 57.203482634s  |
| 29 | 10h 32m 4.413076565s  |
| 30 | 19h 59m 30.093452502s  |
| 31 | 40h 55m 38.340977971s  |
