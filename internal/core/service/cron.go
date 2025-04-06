package service

import (
	"github.com/RomanshkVolkov/test-api/internal/core/domain"
	"github.com/robfig/cron/v3"
)

func InitCron() {
	c := cron.New()

	c.AddFunc("@every 1m", WrapperFuncs(SyncElevaZapier))
	c.AddFunc("@every 1m", WrapperFuncs(SyncBoatyHealthChecker))

	c.Start()
}

func WrapperFuncs(fn func() domain.APIResponse[any]) func() {
	return func() {
		fn()
	}
}
