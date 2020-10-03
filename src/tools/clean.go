package tools

import (
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"kube-tools/src/config"
)

func Clean() {

	_, err := config.KubeClientSet.CoreV1().Namespaces().Get(config.ShellNamespace, metav1.GetOptions{
		TypeMeta: metav1.TypeMeta{Kind: "Namespace", APIVersion: "v1"},
	})
	if err != nil {
		fmt.Println("清理命名空间", config.ShellNamespace, "成功！")
		return
	}
	err = config.KubeClientSet.AppsV1().DaemonSets(config.ShellNamespace).DeleteCollection(&metav1.DeleteOptions{}, metav1.ListOptions{})
	if err != nil {
		fmt.Println("清理命名空间", config.ShellNamespace, "失败！")
		panic(err.Error())
	}
	err = config.KubeClientSet.CoreV1().Namespaces().Delete(config.ShellNamespace, &metav1.DeleteOptions{})
	if err != nil {
		fmt.Println("清理命名空间", config.ShellNamespace, "失败！")
		panic(err.Error())
	}
	fmt.Println("清理命名空间", config.ShellNamespace, "成功！")
}