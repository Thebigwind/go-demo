查看信息图片化方法

获取.pb.gz文件
me@localhostt ~ % go tool pprof http://localhost:6060/debug/pprof/profile
Fetching profile over HTTP from http://localhost:6060/debug/pprof/profile
Saved profile in /Users/me/pprof/pprof.samples.cpu.003.pb.gz
Type: cpu
Time: Jun 18, 2022 at 10:08pm (CST)
Duration: 30s, Total samples = 170ms ( 0.57%)
Entering interactive mode (type "help" for commands, "o" for options)
(pprof)
(pprof) go tool pprof -http=:8080 /Users/me/pprof/pprof.samples.cpu.003.pb.g
unrecognized command: "go"
(pprof) go tool pprof -http=:8080 /Users/me/pprof/pprof.samples.cpu.003.pb.gz
unrecognized command: "go"
(pprof) %

浏览器上查看图片化信息
me@localhostt ~ % go tool pprof -http=:8080 /Users/me/pprof/pprof.samples.cpu.003.pb.gz
Serving web UI on http://localhost:8080