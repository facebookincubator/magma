// Copyright (c) 2004-present Facebook All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package viewer

import (
	"context"
	"net/http"
	"strconv"

	"github.com/facebookincubator/symphony/graph/ent"
	"github.com/facebookincubator/symphony/pkg/log"

	"go.opencensus.io/tag"
	"go.opencensus.io/trace"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	// TenantHeader is the http tenant header.
	TenantHeader = "x-auth-organization"
	// UserHeader is the http user header.
	UserHeader = "x-auth-user-email"
	// ReadOnlyHeader is the http readonly permission header.
	ReadOnlyHeader = "x-auth-user-readonly"
)

// Attributes recorded on the span of the requests.
const (
	TenantAttribute    = "viewer.tenant"
	UserAttribute      = "viewer.user"
	RoleAttribute      = "viewer.role"
	UserAgentAttribute = "viewer.user_agent"
)

// The following tags are applied to context recorded by this package.
var (
	KeyTenant    = tag.MustNewKey(TenantAttribute)
	KeyUser      = tag.MustNewKey(UserAttribute)
	KeyRole      = tag.MustNewKey(RoleAttribute)
	KeyUserAgent = tag.MustNewKey(UserAgentAttribute)
)

// Viewer holds additional per request information.
type Viewer struct {
	Tenant string
	User   string
	Role   string
}

// MarshalLogObject implements zapcore.ObjectMarshaler interface.
func (v *Viewer) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	enc.AddString("tenant", v.Tenant)
	enc.AddString("user", v.User)
	enc.AddString("role", v.Role)
	return nil
}

func (v *Viewer) traceAttrs() []trace.Attribute {
	return []trace.Attribute{
		trace.StringAttribute(TenantAttribute, v.Tenant),
		trace.StringAttribute(UserAttribute, v.User),
		trace.StringAttribute(RoleAttribute, v.Role),
	}
}

func (v *Viewer) tags(r *http.Request) []tag.Mutator {
	return []tag.Mutator{
		tag.Upsert(KeyTenant, v.Tenant),
		tag.Upsert(KeyUser, v.User),
		tag.Upsert(KeyRole, v.Role),
		tag.Upsert(KeyUserAgent, r.UserAgent()),
	}
}

// TenancyHandler adds viewer / tenancy into incoming requests.
func TenancyHandler(h http.Handler, tenancy Tenancy) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tenant := r.Header.Get(TenantHeader)
		if tenant == "" {
			http.Error(w, "missing tenant header", http.StatusBadRequest)
			return
		}

		v := &Viewer{Tenant: tenant, User: r.Header.Get(UserHeader)}
		if ro, err := strconv.ParseBool(r.Header.Get(ReadOnlyHeader)); ro && err == nil {
			v.Role = "readonly"
		}

		ctx := log.NewFieldsContext(r.Context(), zap.Object("viewer", v))
		trace.FromContext(ctx).AddAttributes(v.traceAttrs()...)
		ctx, _ = tag.New(ctx, v.tags(r)...)

		ctx = NewContext(ctx, v)
		if tenancy != nil {
			client, err := tenancy.ClientFor(ctx, tenant)
			if err != nil {
				http.Error(w, "getting tenancy client", http.StatusServiceUnavailable)
				return
			}
			if v.Role == "readonly" {
				client = client.ReadOnly()
			}
			ctx = ent.NewContext(ctx, client)
		}
		h.ServeHTTP(w, r.WithContext(ctx))
	})
}

type contextKey struct{}

// FromContext returns the Viewer stored in a context, or nil if there isn't one.
func FromContext(ctx context.Context) *Viewer {
	v, _ := ctx.Value(contextKey{}).(*Viewer)
	return v
}

// NewContext returns a new context with the given Viewer attached.
func NewContext(parent context.Context, v *Viewer) context.Context {
	return context.WithValue(parent, contextKey{}, v)
}
