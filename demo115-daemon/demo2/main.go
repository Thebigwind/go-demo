package main

import (
	"log"
	"os"
	"os/exec"
)

//go程序调用自身转为后台运行
//我们调用自身程序成功后，是希望子进程可以独自运行，然后父进程退出。这与上面调用外部程序的例子有几点不一样了：
//调用自身程序时，父进程不能以阻塞的方式进行了。因为若阻塞了，那就无法提前退出了
//父进程不能等待获取子进程的结果输出了，同样是为了提前退出

//非阻塞问题：查标准库，exec.Cmd是可以使用func (c *Cmd) Start() error非阻塞式运行外部程序的。若启动外部程序成功则返回nil，否则返回错误信息。
//子进程的结果输出问题：查看本文之前引用的标准库文档exec.Cmd的两个属性Stdout和Stderr（标准输出和错误输出）都是io.Writer接口。 那么我们就可以把标准输出和错误输出定向到日志文件中。
//                  当然若不需要，也可以不用设置Stdout和Stderr两个属性，系统将抛弃子进程标准输出和错误输出的信息。

// !!! 切勿运行此程序 !!!
//示例：self.go

//以上代码粗看起来好像没有问题，但仔细一想，会存在一个问题：启动的子进程也会执行background()方法，再次启动一个子进程。如此循环，会不断的创建子进程。
//也就是说以上例子里，代码无法判断自身是父进程还是子进程。

func main() {
	//background("/tmp/daemon.log")
}

func background(logFile string) error {
	//os.Args 是一个切片,保管了命令行参数，第一个是程序名
	//go程序启动时不包含管道符了,就直接运行了
	cmd := exec.Command(os.Args[0], os.Args[1:]...)

	//若有日志文件, 则把子进程的输出导入到日志文件
	if logFile != "" {
		stdout, err := os.OpenFile(logFile, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
		if err != nil {
			log.Println(os.Getpid(), ": 打开日志文件错误:", err)
			return err
		}
		cmd.Stderr = stdout
		cmd.Stdout = stdout
	}

	//异步启动子进程
	err := cmd.Start()
	if err != nil {
		return err
	}

	return nil
}
