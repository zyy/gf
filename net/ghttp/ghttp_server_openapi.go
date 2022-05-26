// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package ghttp

import (
	"context"
	"reflect"

	"github.com/gogf/gf/v2/internal/intlog"
	"github.com/gogf/gf/v2/net/goai"
	"github.com/gogf/gf/v2/text/gstr"
)

// initOpenApi generates api specification using OpenApiV3 protocol.
func (s *Server) initOpenApi() {
	if s.config.OpenApiPath == "" {
		return
	}
	var (
		ctx    = context.TODO()
		err    error
		method string
	)
	for _, item := range s.GetRoutes() {
		switch item.Type {
		case HandlerTypeMiddleware, HandlerTypeHook:
			continue
		}
		method = item.Method
		if gstr.Equal(method, defaultMethod) {
			method = ""
		}
		if item.Handler.Info.Func == nil {
			oaiInput := goai.AddInput{
				Path:   item.Route,
				Method: method,
				Object: item.Handler.Info.Value.Interface(),
				Func:   gstr.SubStrFromREx(item.Handler.Name, "."),
			}
			if item.Handler.Info.Ctrl.IsValid() {
				if item.Handler.Info.Ctrl.Kind() == reflect.Ptr {
					oaiInput.Ctrl = item.Handler.Info.Ctrl.Type().Elem().Name()
				} else {
					oaiInput.Ctrl = item.Handler.Info.Ctrl.Type().Name()
				}
			}
			if err = s.openapi.Add(oaiInput); err != nil {
				s.Logger().Fatalf(ctx, `%+v`, err)
			}
		}
	}
}

// openapiSpec is a build-in handler automatic producing for openapi specification json file.
func (s *Server) openapiSpec(r *Request) {
	var (
		err error
	)
	if s.config.OpenApiPath == "" {
		r.Response.Write(`OpenApi specification file producing is disabled`)
	} else {
		err = r.Response.WriteJson(s.openapi)
	}

	if err != nil {
		intlog.Errorf(r.Context(), `%+v`, err)
	}
}
