package main

import (
	"GoTravis/api"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"testing"
)

/**
 * :=  created by:  Shuza
 * :=  create date:  11-Apr-2019
 * :=  (C) CopyRight Shuza
 * :=  www.shuza.ninja
 * :=  shuza.sa@gmail.com
 * :=  Fun  :  Coffee  :  Code
 **/

func TestIndexHandler(t *testing.T) {
	req := httptest.NewRequest("GET", "/", nil)
	resp := httptest.NewRecorder()

	api.GetRoutes().ServeHTTP(resp, req)
	assert.Equal(t, 200, resp.Code)
}
