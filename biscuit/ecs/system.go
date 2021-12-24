package ecs

type System interface {
	Init(*World)
	Process(delta float64)
}
