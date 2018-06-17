// Stefan Nilsson 2013-03-13

// This program implements an ELIZA-like oracle (en.wikipedia.org/wiki/ELIZA).
package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

const (
	enough = 2
	star   = "Pythia"
	venue  = "Delphi"
	prompt = "> "
)

func main() {
	fmt.Printf("Welcome to %s, the oracle at %s.\n", star, venue)
	fmt.Println("Your questions will be answered in due time.")

	oracle := Oracle()
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(prompt)
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		fmt.Printf("%s heard: %s\n", star, line)
		oracle <- line // The channel doesn't block.
	}
}

// Oracle returns a channel on which you can send your questions to the oracle.
// You may send as many questions as you like on this channel, it never blocks.
// The answers arrive on stdout, but only when the oracle so decides.
// The oracle also prints sporadic prophecies to stdout even without being asked.
func Oracle() chan<- string {
	questions := make(chan string, enough)
	prophecies := make(chan string, enough)

	go propheciator2(questions, prophecies)
	go answers(prophecies)

	return questions
}

// This is the oracle's secret algorithm.
// It waits for a while and then sends a message on the answer channel.
// TODO: make it better.
func prophecy(question string, answer chan<- string) {

	if question == "" {
		bs := []string{
			"...this is getting awkward.",
			"You're not gonna answer me?!",
			"I guess you're not very talktive...",
		}
		answer <- bs[rand.Intn(len(bs))]
	}


	// Find the longest word.
	longestWord := ""
	words := strings.Fields(question) // Fields extracts the words into a slice.
	for _, w := range words {
		if len(w) > len(longestWord) {
			longestWord = w
		}
	}

	// Cook up some pointless nonsense.
	nonsense := []string{
		"DDDDDJJJJJJ KHALED",
		"Anotha one",
		"MAJOR KEY ALERT",
		"bless upp!",
		"MAAAD fan luv",
		"#WeTheBest",
	}
	answer <- longestWord + "... " + nonsense[rand.Intn(len(nonsense))]
}
func answers(answers <- chan string) {
	for answer := range answers {
		fmt.Print("\r")
		for _, char := range answer {
			time.Sleep(time.Duration(100+rand.Intn(10)) * time.Millisecond)
			fmt.Print(string(char))
		}
		fmt.Print("\n", prompt)

	}
}

func propheciator(question string, ch1 chan<- string) {
	go prophecy(question, ch1)
}

func propheciator2(ch1 <-chan string, ch2 chan<- string) {
	for {
		time.Sleep(time.Duration(10+rand.Intn(5)) * time.Second)
		select{
		case question := <-ch1:
			go propheciator(question, ch2)
		default:
			go propheciator("", ch2)
		}
	}
}



func init() { // Functions called "init" are executed before the main function.
	// Use new pseudo random numbers every time.
	rand.Seed(time.Now().Unix())
}
