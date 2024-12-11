package listener

import "time"

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
	return &Schedule{
		SourceWhere: s.Listener.qe.ses,
		Result:      s.Listener.qe.res,
	}
}
