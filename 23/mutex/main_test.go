package main

import "testing"

func Test_updateMessage(t *testing.T) {
	msg = "Hello, world!"

	wg.Add(1)
	go updateMessage("x")
	go updateMessage("Hi")
	wg.Wait()
	if msg != "Hi" {
		t.Error("incorrect value in msg")
	}
}
