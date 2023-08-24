// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package errors

import (
	"fmt"
	"log"
)


// Unwrap returns the result of calling the Unwrap method on err, if err's
// type contains an Unwrap method returning error.
// Otherwise, Unwrap returns nil.
//
// Unwrap returns nil if the Unwrap method returns []error.
func Unwrap(err error) error {
	u, ok := err.(interface {
		Unwrap() error
	})
	if !ok {
		return nil
	}
	return u.Unwrap()
}

// Wrap wraps an error with additional messages.
//
// If the given error err is not nil, Wrap creates a new error by combining
// messages from the msgs arguments with the text of the original error.
// It returns a new error with the combined data.
// If err is nil, the function returns nil.
func Wrap(err *error, msgs ...interface{}) error {
    if *err == nil {
        return nil
    }

    var message string
    for _, msg := range msgs {
        if msg != nil {
            if message != "" {
                message += "--->"
            }
            message += fmt.Sprint(msg)
        }
    }

    if message != "" {
        *err = New(message + "--->" + (*err).Error())
    }

    return *err
}

// WrapLog wraps an error with additional messages and performs logging.
//
// If the given error err is not nil, WrapLog creates a message by combining
// messages from the msgs arguments with the text of the original error.
// It then logs this message along with information about the request and error.
// If err is nil, the function does not take any action.
func WrapLog(err *error, req interface{}, msgs ...interface{}) {
    if *err == nil {
        return
    }

    var message string
    for _, msg := range msgs {
        if msg != nil {
            if message != "" {
                message += "--->"
            }
            message += fmt.Sprint(msg)
        }
    }

    log.Print(
        message,
        " request: ", req,
        " Error: ", *err,
    )
}

