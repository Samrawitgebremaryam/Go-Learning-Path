package controllers

import (
	"fmt"
	"library_management/models"
	"library_management/services"
)

var libService services.LibraryManager

func StartLibraryManagement() {
	libService = services.NewLibrary()

	for {
		fmt.Println("\t\t\t############################")
		fmt.Println("\t\t\tWelcome to Library Management System")
		fmt.Println("\t\t\t1. Add a new book")
		fmt.Println("\t\t\t2. Remove an existing book")
		fmt.Println("\t\t\t3. Borrow a book")
		fmt.Println("\t\t\t4. Return a book")
		fmt.Println("\t\t\t5. List all available books")
		fmt.Println("\t\t\t6. List all borrowed books by a member")
		fmt.Println("\t\t\t7. Add a new member")
		fmt.Println("\t\t\t8. Exit")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			addBook()
		case 2:
			removeBook()
		case 3:
			borrowBook()
		case 4:
			returnBook()
		case 5:
			listAvailableBooks()
		case 6:
			listBorrowedBooks()
		case 7:
			addMember()
		case 8:
			fmt.Println("Exiting the system...")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}

func addBook() {
	var book models.Book
	fmt.Print("Enter Book ID: ")
	fmt.Scan(&book.ID)
	fmt.Print("Enter Book Title: ")
	fmt.Scan(&book.Title)
	fmt.Print("Enter Book Author: ")
	fmt.Scan(&book.Author)
	book.Status = "Available"
	err := libService.AddBook(book)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Book added successfully.")
	}
}

func removeBook() {
	var bookID int
	fmt.Print("Enter the ID of the book you want to remove: ")
	fmt.Scan(&bookID)
	err := libService.RemoveBook(bookID)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Book removed successfully.")
	}
}

func borrowBook() {
	var bookID, memberID int
	fmt.Print("Enter the ID of the book you want to borrow: ")
	fmt.Scan(&bookID)
	fmt.Print("Enter your member ID: ")
	fmt.Scan(&memberID)
	err := libService.BorrowBook(bookID, memberID)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Book borrowed successfully.")
	}
}

func returnBook() {
	var bookID, memberID int
	fmt.Print("Enter the ID of the book you want to return: ")
	fmt.Scan(&bookID)
	fmt.Print("Enter your member ID: ")
	fmt.Scan(&memberID)
	err := libService.ReturnBook(bookID, memberID)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Book returned successfully.")
	}
}

func listAvailableBooks() {
	books := libService.ListAvailableBooks()
	fmt.Println("Available Books:")
	for _, book := range books {
		fmt.Printf("ID: %d, Title: %s, Author: %s\n", book.ID, book.Title, book.Author)
	}
}

func listBorrowedBooks() {
	var memberID int
	fmt.Print("Enter your member ID: ")
	fmt.Scan(&memberID)
	books, err := libService.ListBorrowedBooks(memberID)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Borrowed Books:")
		for _, book := range books {
			fmt.Printf("ID: %d, Title: %s, Author: %s\n", book.ID, book.Title, book.Author)
		}
	}
}

func addMember() {
	var member models.Member
	fmt.Print("Enter Member ID: ")
	fmt.Scan(&member.ID)
	fmt.Print("Enter Member Name: ")
	fmt.Scan(&member.Name)
	err := libService.AddMember(member)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Member added successfully.")
	}
}
