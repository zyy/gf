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
	"github.com/gogf/gf/v2/internal/json"
	"github.com/gogf/gf/v2/net/goai"
	"github.com/gogf/gf/v2/test/gtest"
	"github.com/gogf/gf/v2/util/gmeta"
)

func TestOpenApiV3_XExtensions_Json(t *testing.T) {
	type CommonResponse struct {
		Code    int         `json:"code"    description:"Error code"`
		Message string      `json:"message" description:"Error message"`
		Data    interface{} `json:"data"    description:"Result data for certain request according API definition"`
	}
	type ProductSearchReq struct {
		gmeta.Meta `path:"/test" method:"get" x-k1:"v1"`
	}
	type ProductSearchRes struct {
		g.Meta `x-k2:"v2"`
	}

	f := func(ctx context.Context, req *ProductSearchReq) (res *ProductSearchRes, err error) {
		return
	}

	gtest.C(t, func(t *gtest.T) {
		var (
			err  error
			oai1 = goai.New()
			oai2 = goai.New()
		)

		oai1.Config.CommonResponse = CommonResponse{}

		err = oai1.Add(goai.AddInput{
			Object: f,
		})
		t.AssertNil(err)

		oai1Json, err := json.Marshal(oai1)
		t.AssertNil(err)

		err = json.Unmarshal(oai1Json, oai2)
		t.AssertNil(err)

		oai2Json, err := json.Marshal(oai2)
		t.AssertNil(err)

		t.Assert(oai1Json, oai2Json)
	})
}
