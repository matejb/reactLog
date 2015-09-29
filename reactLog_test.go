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

	_, err = logger.Write([]byte("ERROR This is a test"))
	if err != nil {
		t.Fatal("Unxepceted error, ", err)
	}
	if buf.String() != "ERROR This is a test" {
		t.Fatalf("Expected %v, recived %v", "ERROR This is a test", buf.String())
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
	_, err = logger.Write([]byte("ERROR This is a test"))
	if err != nil {
		t.Fatal("Unxepceted error, ", err)
	}
	if buf.String() != "ERROR This is a test" {
		t.Fatalf("Expected %v, recived %v", "ERROR This is a test", buf.String())
	}
	if redirectBuf.String() != "" {
		t.Fatalf("Expected %v, recived %v", "", redirectBuf.String())
	}
}

func TestCopy(t *testing.T) {
	buf := &bytes.Buffer{}
	copyBuf := &bytes.Buffer{}

	logger := New(buf)
	logger.AddReaction("user ID 85", &Copy{copyBuf})

	_, err := logger.Write([]byte("This error message concers user ID 85 and will be copied to copyBuf."))
	if err != nil {
		t.Fatal("Unxepceted error, ", err)
	}
	if buf.String() != "This error message concers user ID 85 and will be copied to copyBuf." {
		t.Fatalf("Expected %v, recived %v", "This error message concers user ID 85 and will be copied to copyBuf.", buf.String())
	}
	if copyBuf.String() != "This error message concers user ID 85 and will be copied to copyBuf." {
		t.Fatalf("Expected %v, recived %v", "This error message concers user ID 85 and will be copied to copyBuf.", copyBuf.String())
	}
}

func TestNonWord(t *testing.T) {
	buf := &bytes.Buffer{}

	logger := New(buf)
	logger.AddReaction("custom string sample", &Discard{})

	_, err := logger.Write([]byte("this is complex custom string sample sentence"))
	if err != nil {
		t.Fatal("Unxepceted error, ", err)
	}
	if buf.String() != "" {
		t.Fatalf("Expected %v, recived %v", "", buf.String())
	}
}

func TestPositions(t *testing.T) {
	buf := &bytes.Buffer{}

	logger := New(buf)
	logger.AddReaction("TRIGGER WORD", &Discard{})

	_, err := logger.Write([]byte("TRIGGER WORD at the begining"))
	if err != nil {
		t.Fatal("Unxepceted error, ", err)
	}
	if buf.String() != "" {
		t.Fatalf("Expected %v, recived %v", "", buf.String())
	}

	buf.Reset()
	_, err = logger.Write([]byte("At the end is TRIGGER WORD"))
	if err != nil {
		t.Fatal("Unxepceted error, ", err)
	}
	if buf.String() != "" {
		t.Fatalf("Expected %v, recived %v", "", buf.String())
	}

	buf.Reset()
	_, err = logger.Write([]byte("In the middle is TRIGGER WORD is'n it"))
	if err != nil {
		t.Fatal("Unxepceted error, ", err)
	}
	if buf.String() != "" {
		t.Fatalf("Expected %v, recived %v", "", buf.String())
	}
}

func TestUtf8(t *testing.T) {
	buf := &bytes.Buffer{}

	logger := New(buf)
	logger.AddReaction("ČćžđšŽĆš", &Discard{})

	_, err := logger.Write([]byte("This are croatian charcaters ČćžđšŽĆš"))
	if err != nil {
		t.Fatal("Unxepceted error, ", err)
	}
	if buf.String() != "" {
		t.Fatalf("Expected %v, recived %v", "", buf.String())
	}
}
