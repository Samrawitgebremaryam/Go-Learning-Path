package controllers

import (
	"bufio"
	"fmt"
	"library_management/models"
	"library_management/services"
	"os"
	"strconv"
	"strings"
)

var libService services.LibraryManager

func StartLibraryManagement() {
	libService = services.NewLibrary()

	for {
		displayMainMenu()

		var choice int
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter your choice: ")
		choice = readInt(reader)

		switch choice {
		case 1:
			addBook(reader)
		case 2:
			removeBook(reader)
		case 3:
			borrowBook(reader)
		case 4:
			returnBook(reader)
		case 5:
			listAvailableBooks()
		case 6:
			listBorrowedBooks(reader)
		case 7:
			addMember(reader)
		case 8:
			fmt.Println("Exiting the system... Have a great day!")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
		pause()
	}
}

func displayMainMenu() {
	fmt.Println("\n\t\t\t############################################")
	fmt.Println("\t\t\t#                                          #")
	fmt.Println("\t\t\t#       Welcome to Library Management      #")
	fmt.Println("\t\t\t#                  System                  #")
	fmt.Println("\t\t\t#                                          #")
	fmt.Println("\t\t\t############################################")
	fmt.Println("\t\t\t# 1. Add a new book                        #")
	fmt.Println("\t\t\t# 2. Remove an existing book               #")
	fmt.Println("\t\t\t# 3. Borrow a book                         #")
	fmt.Println("\t\t\t# 4. Return a book                         #")
	fmt.Println("\t\t\t# 5. List all available books              #")
	fmt.Println("\t\t\t# 6. List all borrowed books by a member   #")
	fmt.Println("\t\t\t# 7. Add a new member                      #")
	fmt.Println("\t\t\t# 8. Exit                                  #")
	fmt.Println("\t\t\t############################################")
}

func addBook(reader *bufio.Reader) {
	var book models.Book
	fmt.Println("\n--- Add a New Book ---")
	fmt.Print("Enter Book ID: ")
	book.ID = readInt(reader)
	fmt.Print("Enter Book Title: ")
	book.Title = readString(reader)
	fmt.Print("Enter Book Author: ")
	book.Author = readString(reader)
	book.Status = "Available"
	err := libService.AddBook(book)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Book added successfully.")
	}
}

func removeBook(reader *bufio.Reader) {
	fmt.Println("\n--- Remove an Existing Book ---")
	fmt.Print("Enter the ID of the book you want to remove: ")
	bookID := readInt(reader)
	err := libService.RemoveBook(bookID)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Book removed successfully.")
	}
}

func borrowBook(reader *bufio.Reader) {
	fmt.Println("\n--- Borrow a Book ---")
	fmt.Print("Enter the ID of the book you want to borrow: ")
	bookID := readInt(reader)
	fmt.Print("Enter your member ID: ")
	memberID := readInt(reader)
	err := libService.BorrowBook(bookID, memberID)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Book borrowed successfully.")
	}
}

func returnBook(reader *bufio.Reader) {
	fmt.Println("\n--- Return a Book ---")
	fmt.Print("Enter the ID of the book you want to return: ")
	bookID := readInt(reader)
	fmt.Print("Enter your member ID: ")
	memberID := readInt(reader)
	err := libService.ReturnBook(bookID, memberID)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Book returned successfully.")
	}
}

func listAvailableBooks() {
	fmt.Println("\n--- Available Books ---")
	books := libService.ListAvailableBooks()
	if len(books) == 0 {
		fmt.Println("No available books.")
	} else {
		fmt.Println("-----------------------------------------------------------------")
		fmt.Printf("| %-10s | %-30s | %-20s |\n", "Book ID", "Title", "Author")
		fmt.Println("-----------------------------------------------------------------")
		for _, book := range books {
			fmt.Printf("| %-10d | %-30s | %-20s |\n", book.ID, book.Title, book.Author)
		}
		fmt.Println("-----------------------------------------------------------------")
	}
}

func listBorrowedBooks(reader *bufio.Reader) {
	fmt.Println("\n--- Borrowed Books ---")
	fmt.Print("Enter your member ID: ")
	memberID := readInt(reader)
	books, err := libService.ListBorrowedBooks(memberID)
	if err != nil {
		fmt.Println("Error:", err)
	} else if len(books) == 0 {
		fmt.Println("No borrowed books.")
	} else {
		fmt.Println("-----------------------------------------------------------------")
		fmt.Printf("| %-10s | %-30s | %-20s |\n", "Book ID", "Title", "Author")
		fmt.Println("-----------------------------------------------------------------")
		for _, book := range books {
			fmt.Printf("| %-10d | %-30s | %-20s |\n", book.ID, book.Title, book.Author)
		}
		fmt.Println("-----------------------------------------------------------------")
	}
}

func addMember(reader *bufio.Reader) {
	var member models.Member
	fmt.Println("\n--- Add a New Member ---")
	fmt.Print("Enter Member ID: ")
	member.ID = readInt(reader)
	fmt.Print("Enter Member Name: ")
	member.Name = readString(reader)
	err := libService.AddMember(member)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Member added successfully.")
	}
}

func pause() {
	fmt.Print("Press 'Enter' to continue...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}

func readInt(reader *bufio.Reader) int {
	for {
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		value, err := strconv.Atoi(input)
		if err != nil {
			fmt.Print("Invalid input. Please enter a number: ")
		} else {
			return value
		}
	}
}

func readString(reader *bufio.Reader) string {
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}
