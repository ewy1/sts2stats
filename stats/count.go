package stats

import "sts2stats/model"

type Counter interface {
	Count(run model.Run) error
}

var Counters = []Counter{}

type CountFunc func(run model.Run) error

type CountFuncWrapper struct {
	f CountFunc
}

func (c CountFuncWrapper) Count(run model.Run) error {
	return c.f(run)
}

func Func(f CountFunc) Counter {
	return CountFuncWrapper{f: f}
}

func Count(run model.Run) error {
	for _, c := range Counters {
		err := c.Count(run)
		if err != nil {
			return err
		}
	}
	return nil
}
