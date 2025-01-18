package jobs

import (
	"os/exec"
	"time"

	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/klog/v2"
)

const (
	RepoName = "Axpz/xMinima"
)

type XMinimaCICDJob struct {
}

func (j *XMinimaCICDJob) Do() error {
	return wait.ExponentialBackoff(
		wait.Backoff{ // 配置指数回退的参数
			Duration: 1 * time.Second, // 初始等待时间
			Factor:   2,               // 回退因子，表示每次重试间隔时间倍增
			Jitter:   0.1,             // 允许的随机抖动，以避免集群中多个组件同时重试
			Steps:    3,               // 最大重试次数
		},
		j.do,
	)
}

func (j *XMinimaCICDJob) do() (bool, error) {
	klog.Infof("running CICD job: %s", RepoName)
	cmd := exec.Command("/bin/bash", "-c", "printenv")
	out, err := cmd.CombinedOutput()
	klog.Infof("running printenv: %v, output: %s", err, string(out))

	cmd = exec.Command("/bin/bash", "-c", "kubectl -v 6 apply -f https://raw.githubusercontent.com/Axpz/xMinima/refs/heads/master/script/xminima-cicd-job.yaml")
	out, err = cmd.CombinedOutput()
	klog.Infof("running CICD job: %v, output: %s", err, string(out))
	if err != nil {
		return false, err
	}
	return true, nil
}
