Library Management System Documentation
Overview
The Library Management System is a console-based application developed in Go to manage books and members in a library. The system allows users to perform various operations such as adding and removing books, borrowing and returning books, and listing available and borrowed books. The application is designed using a modular approach with clear separation of concerns.

Folder Structure
go
Copy code
library_management/
├── main.go
├── controllers/
│ └── library_controller.go
├── models/
│ └── book.go
│ └── member.go
├── services/
│ └── library_service.go
├── docs/
│ └── documentation.md
└── go.mod
Components
main.go
Description: The entry point of the application. It initializes the application and starts the library management system.

Key Responsibilities:

Initializes the LibraryManager instance.
Calls the StartLibraryManagement function from the controllers package to begin user interaction.
controllers/library_controller.go
Description: Handles console input and invokes the appropriate service methods based on user choices.

Key Functions:

StartLibraryManagement(): Main function that presents the user interface and processes user input.
addBook(): Collects book details from the user and calls the AddBook method in the service layer.
removeBook(): Requests book ID from the user and calls the RemoveBook method in the service layer.
borrowBook(): Requests book ID and member ID from the user, then calls the BorrowBook method.
returnBook(): Requests book ID and member ID from the user, then calls the ReturnBook method.
listAvailableBooks(): Lists all available books by calling ListAvailableBooks.
listBorrowedBooks(): Requests member ID and lists borrowed books by calling ListBorrowedBooks.
addMember(): Collects member details from the user and calls the AddMember method in the service layer.
models/book.go
Description: Defines the Book struct used to represent books in the library.

Struct:

Book: Contains fields for ID, Title, Author, and Status.
models/member.go
Description: Defines the Member struct used to represent library members.

Struct:

Member: Contains fields for ID, Name, and BorrowedBooks (a slice of Book).
services/library_service.go
Description: Contains the business logic and data manipulation functions. Implements the LibraryManager interface.

Interface:

LibraryManager: Defines methods for adding and removing books, borrowing and returning books, listing available and borrowed books, and adding members.
Struct:

Library: Implements LibraryManager and maintains maps for books and members.
Methods:

AddBook(book models.Book) error: Adds a new book to the library. Returns an error if the book already exists.
RemoveBook(bookID int) error: Removes a book by ID. Returns an error if the book is not found.
BorrowBook(bookID int, memberID int) error: Allows a member to borrow a book if it is available. Returns errors for book not found, book already borrowed, or member not found.
ReturnBook(bookID int, memberID int) error: Allows a member to return a borrowed book. Returns errors for book not found, book not borrowed, or member not found.
ListAvailableBooks() []models.Book: Lists all available books.
ListBorrowedBooks(memberID int) ([]models.Book, error): Lists all books borrowed by a specific member. Returns an error if the member is not found.
AddMember(member models.Member) error: Adds a new member. Returns an error if the member already exists.
Error Handling
The system handles various errors such as:

Book Not Found: When attempting to borrow, return, or remove a book that does not exist.
Member Not Found: When attempting to borrow or return a book for a non-existent member.
Book Already Borrowed: When attempting to borrow a book that is already borrowed.
Member Already Exists: When attempting to add a member who already exists.
Usage
Run the Application: Execute main.go to start the Library Management System.
Interact with the System: Follow the console prompts to perform actions such as adding books, borrowing books, etc.
Exit: Choose the exit option from the menu to close the application.
