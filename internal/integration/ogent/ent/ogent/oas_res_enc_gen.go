// Code generated by ogen, DO NOT EDIT.

package ogent

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"math"
	"math/big"
	"math/bits"
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
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
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
	_ = bits.LeadingZeros64
	_ = big.Rat{}
	_ = validate.Int{}
	_ = ht.NewRequest
	_ = net.IP{}
	_ = otelogen.Version
	_ = attribute.KeyValue{}
	_ = trace.TraceIDFromHex
	_ = otel.GetTracerProvider
	_ = metric.NewNoopMeterProvider
	_ = regexp.MustCompile
	_ = jx.Null
	_ = sync.Pool{}
	_ = codes.Unset
)

func encodeCreateCategoryResponse(response CreateCategoryRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *CategoryCreate:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		e := jx.GetWriter()
		defer jx.PutWriter(e)

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil
	case *R400:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		e := jx.GetWriter()
		defer jx.PutWriter(e)

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil
	case *R409:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(409)
		e := jx.GetWriter()
		defer jx.PutWriter(e)

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil
	case *R500:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500)
		e := jx.GetWriter()
		defer jx.PutWriter(e)

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil
	default:
		return errors.Errorf("/categories"+`: unexpected response type: %T`, response)
	}
}

func encodeCreatePetResponse(response CreatePetRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *PetCreate:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		e := jx.GetWriter()
		defer jx.PutWriter(e)

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil
	case *R400:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		e := jx.GetWriter()
		defer jx.PutWriter(e)

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil
	case *R409:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(409)
		e := jx.GetWriter()
		defer jx.PutWriter(e)

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil
	case *R500:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500)
		e := jx.GetWriter()
		defer jx.PutWriter(e)

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil
	default:
		return errors.Errorf("/pets"+`: unexpected response type: %T`, response)
	}
}

func encodeCreateUserResponse(response CreateUserRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *UserCreate:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		e := jx.GetWriter()
		defer jx.PutWriter(e)

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil
	case *R400:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		e := jx.GetWriter()
		defer jx.PutWriter(e)

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil
	case *R409:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(409)
		e := jx.GetWriter()
		defer jx.PutWriter(e)

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil
	case *R500:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500)
		e := jx.GetWriter()
		defer jx.PutWriter(e)

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil
	default:
		return errors.Errorf("/users"+`: unexpected response type: %T`, response)
	}
}

func encodeDeleteCategoryResponse(response DeleteCategoryRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *DeleteCategoryNoContent:
		w.WriteHeader(204)
		return nil
	case *R400:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		e := jx.GetWriter()
		defer jx.PutWriter(e)

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil
	case *R404:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(404)
		e := jx.GetWriter()
		defer jx.PutWriter(e)

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil
	case *R409:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(409)
		e := jx.GetWriter()
		defer jx.PutWriter(e)

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil
	case *R500:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500)
		e := jx.GetWriter()
		defer jx.PutWriter(e)

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil
	default:
		return errors.Errorf("/categories/{id}"+`: unexpected response type: %T`, response)
	}
}

func encodeDeletePetResponse(response DeletePetRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *DeletePetNoContent:
		w.WriteHeader(204)
		return nil
	case *R400:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		e := jx.GetWriter()
		defer jx.PutWriter(e)

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil
	case *R404:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(404)
		e := jx.GetWriter()
		defer jx.PutWriter(e)

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil
	case *R409:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(409)
		e := jx.GetWriter()
		defer jx.PutWriter(e)

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil
	case *R500:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500)
		e := jx.GetWriter()
		defer jx.PutWriter(e)

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil
	default:
		return errors.Errorf("/pets/{id}"+`: unexpected response type: %T`, response)
	}
}

func encodeDeleteUserResponse(response DeleteUserRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *DeleteUserNoContent:
		w.WriteHeader(204)
		return nil
	case *R400:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		e := jx.GetWriter()
		defer jx.PutWriter(e)

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil
	case *R404:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(404)
		e := jx.GetWriter()
		defer jx.PutWriter(e)

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil
	case *R409:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(409)
		e := jx.GetWriter()
		defer jx.PutWriter(e)

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil
	case *R500:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500)
		e := jx.GetWriter()
		defer jx.PutWriter(e)

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil
	default:
		return errors.Errorf("/users/{id}"+`: unexpected response type: %T`, response)
	}
}

func encodeListCategoryResponse(response ListCategoryRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *ListCategoryOKApplicationJSON:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		e := jx.GetWriter()
		defer jx.PutWriter(e)

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil
	case *R400:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		e := jx.GetWriter()
		defer jx.PutWriter(e)

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil
	case *R404:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(404)
		e := jx.GetWriter()
		defer jx.PutWriter(e)

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil
	case *R409:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(409)
		e := jx.GetWriter()
		defer jx.PutWriter(e)

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil
	case *R500:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500)
		e := jx.GetWriter()
		defer jx.PutWriter(e)

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil
	default:
		return errors.Errorf("/categories"+`: unexpected response type: %T`, response)
	}
}

func encodeListCategoryPetsResponse(response ListCategoryPetsRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *ListCategoryPetsOKApplicationJSON:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		e := jx.GetWriter()
		defer jx.PutWriter(e)

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil
	case *R400:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		e := jx.GetWriter()
		defer jx.PutWriter(e)

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil
	case *R404:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(404)
		e := jx.GetWriter()
		defer jx.PutWriter(e)

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil
	case *R409:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(409)
		e := jx.GetWriter()
		defer jx.PutWriter(e)

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil
	case *R500:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500)
		e := jx.GetWriter()
		defer jx.PutWriter(e)

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil
	default:
		return errors.Errorf("/categories/{id}/pets"+`: unexpected response type: %T`, response)
	}
}

func encodeListPetResponse(response ListPetRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *ListPetOKApplicationJSON:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		e := jx.GetWriter()
		defer jx.PutWriter(e)

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil
	case *R400:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		e := jx.GetWriter()
		defer jx.PutWriter(e)

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil
	case *R404:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(404)
		e := jx.GetWriter()
		defer jx.PutWriter(e)

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil
	case *R409:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(409)
		e := jx.GetWriter()
		defer jx.PutWriter(e)

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil
	case *R500:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500)
		e := jx.GetWriter()
		defer jx.PutWriter(e)

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil
	default:
		return errors.Errorf("/pets"+`: unexpected response type: %T`, response)
	}
}

func encodeListPetCategoriesResponse(response ListPetCategoriesRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *ListPetCategoriesOKApplicationJSON:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		e := jx.GetWriter()
		defer jx.PutWriter(e)

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil
	case *R400:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		e := jx.GetWriter()
		defer jx.PutWriter(e)

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil
	case *R404:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(404)
		e := jx.GetWriter()
		defer jx.PutWriter(e)

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil
	case *R409:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(409)
		e := jx.GetWriter()
		defer jx.PutWriter(e)

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil
	case *R500:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500)
		e := jx.GetWriter()
		defer jx.PutWriter(e)

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil
	default:
		return errors.Errorf("/pets/{id}/categories"+`: unexpected response type: %T`, response)
	}
}

func encodeListPetFriendsResponse(response ListPetFriendsRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *ListPetFriendsOKApplicationJSON:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		e := jx.GetWriter()
		defer jx.PutWriter(e)

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil
	case *R400:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		e := jx.GetWriter()
		defer jx.PutWriter(e)

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil
	case *R404:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(404)
		e := jx.GetWriter()
		defer jx.PutWriter(e)

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil
	case *R409:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(409)
		e := jx.GetWriter()
		defer jx.PutWriter(e)

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil
	case *R500:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500)
		e := jx.GetWriter()
		defer jx.PutWriter(e)

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil
	default:
		return errors.Errorf("/pets/{id}/friends"+`: unexpected response type: %T`, response)
	}
}

func encodeListUserResponse(response ListUserRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *ListUserOKApplicationJSON:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		e := jx.GetWriter()
		defer jx.PutWriter(e)

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil
	case *R400:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		e := jx.GetWriter()
		defer jx.PutWriter(e)

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil
	case *R404:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(404)
		e := jx.GetWriter()
		defer jx.PutWriter(e)

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil
	case *R409:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(409)
		e := jx.GetWriter()
		defer jx.PutWriter(e)

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil
	case *R500:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500)
		e := jx.GetWriter()
		defer jx.PutWriter(e)

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil
	default:
		return errors.Errorf("/users"+`: unexpected response type: %T`, response)
	}
}

func encodeListUserPetsResponse(response ListUserPetsRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *ListUserPetsOKApplicationJSON:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		e := jx.GetWriter()
		defer jx.PutWriter(e)

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil
	case *R400:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		e := jx.GetWriter()
		defer jx.PutWriter(e)

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil
	case *R404:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(404)
		e := jx.GetWriter()
		defer jx.PutWriter(e)

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil
	case *R409:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(409)
		e := jx.GetWriter()
		defer jx.PutWriter(e)

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil
	case *R500:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500)
		e := jx.GetWriter()
		defer jx.PutWriter(e)

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil
	default:
		return errors.Errorf("/users/{id}/pets"+`: unexpected response type: %T`, response)
	}
}

func encodeReadCategoryResponse(response ReadCategoryRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *CategoryRead:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		e := jx.GetWriter()
		defer jx.PutWriter(e)

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil
	case *R400:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		e := jx.GetWriter()
		defer jx.PutWriter(e)

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil
	case *R404:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(404)
		e := jx.GetWriter()
		defer jx.PutWriter(e)

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil
	case *R409:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(409)
		e := jx.GetWriter()
		defer jx.PutWriter(e)

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil
	case *R500:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500)
		e := jx.GetWriter()
		defer jx.PutWriter(e)

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil
	default:
		return errors.Errorf("/categories/{id}"+`: unexpected response type: %T`, response)
	}
}

func encodeReadPetResponse(response ReadPetRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *PetRead:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		e := jx.GetWriter()
		defer jx.PutWriter(e)

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil
	case *R400:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		e := jx.GetWriter()
		defer jx.PutWriter(e)

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil
	case *R404:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(404)
		e := jx.GetWriter()
		defer jx.PutWriter(e)

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil
	case *R409:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(409)
		e := jx.GetWriter()
		defer jx.PutWriter(e)

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil
	case *R500:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500)
		e := jx.GetWriter()
		defer jx.PutWriter(e)

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil
	default:
		return errors.Errorf("/pets/{id}"+`: unexpected response type: %T`, response)
	}
}

func encodeReadPetOwnerResponse(response ReadPetOwnerRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *PetOwnerRead:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		e := jx.GetWriter()
		defer jx.PutWriter(e)

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil
	case *R400:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		e := jx.GetWriter()
		defer jx.PutWriter(e)

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil
	case *R404:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(404)
		e := jx.GetWriter()
		defer jx.PutWriter(e)

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil
	case *R409:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(409)
		e := jx.GetWriter()
		defer jx.PutWriter(e)

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil
	case *R500:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500)
		e := jx.GetWriter()
		defer jx.PutWriter(e)

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil
	default:
		return errors.Errorf("/pets/{id}/owner"+`: unexpected response type: %T`, response)
	}
}

func encodeReadUserResponse(response ReadUserRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *UserRead:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		e := jx.GetWriter()
		defer jx.PutWriter(e)

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil
	case *R400:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		e := jx.GetWriter()
		defer jx.PutWriter(e)

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil
	case *R404:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(404)
		e := jx.GetWriter()
		defer jx.PutWriter(e)

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil
	case *R409:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(409)
		e := jx.GetWriter()
		defer jx.PutWriter(e)

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil
	case *R500:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500)
		e := jx.GetWriter()
		defer jx.PutWriter(e)

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil
	default:
		return errors.Errorf("/users/{id}"+`: unexpected response type: %T`, response)
	}
}

func encodeReadUserBestFriendResponse(response ReadUserBestFriendRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *UserBestFriendRead:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		e := jx.GetWriter()
		defer jx.PutWriter(e)

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil
	case *R400:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		e := jx.GetWriter()
		defer jx.PutWriter(e)

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil
	case *R404:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(404)
		e := jx.GetWriter()
		defer jx.PutWriter(e)

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil
	case *R409:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(409)
		e := jx.GetWriter()
		defer jx.PutWriter(e)

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil
	case *R500:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500)
		e := jx.GetWriter()
		defer jx.PutWriter(e)

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil
	default:
		return errors.Errorf("/users/{id}/best-friend"+`: unexpected response type: %T`, response)
	}
}

func encodeUpdateCategoryResponse(response UpdateCategoryRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *CategoryUpdate:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		e := jx.GetWriter()
		defer jx.PutWriter(e)

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil
	case *R400:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		e := jx.GetWriter()
		defer jx.PutWriter(e)

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil
	case *R404:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(404)
		e := jx.GetWriter()
		defer jx.PutWriter(e)

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil
	case *R409:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(409)
		e := jx.GetWriter()
		defer jx.PutWriter(e)

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil
	case *R500:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500)
		e := jx.GetWriter()
		defer jx.PutWriter(e)

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil
	default:
		return errors.Errorf("/categories/{id}"+`: unexpected response type: %T`, response)
	}
}

func encodeUpdatePetResponse(response UpdatePetRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *PetUpdate:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		e := jx.GetWriter()
		defer jx.PutWriter(e)

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil
	case *R400:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		e := jx.GetWriter()
		defer jx.PutWriter(e)

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil
	case *R404:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(404)
		e := jx.GetWriter()
		defer jx.PutWriter(e)

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil
	case *R409:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(409)
		e := jx.GetWriter()
		defer jx.PutWriter(e)

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil
	case *R500:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500)
		e := jx.GetWriter()
		defer jx.PutWriter(e)

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil
	default:
		return errors.Errorf("/pets/{id}"+`: unexpected response type: %T`, response)
	}
}

func encodeUpdateUserResponse(response UpdateUserRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *UserUpdate:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		e := jx.GetWriter()
		defer jx.PutWriter(e)

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil
	case *R400:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		e := jx.GetWriter()
		defer jx.PutWriter(e)

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil
	case *R404:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(404)
		e := jx.GetWriter()
		defer jx.PutWriter(e)

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil
	case *R409:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(409)
		e := jx.GetWriter()
		defer jx.PutWriter(e)

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil
	case *R500:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500)
		e := jx.GetWriter()
		defer jx.PutWriter(e)

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil
	default:
		return errors.Errorf("/users/{id}"+`: unexpected response type: %T`, response)
	}
}
