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

package gout

import (
	"io"
	"net/http"
	"os"
)

// DownloadFile downloads a file and saves it
func DownloadFile(url string, path string) error {
	// create new file
	output, err := os.Create(path)
	defer output.Close()
	if err != nil {
		return err
	}

	// get response
	response, err := http.Get(url)
	defer response.Body.Close()
	if err != nil {
		return err
	}

	// write to disk
	_, err = io.Copy(output, response.Body)
	if err != nil {
		return err
	}

	return nil
}
