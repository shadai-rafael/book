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
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncodeBook(t *testing.T) {
	b := Book{Title: "Effective Go", Author: "Rafael", ISBN: "123456789"}
	jb := b.encodeBook()
	assert.Equal(t, `{"title":"Effective Go","author":"Rafael","isbn":"123456789"}`,
		string(jb), "book marshaling wrong")
}

func TestDecodeBook(t *testing.T) {
	j := []byte(`{"title":"Effective Go","author":"Rafael","isbn":"123456789"}`)
	book := decodeBook(j)
	assert.Equal(t, Book{Title: "Effective Go", Author: "Rafael", ISBN: "123456789"},
		book, "book unmarshaling wrong")
}
