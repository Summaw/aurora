package style

import (
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/Summaw/aurora/pkg/color"
)

type Spinner struct {
	message  string
	frames   []string
	interval time.Duration
	gradient color.Gradient
	output   *os.File
	stop     chan struct{}
	done     chan struct{}
	mu       sync.Mutex
	running  bool
}

var defaultFrames = []string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"}

func NewSpinner(message string) *Spinner {
	s := &Spinner{
		message:  message,
		frames:   defaultFrames,
		interval: 80 * time.Millisecond,
		gradient: color.GradientAurora,
		output:   os.Stdout,
		stop:     make(chan struct{}),
		done:     make(chan struct{}),
	}
	s.Start()
	return s
}

func (s *Spinner) Frames(frames []string) *Spinner {
	s.frames = frames
	return s
}

func (s *Spinner) Interval(d time.Duration) *Spinner {
	s.interval = d
	return s
}

func (s *Spinner) Gradient(name string) *Spinner {
	s.gradient = color.GetGradient(name)
	return s
}

func (s *Spinner) Start() {
	s.mu.Lock()
	if s.running {
		s.mu.Unlock()
		return
	}
	s.running = true
	s.mu.Unlock()

	go func() {
		i := 0
		for {
			select {
			case <-s.stop:
				close(s.done)
				return
			default:
				frame := s.frames[i%len(s.frames)]
				coloredFrame := s.gradient.Apply(frame)
				fmt.Fprintf(s.output, "\r  %s %s", coloredFrame, s.message)
				i++
				time.Sleep(s.interval)
			}
		}
	}()
}

func (s *Spinner) Stop() {
	s.mu.Lock()
	if !s.running {
		s.mu.Unlock()
		return
	}
	s.running = false
	s.mu.Unlock()

	close(s.stop)
	<-s.done
	fmt.Fprint(s.output, "\r\033[K")
}

func (s *Spinner) Success(msg string) {
	s.Stop()
	icon := color.Colorize("✓", color.Green)
	fmt.Fprintf(s.output, "  %s %s\n", icon, msg)
}

func (s *Spinner) Fail(msg string) {
	s.Stop()
	icon := color.Colorize("✖", color.Red)
	fmt.Fprintf(s.output, "  %s %s\n", icon, msg)
}

func (s *Spinner) Warn(msg string) {
	s.Stop()
	icon := color.Colorize("⚠", color.Yellow)
	fmt.Fprintf(s.output, "  %s %s\n", icon, msg)
}

func (s *Spinner) Info(msg string) {
	s.Stop()
	icon := color.Colorize("●", color.Blue)
	fmt.Fprintf(s.output, "  %s %s\n", icon, msg)
}
