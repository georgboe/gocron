// Tests for gocron
package gocron

import (
	"fmt"
	"testing"
	"time"
)

var err = 1

func task() {
	fmt.Println("I am a running job.")
}

func taskWithParams(a int, b string) {
	fmt.Println(a, b)
}

func TestSecond(*testing.T) {
	defaultScheduler.Every(1).Second().Do(task)
	defaultScheduler.Every(1).Second().Do(taskWithParams, 1, "hello")
	defaultScheduler.Start()
	time.Sleep(10 * time.Second)
}

func TestDaylightSavingsTime(t *testing.T) {
	loc, _ := time.LoadLocation("Europe/Oslo")
	utc, _ := time.LoadLocation("UTC")
	scheduler := NewScheduler()
	scheduler.Clock = MockClock{now: time.Date(2017, 10, 26, 20, 00, 00, 00, loc)}

	ChangeLoc(loc)
	scheduler.Every(1).Monday().At("12:05").Do(func() {})

	_, next := scheduler.NextRun()

	expected := time.Date(2017, 10, 30, 11, 05, 0, 0, utc)
	if next.UTC() != expected {
		t.Errorf("%v not equal to %v", next, expected)
	}
}

func TestMonthChange(t *testing.T) {
	utc, _ := time.LoadLocation("UTC")
	scheduler := NewScheduler()
	scheduler.Clock = MockClock{now: time.Date(2017, 10, 04, 20, 00, 00, 00, utc)}

	ChangeLoc(utc)
	scheduler.Every(1).Friday().At("12:05").Do(func() {})

	_, next := scheduler.NextRun()

	expected := time.Date(2017, 10, 06, 12, 05, 0, 0, utc)
	if next.UTC() != expected {
		t.Errorf("%v not equal to %v", next, expected)
	}
}
