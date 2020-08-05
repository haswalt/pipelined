module github.com/haswalt/pipelined

go 1.14

// replace pipelined.dev/audio => github.com/pipelined/audio v0.2.1
// replace pipelined.dev/wav => github.com/pipelined/wav v0.2.1
// replace pipelined.dev/signal => github.com/pipelined/signal v0.2.1

require (
	pipelined.dev/audio v0.2.2-0.20200804074927-aa4a1cccf461
	pipelined.dev/audio/mp3 v0.4.1-0.20200630191513-53012f1084c1 // indirect
	pipelined.dev/audio/wav v0.4.1-0.20200630072250-b574a19d81fc
	pipelined.dev/pipe v0.8.2
	pipelined.dev/repeat v0.0.0-20200620185127-98404aa74e08 // indirect
)
