package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const finalWord = "Go!"
const countdownStart = 3
const writeOp = "write"
const sleepOp = "sleep"

type Sleeper interface {
	Sleep()
}

type DefaultSleeper struct {
}

type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

func (cs *ConfigurableSleeper) Sleep() {
	cs.sleep(cs.duration)
}

type SpyTime struct {
	durationSlept time.Duration
}

func (s *SpyTime) Sleep(duration time.Duration) {
	s.durationSlept = duration
}

func (s *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

// Ordered list of ops (for improved testing) - implements ops for BOTH Sleep() + Write()
type CountdownOperationsSpy struct {
	Calls []string
}

// implements Sleep
func (s *CountdownOperationsSpy) Sleep() {
	s.Calls = append(s.Calls, sleepOp)
}

// implements io.Writer
func (s *CountdownOperationsSpy) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, writeOp)
	return
}

func Countdown(writer io.Writer, s Sleeper) {
	// Mock out "time.Sleep" so that we do NOT have to wait 4 seconds for each test cycle

	// for i := countdownStart; i > 0; i-- {
	// 	fmt.Fprintln(writer, i)
	// }
	// for i := countdownStart; i > 0; i-- {
	// 	s.Sleep()
	// }
	// s.Sleep()
	// fmt.Fprint(writer, finalWord)

	for i := countdownStart; i > 0; i-- {
		s.Sleep()
		fmt.Fprintln(writer, i)
	}
	s.Sleep()
	fmt.Fprint(writer, finalWord)

	// for i := countdownStart; i > 0; i-- {
	// 	s.Sleep()
	// 	fmt.Fprintln(writer, i)
	// }
	// s.Sleep()
	// fmt.Fprint(writer, finalWord)
}

func main() {
	Countdown(os.Stdout, &DefaultSleeper{})
}
