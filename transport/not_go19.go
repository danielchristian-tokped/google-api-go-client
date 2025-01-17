// Copyright 2018 Google LLC.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build !go1.9

package transport

import (
	"context"

	"github.com/danielchristian-tokped/google-api-go-client/internal"
	"github.com/danielchristian-tokped/google-api-go-client/option"
	"golang.org/x/oauth2/google"
)

// Creds constructs a google.DefaultCredentials from the information in the options,
// or obtains the default credentials in the same way as google.FindDefaultCredentials.
func Creds(ctx context.Context, opts ...option.ClientOption) (*google.DefaultCredentials, error) {
	var ds internal.DialSettings
	for _, opt := range opts {
		opt.Apply(&ds)
	}
	return internal.Creds(ctx, &ds)
}
