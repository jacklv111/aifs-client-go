/*
 * Created on Tue Jul 11 2023
 *
 * Copyright (c) 2023 Company-placeholder. All rights reserved.
 *
 * Author Yubinlv.
 */

package httpclient

import "net/http"

//go:generate mockgen -source=interface.go -destination=./mock/mock_interface.go -package=mock

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}
