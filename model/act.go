package model

type Act struct {
	Floor int
	Label string
	Key   string
}

var Acts = make(map[string]*Act)
