package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {

		fmt.Println("Please choose an option:")
		fmt.Println("1 . Count the frequency of words")
		fmt.Println("2 . Check if a word is a palindrome")
		fmt.Println("0 . Exit the program")

		choiceStr, _ := reader.ReadString('\n')
		choiceStr = strings.TrimSpace(choiceStr)
		choice, err := strconv.Atoi(choiceStr)
		if err != nil {
			fmt.Println("Invalid input, please enter a valid number.")
			pause()
			continue
		}

		switch choice {
		case 1:
			FrequencyCounter(reader)
		case 2:
			palindrome(reader)
		case 0:
			fmt.Println("\nExiting ... ")
			return
		default:
			fmt.Println("Invalid choice, please select a valid option.")
			pause()
		}
	}
}

func FrequencyCounter(reader *bufio.Reader) {
	count := make(map[string]int)

	fmt.Println("\n--- Word Frequency Counter ---")
	fmt.Println("Please enter the text:")

	text, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	text = strings.TrimSpace(text)
	text = strings.ToLower(text)
	text = removePunctuation(text)
	words := strings.Fields(text)

	for _, word := range words {
		count[word]++
	}

	fmt.Println("\nWord Frequencies:")
	for word, freq := range count {
		fmt.Printf("%s: %d\n", word, freq)
	}

	pause()
}

func palindrome(reader *bufio.Reader) {
	fmt.Println("\n--- Palindrome Checker ---")
	fmt.Print("Please enter a word or phrase: ")

	word, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	word = strings.TrimSpace(word)
	word = strings.ToLower(word)
	word = removePunctuation(word)

	left, right := 0, len(word)-1
	isPalindrome := true
	for left < right {
		if word[left] != word[right] {
			isPalindrome = false
			break
		}
		left++
		right--
	}

	if isPalindrome {
		fmt.Println("\nThis word or phrase is a palindrome.")
	} else {
		fmt.Println("\nThis word or phrase is not a palindrome.")
	}

	pause()
}

func removePunctuation(text string) string {
	var ans []rune
	for _, val := range text {
		if unicode.IsLetter(val) || unicode.IsDigit(val) || unicode.IsSpace(val) {
			ans = append(ans, val)
		}
	}
	return string(ans)
}

func pause() {
	fmt.Println("\nPress Enter to continue...")
	bufio.NewReader(os.Stdin).ReadString('\n')
}
