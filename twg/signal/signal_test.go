package signal

import "testing"

func TestHandler(t *testing.T) {
	t.Error("here is a plain message") // Print the message but didn't stop the test
	t.Log("Showing other message")
	t.Fatal("this is the last message") // Print message and stop test
	t.Log("This message didn't show")

	/*
		t.Error => use in case to get more information continuing with the test
		t.Fatal => in case that you can not get more useful information
	*/
}
