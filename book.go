/*
 * Copyright 2019 Shadai Rafael Lopez Garcia
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
package book

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Book struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	ISBN   string `json:"isbn"`
}

const (
	LocationBooks = "/api/books"
	LocationBook  = "/api/books/"
)

func (b *Book) encodeBook() []byte {
	j, err := json.Marshal(b)
	if err != nil {
		panic(err)
	}
	return j
}

func decodeBook(data []byte) Book {
	b := Book{}
	e := json.Unmarshal(data, &b)
	if e != nil {
		panic(e)
	}
	return b
}

var mBooks = map[string]Book{
	"123456789": Book{Title: "Effective Go", Author: "Shadai Lopez", ISBN: "123456789"},
	"123456790": Book{Title: "Effective Go 2", Author: "Shadai Lopez", ISBN: "123456790"},
}

func BooksHandler(rw http.ResponseWriter, r *http.Request) {
	switch method := r.Method; method {
	case http.MethodGet:
		books := getAllBooks()
		//rw.Header().Add("Content-Type", "application/json; charset=utf-8")
		writeBooks(&rw, books)
	case http.MethodPost:
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
		} else {
			book := decodeBook(body)
			i, c := createBook(book)
			if c {
				rw.Header().Add("Location", (LocationBooks + "/" + i))
				rw.WriteHeader(http.StatusCreated)
			} else {
				rw.WriteHeader(http.StatusConflict)
			}
		}
	default:
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte("Unsupported request method"))
	}

}

func BookHandler(rw http.ResponseWriter, r *http.Request) {
	isbn := r.URL.Path[len(LocationBook):]
	b, ok := mBooks[isbn]
	switch method := r.Method; method {
	case http.MethodGet:
		if ok == true {
			rw.Write(b.encodeBook())
			rw.WriteHeader(http.StatusOK)
			rw.Header().Add("Content-Type", "application/json; charset=utf-8")
		} else {
			rw.WriteHeader(http.StatusNotFound)
			rw.Write([]byte("Book not found"))
		}
	case http.MethodPut:
		if ok == true {
			body, err := ioutil.ReadAll(r.Body)
			if err != nil {
				rw.WriteHeader(http.StatusInternalServerError)
			} else {
				bNew := decodeBook(body)
				updateBook(bNew, isbn)
				rw.WriteHeader(http.StatusOK)
				rw.Write([]byte("Success"))
			}
		} else {
			rw.WriteHeader(http.StatusNotFound)
			rw.Write([]byte("Book not found"))
		}
	case http.MethodDelete:
		if ok == true {
			delete(mBooks, isbn)
			rw.Write([]byte("Success"))
		} else {
			rw.WriteHeader(http.StatusNotFound)
			rw.Write([]byte("Book not found"))
		}
	default:
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte("Unsupported request method"))
	}

}

func getAllBooks() []Book {
	allBooks := []Book{}
	for _, book := range mBooks {
		allBooks = append(allBooks, book)
	}
	return allBooks
}

func writeBooks(rw *http.ResponseWriter, books []Book) {
	for _, book := range books {
		(*rw).Write(book.encodeBook())
	}
}

func createBook(b Book) (i string, c bool) {
	_, ok := mBooks[b.ISBN]
	if ok == false {
		mBooks[b.ISBN] = b
		return b.ISBN, true
	}
	return "", false
}

func updateBook(bNew Book, isbn string) {
	mBooks[isbn] = bNew
}
