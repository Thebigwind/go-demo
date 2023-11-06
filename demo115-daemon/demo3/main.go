package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"
)

//go程序调用自身转为后台运行 区分父进程子进程的问题
/*
为了区分子进程父进程，大多数开源库的解决方案是设置特殊的参数。这种方案是入侵式的，新设置的参数，有可能和go程序原有参数冲突。
虽然设置一些奇怪的参数名来降低冲突概率，但至少在使用过程中，并非完全保持参数原样启动子进程，可能会造成使用者的迷惑。这种方案不太完美，先舍弃。

前文提过github.com/sevlyar/go-daemon是巧妙的使用了环境变量，用来区分子进程和父进程。这种方案对go程序影响更小，产生冲突的可能性更小，也避免了使用者对参数变化的迷惑。
其原理是利用的是exec.Cmd的Env属性设置子进程的环境变量时，添加一个特殊的环境变量，用以标记子程序。用这个思路，我们把上面的例子修正一下。
模仿C语言里的fork，返回一个可用用于判断是子进程还是父进程的数据。
*/
//示例:self1.go

//注意此例子不建议使用go run直接运行，因为go run会先编译可执行文件到一个临时目录，然后再运行，其执行输出可能会有些让人迷惑。建议先编译为可执行文件后执行
func main() {
	cmd, err := background("/tmp/daemon.log")
	if err != nil {
		log.Fatal("启动子进程失败:", err)
	}

	//根据返回值区分父进程子进程
	if cmd != nil { //父进程
		log.Println("我是父进程:", os.Getpid(), "; 启动了子进程:", cmd.Process.Pid, "; 运行参数", os.Args)
		return //父进程退出
	} else { //子进程
		log.Println("我是子进程:", os.Getpid(), "; 运行参数:", os.Args)
	}

	//以下代码只有子进程会执行
	log.Println("只有子进程会运行:", os.Getpid(), "; 开始...")
	time.Sleep(time.Second * 20) //休眠20秒
	log.Println("只有子进程会运行:", os.Getpid(), "; 结束")
}

//@link https://www.zhihu.com/people/zh-five
func background(logFile string) (*exec.Cmd, error) {
	envName := "XW_DAEMON" //环境变量名称
	envValue := "SUB_PROC" //环境变量值

	val := os.Getenv(envName) //读取环境变量的值,若未设置则为空字符串
	if val == envValue {      //监测到特殊标识, 判断为子进程,不再执行后续代码
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
	cmd.Env = append(cmd.Env, fmt.Sprintf("%s=%s", envName, envValue))

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
	err := cmd.Start()
	if err != nil {
		return nil, err
	}

	return cmd, nil
}

/*
#编译
$ go build self1.go

#随便设置一些参数，查看执行效果
$ ./self1 -a -b
2020/06/05 19:05:44 我是父进程: 37886 ; 启动了子进程: 37887 ; 运行参数 [./self1 -a -b]

#查看子进程 37887
$ ps -ef |grep self1
  501 37887     1   0  7:05下午 ttys003    0:00.01 ./self1 -a -b

#查看子进程输出日志
$ tail /tmp/daemon.log
2020/06/05 19:05:44 我是子进程: 37887 ; 运行参数: [./self1 -a -b]
2020/06/05 19:05:44 只有子进程会运行: 37887 ; 开始...
2020/06/05 19:06:04 只有子进程会运行: 37887 ; 结束
*/
