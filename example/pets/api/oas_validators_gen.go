// Code generated by ogen, DO NOT EDIT.

package api

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"math"
	"net"
	"net/http"
	"net/url"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/go-faster/errors"
	"github.com/go-faster/jx"
	"github.com/google/uuid"
	"github.com/ogen-go/ogen/conv"
	ht "github.com/ogen-go/ogen/http"
	"github.com/ogen-go/ogen/json"
	"github.com/ogen-go/ogen/otelogen"
	"github.com/ogen-go/ogen/uri"
	"github.com/ogen-go/ogen/validate"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/trace"
)

// No-op definition for keeping imports.
var (
	_ = context.Background()
	_ = fmt.Stringer(nil)
	_ = strings.Builder{}
	_ = errors.Is
	_ = sort.Ints
	_ = http.MethodGet
	_ = io.Copy
	_ = json.Marshal
	_ = bytes.NewReader
	_ = strconv.ParseInt
	_ = time.Time{}
	_ = conv.ToInt32
	_ = uuid.UUID{}
	_ = uri.PathEncoder{}
	_ = url.URL{}
	_ = math.Mod
	_ = validate.Int{}
	_ = ht.NewRequest
	_ = net.IP{}
	_ = otelogen.Version
	_ = trace.TraceIDFromHex
	_ = otel.GetTracerProvider
	_ = metric.NewNoopMeterProvider
	_ = regexp.MustCompile
	_ = jx.Null
	_ = sync.Pool{}
)

func (s ListCategoryOKApplicationJSON) Validate() error {
	if s == nil {
		return errors.New("nil is invalid value")
	}
	return nil
}
func (s ListCategoryPetsOKApplicationJSON) Validate() error {
	if s == nil {
		return errors.New("nil is invalid value")
	}
	return nil
}
func (s ListPetCategoriesOKApplicationJSON) Validate() error {
	if s == nil {
		return errors.New("nil is invalid value")
	}
	return nil
}
func (s ListPetFriendsOKApplicationJSON) Validate() error {
	if s == nil {
		return errors.New("nil is invalid value")
	}
	return nil
}
func (s ListPetOKApplicationJSON) Validate() error {
	if s == nil {
		return errors.New("nil is invalid value")
	}
	return nil
}
func (s ListUserOKApplicationJSON) Validate() error {
	if s == nil {
		return errors.New("nil is invalid value")
	}
	return nil
}
func (s ListUserPetsOKApplicationJSON) Validate() error {
	if s == nil {
		return errors.New("nil is invalid value")
	}
	return nil
}