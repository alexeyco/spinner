package spinner

import "time"

func ExampleNew() {
	s := New(Dots)
	s.Tick("Dots", func(s *Spinner) {
		time.Sleep(3 * time.Second)
	})

	s = New(Dots2)
	s.Tick("Dots2", func(s *Spinner) {
		time.Sleep(3 * time.Second)
	})

	s = New(Line)
	s.Tick("Line", func(s *Spinner) {
		time.Sleep(3 * time.Second)
	})

	s = New(Star)
	s.Tick("Star", func(s *Spinner) {
		time.Sleep(3 * time.Second)
	})

	s = New(Balloon)
	s.Tick("Balloon", func(s *Spinner) {
		time.Sleep(3 * time.Second)
	})
}

func ExampleSpinner_Done() {
	s := New(Dots)
	s.Tick("Wait for something", func(s *Spinner) {
		time.Sleep(3 * time.Second)
		s.Done("Complete")
	})
}
