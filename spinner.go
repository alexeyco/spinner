package spinner

import "time"

const (
	duration = 130 * time.Millisecond
	reset    = "\033[2K\r"
)

var (
	// Dots spinner frames (see example)
	Dots = []string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"}
	// Dots2 spinner frames (see example)
	Dots2 = []string{"⠁", "⠂", "⠄", "⡀", "⢀", "⠠", "⠐", "⠈"}
	// Line spinner frames (see example)
	Line = []string{"-", "\\", "|", "/"}
	// Star spinner frames (see example)
	Star = []string{"+", "x", "*"}
	// Balloon spinner frames (see example)
	Balloon = []string{".", "o", "O", "°", "O", "o", "."}
)

// Handler spinner handler
type Handler func(s *Spinner)

// Spinner spinner struct
type Spinner struct {
	frames []string
	len    int
	done   string
}

// Tick executes unpredictable task with specified message and spinner
func (s *Spinner) Tick(message string, h Handler) {
	stop := make(chan bool)

	go s.spin(message, stop)
	h(s)

	stop <- true
	close(stop)

	s.reset()
	if s.done != "" {
		println(s.done)
	}

	s.done = ""
}

// Done draws complete message
func (s *Spinner) Done(message string) {
	s.done = "  " + message
}

func (s *Spinner) reset() {
	print(reset)
}

func (s *Spinner) draw(message string, cur int) {
	print(s.frames[cur] + " " + message)
}

func (s *Spinner) next(cur int) int {
	cur++
	if cur >= s.len-1 {
		cur = 0
	}

	return cur
}

func (s *Spinner) spin(message string, stop chan bool) {
	var cur int
	s.draw(message, cur)

	ticker := time.NewTicker(duration)

	for {
		select {
		case <-stop:
			return
		case <-ticker.C:
			cur = s.next(cur)

			s.reset()
			s.draw(message, cur)
		}
	}
}

// New returns new spinner with specified frames
func New(frames []string) *Spinner {
	return &Spinner{
		frames: frames,
		len:    len(frames),
	}
}
