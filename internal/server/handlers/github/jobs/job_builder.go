package jobs

import (
	"fmt"

	"github.com/hashicorp/go-multierror"
	"k8s.io/klog/v2"
)

type JobBuilder struct {
	jobs map[string]Job
}

func NewJobBuilder() *JobBuilder {

	return &JobBuilder{
		jobs: map[string]Job{
			RepoName: &XMinimaCICDJob{},
		},
	}
}

func (builder *JobBuilder) Complete(jobs ...string) error {
	var allErrors *multierror.Error

	if len(jobs) > 0 {
		klog.Infof("running jobs: %v", jobs)
		for _, job := range jobs {
			klog.Infof("job: %s", job)
			if j, ok := builder.jobs[job]; ok {
				if err := j.Do(); err != nil {
					allErrors = multierror.Append(allErrors, fmt.Errorf("job %s failed: %v", job, err))
				}
			}
		}
	} else {
		klog.Infof("running all the jobs: %v", builder.jobs)
		for jobName, j := range builder.jobs {
			if err := j.Do(); err != nil {
				allErrors = multierror.Append(allErrors, fmt.Errorf("job %s failed: %v", jobName, err))
			}
		}
	}

	return allErrors.ErrorOrNil()
}
