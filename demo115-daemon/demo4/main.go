package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"time"
)

//go程序调用自身转为后台运行 区分父进程子进程的问题 在子进程中再次启动子进程
/*
我们若在子进程中调用background方法，会发现是无法启动新的子进程的。原因是不管第几次调用background方法，环境变量的判断结果都是一样。有两个因素并没有考虑到：

第几次调用background
当前子进程是第几代子进程
结合这两个因素，我们似乎可以设计出一个判断策略，background知道什么时候我该启动子进程，什么时候不该启动。我们设计一个变量ruuIdx记录调用background的次数，
启动子进程时把此计数写入到子进程的环境变量中，用于标记此进程是第几代子进程（envIdx）。
显然，在子进程中，若runIdx等于envIdx时，那父进程正是调用了此次的background而启动了这个子进程。推导判断一下其它情况，可制定完成的策略如下：

runIdx = envIdx时：代表意义如上所述，不启动子进程
runIdx < envIdx时：表示是启动前几代子进程的调用，不启动子进程
runIdx > envIdx时：表示需要启动新启动一个子进程
*/
//示例:self1.go

//注意此例子不建议使用go run直接运行，因为go run会先编译可执行文件到一个临时目录，然后再运行，其执行输出可能会有些让人迷惑。建议先编译为可执行文件后执行
//示例:self2.go

func main() {
	logFile := "/tmp/daemon.log"
	background(logFile, true) //启动子进程后退出
	background(logFile, true) //启动子进程后退出
	background(logFile, true) //启动子进程后退出

	//以下代码只有最后一代子进程会执行
	log.Println(os.Getpid(), "业务代码开始...")
	time.Sleep(time.Second * 20) //休眠20秒
	log.Println(os.Getpid(), "业务代码结束")
}

var runIdx int = 0               //background调用计数
const ENV_NAME = "XW_DAEMON_IDX" //环境变量名

//@link https://www.zhihu.com/people/zh-five
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

#随便设置一些参数执行
$ ./self2 -a -b -c 123
2020/06/05 19:58:27 38984 : 启动子进程成功: -> 38985

#查看进程，看到的是最终子进程
$ ps -ef |grep self2
  501 38990     1   0  7:58下午 ttys003    0:00.01 ./self2 -a -b -c 123

#查看日志
$ tail /tmp/daemon.log
2020/06/05 19:58:27 38985 : 启动子进程成功: -> 38988

2020/06/05 19:58:27 38988 : 启动子进程成功: -> 38990

2020/06/05 19:58:28 38990 业务代码开始...
2020/06/05 19:58:48 38990 业务代码结束
*/

/*
由日志可以看出，成功的启动了3代子进程：38984（父进程）-> 38985 -> 38988 -> 38990。最终的38990子进程执行了业务代码。

注意：此种策略判断的前提条件是，逐代启动子进程。若某进程里重复启动了多个子进程，那么其子进程若想再启动子进程，可能会失败。如以下例子
*/

//非逐代启动子进程的异常情况
func main1() {
	logFile := "/tmp/daemon.log"
	cmd, err := background(logFile, false) //启动子进程后不自动退出
	if err != nil {
		log.Fatal("启动子进程失败:", err)
	}

	//根据返回值区分父进程子进程
	if cmd != nil { //父进程
		//父进程再次启动一个子进程, 非逐代启动了
		background(logFile, true) //启动子进程后退出
		return                    //父进程退出
	}

	//父进程里第2次启动的子进程, 此处调用出现异常情况: 将不会启动子进程,而会直接略过执行后面的代码
	background(logFile, true) //启动子进程后退出

	//以下代码只有最后一代子进程会执行
	log.Println(os.Getpid(), "业务代码开始...")
	time.Sleep(time.Second * 20) //休眠20秒
	log.Println(os.Getpid(), "业务代码结束")
}

//执行结果为
/*
#编译
$ go build self2.go

#执行。启动了两个子进程，注意第2此启动39291进程将有异常
$ ./self2 -a -b -c 123
2020/06/05 20:16:58 39289 : 启动子进程成功: -> 39290

2020/06/05 20:16:58 39289 : 启动子进程成功: -> 39291

#查看进程
$ ps -ef |grep self2
  501 39291     1   0  8:16下午 ttys003    0:00.01 ./self2 -a -b -c 123
  501 39292     1   0  8:16下午 ttys003    0:00.01 ./self2 -a -b -c 123

#查看日志。主要只有39290再次启动了子进程，而39291则直接执行了业务代码
$ tail /tmp/daemon.log
2020/06/05 20:16:58 39290 : 启动子进程成功: -> 39292

2020/06/05 20:16:58 39291 业务代码开始...
2020/06/05 20:16:58 39292 业务代码开始...
2020/06/05 20:17:18 39291 业务代码结束
2020/06/05 20:17:18 39292 业务代码结束
*/
//若是重复启动的子进程不再启动子进程，则无影响。后续守护进程的实现，会有这种情况。
