// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package goai_test

import (
	"context"
	"testing"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/goai"
	"github.com/gogf/gf/v2/test/gtest"
	"github.com/gogf/gf/v2/util/gmeta"
)

func TestOpenApiV3_XExtensions(t *testing.T) {
	type CommonResponse struct {
		Code    int         `json:"code"    description:"Error code"`
		Message string      `json:"message" description:"Error message"`
		Data    interface{} `json:"data"    description:"Result data for certain request according API definition"`
	}
	type ProductSearchReq struct {
		gmeta.Meta `path:"/test" method:"get" x-k1:"v1" x-k2:"v2"`
	}
	type ProductSearchRes struct {
		g.Meta `x-k3:"v3" x-k4:"v4"`
	}

	f := func(ctx context.Context, req *ProductSearchReq) (res *ProductSearchRes, err error) {
		return
	}

	gtest.C(t, func(t *gtest.T) {
		var (
			err error
			oai = goai.New()
		)

		oai.Config.CommonResponse = CommonResponse{}

		err = oai.Add(goai.AddInput{
			Object: f,
		})
		t.AssertNil(err)
		// Schema asserts.
		// fmt.Println(oai.String())
		t.Assert(len(oai.Components.Schemas.Map()), 3)
		t.Assert(len(oai.Paths), 1)
		t.Assert(oai.Paths["/test"].Get.XExtensions["x-k1"], "v1")
		t.Assert(oai.Paths["/test"].Get.XExtensions["x-k2"], "v2")
		t.Assert(oai.Paths["/test"].Get.Responses["200"].Value.XExtensions["x-k3"], "v3")
		t.Assert(oai.Paths["/test"].Get.Responses["200"].Value.XExtensions["x-k4"], "v4")
	})
}
