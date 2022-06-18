交互式命令行



demo:
me@localhostt pprof % go tool pprof http://localhost:6060/debug/pprof/profile
Fetching profile over HTTP from http://localhost:6060/debug/pprof/profile
Saved profile in /Users/me/pprof/pprof.samples.cpu.002.pb.gz
Type: cpu
Time: Jun 18, 2022 at 9:52pm (CST)
Duration: 30s, Total samples = 160ms ( 0.53%)
Entering interactive mode (type "help" for commands, "o" for options)
(pprof) top10
Showing nodes accounting for 160ms, 100% of 160ms total
Showing top 10 nodes out of 24
flat  flat%   sum%        cum   cum%
50ms 31.25% 31.25%       80ms 50.00%  runtime.kevent
40ms 25.00% 56.25%       50ms 31.25%  syscall.syscall
30ms 18.75% 75.00%       30ms 18.75%  runtime.libcCall
30ms 18.75% 93.75%       30ms 18.75%  runtime.pthread_cond_wait
10ms  6.25%   100%       10ms  6.25%  runtime.reentersyscall
0     0%   100%       50ms 31.25%  internal/poll.(*FD).Write
0     0%   100%       50ms 31.25%  internal/poll.ignoringEINTRIO (inline)
0     0%   100%       50ms 31.25%  log.(*Logger).Output
0     0%   100%       50ms 31.25%  log.Printf
0     0%   100%       50ms 31.25%  main.main.func1
(pprof) web
(pprof)
此时会生成一个 .svg 格式的图片，在web浏览器中打开。


