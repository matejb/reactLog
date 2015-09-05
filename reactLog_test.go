package reactLog

import (
	"bytes"
	"testing"
)

func TestSimpleWrite(t *testing.T) {
	buf := &bytes.Buffer{}

	logger := New(buf)

	_, err := logger.Write([]byte("This is a test"))
	if err != nil {
		t.Fatal("Unxepceted error, ", err)
	}

	if buf.String() != "This is a test" {
		t.Fatalf("Expected %v, recived %v", "This is a test", buf.String())
	}
}

func TestDiscard(t *testing.T) {
	buf := &bytes.Buffer{}

	logger := New(buf)
	logger.AddReaction("INFO", &Discard{})

	_, err := logger.Write([]byte("INFO This is a test"))
	if err != nil {
		t.Fatal("Unxepceted error, ", err)
	}
	if buf.String() != "" {
		t.Fatalf("Expected %v, recived %v", "", buf.String())
	}

	_, err = logger.Write([]byte("NOT_INFO This is a test"))
	if err != nil {
		t.Fatal("Unxepceted error, ", err)
	}
	if buf.String() != "NOT_INFO This is a test" {
		t.Fatalf("Expected %v, recived %v", "NOT_INFO This is a test", buf.String())
	}
}

func TestRedirect(t *testing.T) {
	buf := &bytes.Buffer{}
	redirectBuf := &bytes.Buffer{}

	logger := New(buf)
	logger.AddReaction("INFO", &Redirect{redirectBuf})

	_, err := logger.Write([]byte("INFO This is a test"))
	if err != nil {
		t.Fatal("Unxepceted error, ", err)
	}
	if buf.String() != "" {
		t.Fatalf("Expected %v, recived %v", "", buf.String())
	}
	if redirectBuf.String() != "INFO This is a test" {
		t.Fatalf("Expected %v, recived %v", "INFO This is a test", redirectBuf.String())
	}

	redirectBuf.Reset()
	_, err = logger.Write([]byte("NOT_INFO This is a test"))
	if err != nil {
		t.Fatal("Unxepceted error, ", err)
	}
	if buf.String() != "NOT_INFO This is a test" {
		t.Fatalf("Expected %v, recived %v", "NOT_INFO This is a test", buf.String())
	}
	if redirectBuf.String() != "" {
		t.Fatalf("Expected %v, recived %v", "", redirectBuf.String())
	}
}

func TestCopy(t *testing.T) {
	buf := &bytes.Buffer{}
	copyBuf := &bytes.Buffer{}

	logger := New(buf)
	logger.AddReaction("INFO", &Copy{copyBuf})

	_, err := logger.Write([]byte("INFO This is a test"))
	if err != nil {
		t.Fatal("Unxepceted error, ", err)
	}
	if buf.String() != "INFO This is a test" {
		t.Fatalf("Expected %v, recived %v", "INFO This is a test", buf.String())
	}
	if copyBuf.String() != "INFO This is a test" {
		t.Fatalf("Expected %v, recived %v", "INFO This is a test", copyBuf.String())
	}
}
