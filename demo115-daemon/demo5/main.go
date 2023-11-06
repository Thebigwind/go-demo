package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"time"
)

//go程序调用自身转为后台运行 区分父进程子进程的问题 在子进程中再次启动子进程 守护进程的实现
/*
查标准库中有一个func (c *Cmd) Wait() error方法，可以阻塞等待子进程执行结束。
守护进程的逻辑就是启动一个子进程（处理业务逻辑，可称为业务进程），然后Wait()住。
若子进程退出了，则Wait()解除阻塞，再次重复一次之前的步骤。
如此循环，则相当于守护了一个业务进程常驻内存，保证服务的持续性。
*/

//示例:self1.go

//守护进程的实现, 基于之前的 background() 。可以替换示例self2.go中的main()函数进行测试
func main() {
	logFile := "/tmp/daemon.log"

	//启动一个子进程作为守护进程
	background(logFile, true) //启动子进程后退出

	//在守护进程中循环启动子进程
	for {
		cmd, err := background(logFile, false) //启动子进程后不自动退出
		if err != nil {
			log.Fatal("启动子进程失败:", err)
		}

		//根据返回值区分父进程子进程
		if cmd != nil { //父进程
			cmd.Wait() //等等子进程执行结束(监视子进程)
		} else { //子进程, 跳出让其执行后续业务代码
			break
		}
	}

	//以下是业务代码
	log.Println(os.Getpid(), "业务代码开始...")
	time.Sleep(time.Second * 20) //休眠20秒
	log.Println(os.Getpid(), "业务代码结束")

}

var runIdx int = 0               //background调用计数
const ENV_NAME = "XW_DAEMON_IDX" //环境变量名

func background(logFile string, isExit bool) (*exec.Cmd, error) {
	//判断子进程还是父进程
	runIdx++
	envIdx, err := strconv.Atoi(os.Getenv(ENV_NAME))
	if err != nil {
		envIdx = 0
	}
	if runIdx <= envIdx { //子进程, 退出
		return nil, nil
	}

	/*以下是父进程执行的代码*/

	//因为要设置更多的属性, 这里不使用`exec.Command`方法, 直接初始化`exec.Cmd`结构体
	cmd := &exec.Cmd{
		Path: os.Args[0],
		Args: os.Args,      //注意,此处是包含程序名的
		Env:  os.Environ(), //父进程中的所有环境变量
	}

	//为子进程设置特殊的环境变量标识
	cmd.Env = append(cmd.Env, fmt.Sprintf("%s=%d", ENV_NAME, runIdx))

	//若有日志文件, 则把子进程的输出导入到日志文件
	if logFile != "" {
		stdout, err := os.OpenFile(logFile, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
		if err != nil {
			log.Println(os.Getpid(), ": 打开日志文件错误:", err)
			return nil, err
		}
		cmd.Stderr = stdout
		cmd.Stdout = stdout
	}

	//异步启动子进程
	err = cmd.Start()
	if err != nil {
		log.Println(os.Getpid(), "启动子进程失败:", err)
		return nil, err
	} else {
		//执行成功
		log.Println(os.Getpid(), ":", "启动子进程成功:", "->", cmd.Process.Pid, "\n ")
	}

	//若启动子进程成功, 父进程是否直接退出
	if isExit {
		os.Exit(0)
	}

	return cmd, nil
}

/*
#编译
$ go build self2.go

#执行，启动的守护进程为 39541
$ ./self2 -a -b -c 123
2020/06/05 20:36:05 39540 : 启动子进程成功: -> 39541

#查看进程。可以看出，业务进程是39543，其父进程是39541
$ ps -ef |grep self2
  501 39541     1   0  8:36下午 ttys003    0:00.01 ./self2 -a -b -c 123
  501 39543 39541   0  8:36下午 ttys003    0:00.01 ./self2 -a -b -c 123

#查看日志。可以看到业务进程39543退出后，守护进程及时的启动了另一个业务进程39574
$ tail /tmp/daemon.log
2020/06/05 20:36:05 39541 : 启动子进程成功: -> 39543

2020/06/05 20:36:05 39543 业务代码开始...
2020/06/05 20:36:25 39543 业务代码结束
2020/06/05 20:36:25 39541 : 启动子进程成功: -> 39574

2020/06/05 20:36:25 39574 业务代码开始...
*/

/*
到此，守护进程的功能已经实现了。但作为一个库，对使用者还不太友好，我们需要封装一下。并且结合业务场景似乎还有一些细节问题需要考虑一下：

一个正常服务进程一般不会异常退出，可能并不需要无限的循环重启，这可以让使用者自定义最大重启次数
若业务进程连续不断的异常退出，是不应该继续不断重启了。可设置一个允许的最大连续异常退出次数
实际编写的服务程序，异常退出时不一定退出码就是非0。可以设置一个最短运行时间，协助判断是否是异常退出
最后封装为xdaemon库，开源在https://github.com/zh-five/xdaemon
*/
