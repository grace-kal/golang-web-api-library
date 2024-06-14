package models

import (
	"WebApiLibrary/database"
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAllBooks(c *gin.Context) {
	rows, err := database.GetDb().Query("SELECT * FROM books")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var books []Book
	for rows.Next() {
		var book Book
		if err := rows.Scan(&book.ID, &book.Title, &book.ISBN, &book.Author, &book.Release); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		books = append(books, book)
	}

	c.JSON(http.StatusOK, books)
}

func GetBookByID(c *gin.Context) {
	id := c.Param("id")
	row := database.GetDb().QueryRow("SELECT * FROM books WHERE id = ?", id)

	var book Book
	if err := row.Scan(&book.ID, &book.Title, &book.ISBN, &book.Author, &book.Release); err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, book)
}

func CreateBook(c *gin.Context) {
	var book Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := database.GetDb().Exec("INSERT INTO books (title, isbn, author, release) VALUES (?, ?, ?, ?)", book.Title, book.ISBN, book.Author, book.Release)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	id, err := result.LastInsertId()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	book.ID = int(id)
	c.JSON(http.StatusCreated, book)
}

func UpdateBook(c *gin.Context) {
	id := c.Param("id")

	var existingBook Book
	row := database.GetDb().QueryRow("SELECT * FROM books WHERE id = ?", id)
	err := row.Scan(&existingBook.ID, &existingBook.Title, &existingBook.ISBN, &existingBook.Author, &existingBook.Release)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	var updates map[string]interface{}
	if err := c.ShouldBindJSON(&updates); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if title, ok := updates["title"].(string); ok {
		existingBook.Title = title
	}
	if isbn, ok := updates["isbn"].(string); ok {
		existingBook.ISBN = isbn
	}
	if author, ok := updates["author"].(string); ok {
		existingBook.Author = author
	}
	if release, ok := updates["release"].(float64); ok {
		existingBook.Release = int(release)
	}

	_, err = database.GetDb().Exec("UPDATE books SET title = ?, isbn = ?, author = ?, release = ? WHERE id = ?", existingBook.Title, existingBook.ISBN, existingBook.Author, existingBook.Release, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, existingBook)
}

func DeleteBook(c *gin.Context) {
	id := c.Param("id")
	_, err := database.GetDb().Exec("DELETE FROM books WHERE id = ?", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
