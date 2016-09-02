// Copyright (c) 2016. See AUTHORS file.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package strings

import (
	"crypto/rand"
)

const (
	// AlphaLowerCase alphabet a-z
	AlphaLowerCase = "abcdefghijklmnopqrstuvwxyz"
	// AlphaUpperCase alphabet A-Z
	AlphaUpperCase = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	// Numeric alphabet 0-9
	Numeric = "0123456789"
	// Alpha alphabet a-zA-Z
	Alpha = AlphaLowerCase + AlphaUpperCase
	// AlphaNumeric alphabet a-zA-Z0-9
	AlphaNumeric = Alpha + Numeric
)

// RandomString generates a random string
func RandomString(length int, alphabet string) string {
	alphabetLen := byte(len(alphabet))

	// make generate random byte array
	id := make([]byte, length)
	rand.Read(id)

	// replace rand num with char from alphabet
	for i, b := range id {
		id[i] = alphabet[b%alphabetLen]
	}

	return string(id)
}
