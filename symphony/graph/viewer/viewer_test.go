// Copyright (c) 2004-present Facebook All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package viewer

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/facebookincubator/symphony/cloud/log"
	"github.com/facebookincubator/symphony/graph/ent"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"go.opencensus.io/trace"
	"go.uber.org/zap"
	"go.uber.org/zap/zaptest/observer"
)

func TestViewerHandler(t *testing.T) {
	tests := []struct {
		name    string
		prepare func(*http.Request)
		expect  func(*testing.T, *httptest.ResponseRecorder)
	}{
		{
			name: "TestTenant",
			prepare: func(req *http.Request) {
				req.Header.Set(TenantHeader, "test")
			},
			expect: func(t *testing.T, rec *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusOK, rec.Code)
				assert.Equal(t, "test", rec.Body.String())
			},
		},
		{
			name: "NoTenant",
			expect: func(t *testing.T, rec *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusBadRequest, rec.Code)
				assert.NotZero(t, rec.Body.Len())
			},
		},
	}

	h := TenancyHandler(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			viewer := FromContext(r.Context())
			require.NotNil(t, viewer)
			assert.NotNil(t, log.FieldsFromContext(r.Context()))
			_, _ = io.WriteString(w, viewer.Tenant)
		}),
		nil,
	)
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "/", nil)
			if tc.prepare != nil {
				tc.prepare(req)
			}
			rec := httptest.NewRecorder()
			h.ServeHTTP(rec, req)
			tc.expect(t, rec)
		})
	}
}

func TestViewerMarshalLog(t *testing.T) {
	core, o := observer.New(zap.InfoLevel)
	logger := zap.New(core)
	v := &Viewer{Tenant: "test", User: "tester"}
	logger.Info("viewer log test", zap.Object("viewer", v))

	logs := o.TakeAll()
	require.Len(t, logs, 1)
	field, ok := logs[0].ContextMap()["viewer"].(map[string]interface{})
	require.True(t, ok)
	assert.Equal(t, v.Tenant, field["tenant"])
	assert.Equal(t, v.User, field["user"])
}

type testExporter struct {
	mock.Mock
}

func (e *testExporter) ExportSpan(s *trace.SpanData) {
	e.Called(s)
}

func TestViewerSpanAttributes(t *testing.T) {
	h := TenancyHandler(
		http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
			w.WriteHeader(http.StatusAccepted)
		}),
		nil,
	)
	t.Run("WithSpan", func(t *testing.T) {
		var e testExporter
		trace.RegisterExporter(&e)
		defer trace.UnregisterExporter(&e)

		e.On("ExportSpan", mock.AnythingOfType("*trace.SpanData")).
			Run(func(args mock.Arguments) {
				s := args.Get(0).(*trace.SpanData)
				assert.Equal(t, "test", s.Attributes["viewer.tenant"])
				assert.Equal(t, "test", s.Attributes["viewer.user"])
			}).
			Once()
		defer e.AssertExpectations(t)

		ctx, span := trace.StartSpan(context.Background(), "test",
			trace.WithSampler(trace.AlwaysSample()),
		)
		defer span.End()

		req := httptest.NewRequest(http.MethodGet, "/", nil).WithContext(ctx)
		req.Header.Set(TenantHeader, "test")
		req.Header.Set(UserHeader, "test")
		rec := httptest.NewRecorder()
		h.ServeHTTP(rec, req)
		assert.Equal(t, http.StatusAccepted, rec.Code)
	})
	t.Run("WithoutSpan", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(TenantHeader, "test")
		rec := httptest.NewRecorder()
		assert.NotPanics(t, func() { h.ServeHTTP(rec, req) })
		assert.Equal(t, http.StatusAccepted, rec.Code)
	})
}

func TestViewerTenancy(t *testing.T) {
	t.Run("WithTenancy", func(t *testing.T) {
		var client ent.Client
		h := TenancyHandler(
			http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				assert.True(t, &client == ent.FromContext(r.Context()))
				w.WriteHeader(http.StatusAccepted)
			}),
			NewFixedTenancy(&client),
		)
		req := httptest.NewRequest(http.MethodPost, "/", nil)
		req.Header.Set(TenantHeader, "test")
		rec := httptest.NewRecorder()
		h.ServeHTTP(rec, req)
		assert.Equal(t, http.StatusAccepted, rec.Code)
	})
	t.Run("WithoutTenancy", func(t *testing.T) {
		h := TenancyHandler(
			http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				assert.Nil(t, ent.FromContext(r.Context()))
				w.WriteHeader(http.StatusAccepted)
			}),
			nil,
		)
		req := httptest.NewRequest(http.MethodPost, "/", nil)
		req.Header.Set(TenantHeader, "test")
		rec := httptest.NewRecorder()
		h.ServeHTTP(rec, req)
		assert.Equal(t, http.StatusAccepted, rec.Code)
	})
}
