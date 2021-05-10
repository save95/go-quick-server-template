package example

import (
	"log"

	"github.com/save95/go-pkg/job"
)

type simpleJob struct {
}

func NewSimpleJob() job.IJob {
	return &simpleJob{}
}

func (s simpleJob) Run() error {
	log.Println("example simple job, only print")

	return nil
}
