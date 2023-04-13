package main

import (
	"bufio"
	"io"
	"os"
	"strings"
	"testing"
)

func Test_isPrime(t *testing.T) {
	primeTests := []struct {
		name     string
		testNum  int
		expected bool
		msg      string
	}{
		{"prime", 7, true, "7 is a prime number!"},
		{"not prime", 8, false, "8 is not a prime number because it is divisible by 2!"},
		{"zero", 0, false, "0 is not prime, by definition!"},
		{"one", 1, false, "1 is not prime, by definition!"},
		{"negative number", -11, false, "Negative numbers are not prime, by definition!"},
	}

	for _, e := range primeTests {
		result, msg := isPrime(e.testNum)
		if e.expected && !result {
			t.Errorf("%s: expected true but got false", e.name)
		}

		if !e.expected && result {
			t.Errorf("%s: expected false but got true", e.name)
		}

		if e.msg != msg {
			t.Errorf("%s: expected %s but got %s", e.name, e.msg, msg)
		}
	}
}

func Test_prompt(t *testing.T) {
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	prompt()

	_ = w.Close()
	os.Stdout = oldStdout

	result, _ := io.ReadAll(r)
	if string(result) != "-> " {
		t.Fatal("wrong prompt")
	}

}

func Test_intro(t *testing.T) {
	previousOutput := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	intro()

	_ = w.Close()

	os.Stdout = previousOutput

	out, _ := io.ReadAll(r)

	if strings.Contains(string(out), "Enter whole number") {
		t.Fatal("wrong intro")
	}
}

func Test_checkNumbers(t *testing.T) {
	/*	scanner := bufio.NewScanner(strings.NewReader("q"))
		//scanner.Bytes()
		result, err := checkNumbers(scanner)
		if result != "" && err != true {
			t.Fatal()
		}*/

	checkNumbersTest := []struct {
		name                string
		value               string
		expectedStringValue string
		expectedBoolValue   bool
	}{
		{"exit signal", "q", "", true},
		{"wrong string value", "wrong number", "Please enter a whole number!", false},
		{"wrong float value", "7.5", "Please enter a whole number!", false},
		{"right int value", "32", "32 is not a prime number because it is divisible by 2!", false},
	}

	for _, test := range checkNumbersTest {
		scanner := bufio.NewScanner(strings.NewReader(test.value))
		result, boolVal := checkNumbers(scanner)
		if result != test.expectedStringValue && boolVal != test.expectedBoolValue {
			t.Fatal()
		}
	}
}

func Test_readUserInput(t *testing.T) {
	testChain := make(chan bool)
	go readUserInput(bufio.NewReader(strings.NewReader("123")), testChain)
	go readUserInput(bufio.NewReader(strings.NewReader("q")), testChain)
	<-testChain
	close(testChain)
}
