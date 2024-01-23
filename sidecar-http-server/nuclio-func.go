/*
Copyright 2023 The Nuclio Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package sidecar_http_server

import (
	"github.com/valyala/fasthttp"
	"net/url"
)

func Handler(context *nuclio.Context, event nuclio.Event) (interface{}, error) {
	context.Logger.Info("Got request, sending it to sidecar container")
	req := fasthttp.AcquireRequest()
	req.SetBody(event.GetBody())
	sidecarHost, _ := url.JoinPath("http://0.0.0.0:9000", event.GetPath())
	req.SetRequestURI(sidecarHost)
	resp := fasthttp.AcquireResponse()
	err := fasthttp.Do(req, resp)
	return nuclio.Response{
		StatusCode:  resp.StatusCode(),
		ContentType: "application/text",
		Body:        resp.Body(),
	}, err
}
