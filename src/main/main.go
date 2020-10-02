package main

import (
	"flag"
	"fmt"
	"kube-tools/src/config"
	"kube-tools/src/tools"
)

func main() {

	config.InitKube()
	action := flag.Arg(0)
	switch action {
	case "node":
		// 返回节点信息列表
		tools.Node()
	case "install":
		// 部署用于执行 node shell 的 ds
		fmt.Println("飞速开发中...")
	case "uninstall":
		// 清理 node shell 的 ds
		fmt.Println("飞速开发中...")
	case "cat":
		// 获取node内的文件内容
		tools.Cat()
	default:
		flag.Usage()
	}

}
