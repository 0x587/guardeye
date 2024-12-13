package listener

import (
	"time"

	"github.com/samber/lo"
)

func NewScheduler() *Scheduler {
	return &Scheduler{
		Listener: NewResListener(),
	}
}

type Scheduler struct {
	Listener *ResListener
}

type Schedule struct {
	TimeWhere *struct {
		StartAt time.Time
		EndAt   time.Time
	}
	SourceWhere []*SourceEntry
	Result      []*ResultEntry
}

func (s *Scheduler) GetSchedule() *Schedule {
	res := &Schedule{
		SourceWhere: s.Listener.qe.ses,
		Result:      s.Listener.qe.res,
	}
	for _, r := range res.Result {
		for _, s := range res.SourceWhere {
			s.NeedKeys = append(s.NeedKeys, r.Value.SourceDepend.NeedKeys...)
			s.NeedKeys = lo.Uniq(s.NeedKeys)
		}
	}
	return res
}
