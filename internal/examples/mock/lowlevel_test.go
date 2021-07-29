// Copyright 2021 Google LLC.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mock

import (
	"testing"

	"github.com/danielchristian-tokped/google-api-go-client/googleapi"
	"github.com/danielchristian-tokped/google-api-go-client/translate/v3"
)

// mockCall fullfills the TranslateTextCall and matches the `Do` call on
// `translate.ProjectsTranslateTextCall`.
type mockCall struct{}

func (*mockCall) Do(opts ...googleapi.CallOption) (*translate.TranslateTextResponse, error) {
	resp := &translate.TranslateTextResponse{
		Translations: []*translate.Translation{
			{TranslatedText: "Hello World"},
		},
	}
	return resp, nil
}

func TestTranslateTextLowLevel(t *testing.T) {
	call := &mockCall{}
	text, err := TranslateTextLowLevel(call)
	if err != nil {
		t.Fatal(err)
	}
	if text != "Hello World" {
		t.Fatalf("got %q, want Hello World", text)
	}
}
