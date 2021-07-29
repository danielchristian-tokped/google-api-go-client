// Copyright 2021 Google LLC.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mock

import (
	"fmt"

	"github.com/danielchristian-tokped/google-api-go-client/googleapi"
	"github.com/danielchristian-tokped/google-api-go-client/translate/v3"
)

// TranslateTextCall is used to translate text and is fullfilled by a
// `translate.ProjectsTranslateTextCall`.
type TranslateTextCall interface {
	Do(opts ...googleapi.CallOption) (*translate.TranslateTextResponse, error)
}

// TranslateTextLowLevel executes the call and returns the translated text.
func TranslateTextLowLevel(call TranslateTextCall) (string, error) {
	resp, err := call.Do()
	if err != nil {
		return "", fmt.Errorf("unable to translate text: %v", err)
	}
	return resp.Translations[0].TranslatedText, nil
}
