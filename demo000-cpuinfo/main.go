package main

import (
	"fmt"
	"runtime"
)

func main() {
	// 获取操作系统
	os := runtime.GOOS
	fmt.Println("Operating System:", os)

	// 获取架构
	arch := runtime.GOARCH
	fmt.Println("Architecture:", arch)
}

/*
在 Shell 中，您可以使用 uname 命令来查看操作系统的信息，包括架构。以下是一些常用的 uname 命令：

查看操作系统名称：

bash
Copy code
uname -s
查看操作系统版本：

bash
Copy code
uname -v
查看操作系统和版本：

bash
Copy code
uname -a
查看架构（硬件平台）：

bash
Copy code
uname -m
查看处理器类型：

bash
Copy code
uname -p
例如，要查看操作系统和架构，您可以运行以下命令：

bash
Copy code
uname -a
这将显示包括操作系统名称、版本和架构在内的详细信息。

如果您只关心架构信息，可以运行：

bash
Copy code
uname -m
这将显示当前系统的硬件架构。在 ARM 架构系统上，可能会显示 arm。在 x86 架构系统上，可能会显示 x86_64 或 i686 等。
*/
