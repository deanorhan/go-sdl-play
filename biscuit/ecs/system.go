package ecs

type (
	System interface {
		Process(delta float64)
	}

	SystemInit interface {
		Init(*World)
	}
)
