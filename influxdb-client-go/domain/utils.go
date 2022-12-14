// Copyright 2020-2021 InfluxData, Inc. All rights reserved.
// Use of this source code is governed by MIT
// license that can be found in the LICENSE file.

package domain

import (
	"gerrit.o-ran-sc.org/r/scp/ric-app/kpimon/influxdb-client-go/api/http"
)

// ErrorToHTTPError creates http.Error from domain.Error
func ErrorToHTTPError(error *Error, statusCode int) *http.Error {
	err := &http.Error{
		StatusCode: statusCode,
		Code:       string(error.Code),
	}
	if error.Message != nil {
		err.Message = *error.Message
	}
	return err
}
