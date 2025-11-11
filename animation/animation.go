package animation

type Animation struct {
	first        int
	last         int
	step         int
	speedInTps   float32
	frameCounter float32
	frame        int
}

func NewAnimation(first, last, step int, speed float32) *Animation {
	return &Animation{
		first:        first,
		last:         last,
		step:         step,
		speedInTps:   speed,
		frameCounter: speed,
		frame:        first,
	}
}

func (a *Animation) Frame() int {
	return a.frame
}

func (a *Animation) Update() {
	a.frameCounter -= 1.0
	if a.frameCounter >= 0 {
		return
	}

	a.frameCounter = a.speedInTps
	a.frame += a.step

	if a.frame > a.last {
		a.frame = a.first
	}
}
