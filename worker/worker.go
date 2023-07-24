package worker

import (
	"context"
	"sync"
	"time"

	"github.com/pcpratheesh/go-healthwatch/constants"
	"github.com/pcpratheesh/go-healthwatch/service"
	"github.com/pcpratheesh/go-healthwatch/utils"
	"github.com/sirupsen/logrus"
)

type Task struct {
	name     string
	interval time.Duration
	job      service.ServiceChecker
}

func NewTask(name string, j service.ServiceChecker, interval time.Duration) *Task {
	return &Task{
		name:     name,
		interval: interval,
		job:      j,
	}
}

type Tasks []*Task

func NewTasks(t ...*Task) Tasks {
	return t
}
func NewWorker() *Tasks {
	return &Tasks{}
}

// Add
// Add worker job into the pool
func (t Tasks) Add(name string, j service.ServiceChecker, interval time.Duration) Tasks {
	return append(t, NewTask(name, j, interval))
}

// Start the eexecution from pool
func (t Tasks) Start(ctx context.Context) {
	var wg sync.WaitGroup
	for i := range t {
		wg.Add(1)
		go func(t *Task) {
			defer wg.Done()
			t.Run(ctx)
		}(t[i])
	}
	wg.Wait()
}

// Run the tasks periodically
func (j *Task) Run(ctx context.Context) {
	if j.interval <= 0 {
		logrus.Infof("No interval found, [%s] checks once\n", j.name)
		if err := j.job.Check(ctx); err != nil && ctx.Err() == nil {
			utils.SendStatusNotification(j.job, constants.Failure, err)
			return
		}
		utils.SendStatusNotification(j.job, constants.Success, nil)
		return
	}
	logrus.Infof("[%s] checks at every %s\n", j.name, j.interval)
	t := time.NewTicker(j.interval)
	for {
		select {
		case <-t.C:
			if err := j.job.Check(ctx); err != nil && ctx.Err() == nil {
				utils.SendStatusNotification(j.job, constants.Failure, err)
				continue
			}
			utils.SendStatusNotification(j.job, constants.Success, nil)
			// logrus.Infof("[%v] health check success", j.name)

		case <-ctx.Done():
			t.Stop()
			logrus.Infof("[%s] terminated\n", j.name)
			return
		}
	}
}
