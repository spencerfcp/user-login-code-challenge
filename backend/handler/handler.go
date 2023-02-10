package handler

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/spencerfcp/scoirapi/v2/env"
	"github.com/spencerfcp/scoirapi/v2/httperr"
	"github.com/spencerfcp/scoirapi/v2/logerr"
	"github.com/spencerfcp/scoirapi/v2/pb"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

type Methods struct {
	Get    http.HandlerFunc
	Post   http.HandlerFunc
	Patch  http.HandlerFunc
	Delete http.HandlerFunc
}

const (
	ContentTypeJSON        = "application/json"
	ContentTypeFormEncoded = "application/x-www-form-urlencoded"
	ContentTypeOctetStream = "application/octet-stream"
)

const (
	ContentType   = "Content-Type"
	Authorization = "Authorization"
)

func Path(mux *http.ServeMux, basePath string, methods Methods) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		switch {
		case r.Method == http.MethodGet && methods.Get != nil:
			methods.Get(w, r)
		case r.Method == http.MethodPost && methods.Post != nil:
			methods.Post(w, r)
		case r.Method == http.MethodPatch && methods.Patch != nil:
			methods.Patch(w, r)
		case r.Method == http.MethodDelete && methods.Delete != nil:
			methods.Delete(w, r)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}

	}

	mux.HandleFunc(basePath, handler)
}

func writeErrorToResponseAndLog(w http.ResponseWriter, err error) {
	switch cause := errors.Cause(err).(type) {
	case httperr.HTTPError:
		w.WriteHeader(cause.StatusCode)
	case httperr.HumanReadableErr:
		w.WriteHeader(http.StatusBadRequest)
		_, err := w.Write(([]byte)(cause.Error()))
		if err != nil {
			logerr.FromError(errors.Wrap(err, "write humanReadableErr failed"))
		}
	default:
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func writePanicToResponseAndLog(w http.ResponseWriter, errInterface interface{}) {
	logerr.FromHTTPRequestPanic(errInterface)
	w.WriteHeader(http.StatusInternalServerError)
}

type UnauthenticatedParams struct {
	Message       proto.Message
	DB            *sqlx.DB
	RequestHeader http.Header
	Request       *http.Request
	Env           env.Env
}

func loggedProtoWithRequest(db *sqlx.DB, factory func() proto.Message, fn func(*http.Request, proto.Message, time.Time) (proto.Message, error)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		now := time.Now()
		defer func() {
			if err := recover(); err != nil {
				writePanicToResponseAndLog(w, err)
			}
		}()
		log.Printf("%s method=%s path=%s\n", r.RemoteAddr, r.Method, r.URL)
		var reader io.Reader
		if r.Method == http.MethodGet {
			queryStr, err := url.QueryUnescape(r.URL.RawQuery)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			reader = strings.NewReader(queryStr)
		} else {
			reader = r.Body
		}

		message := factory()
		_, isEmpty := message.(*pb.EmptyRequest)
		if !isEmpty {
			var rawJson json.RawMessage
			if err := json.NewDecoder(reader).Decode(&rawJson); err != nil {
				if errors.Cause(err) == io.ErrUnexpectedEOF || errors.Cause(err) == io.EOF {
					w.WriteHeader(http.StatusBadRequest)
					return
				}

				panic(errors.Wrapf(err, "decode json failed from %s", r.Method))
			}

			unmarshaller := protojson.UnmarshalOptions{DiscardUnknown: true}
			err := unmarshaller.Unmarshal(rawJson, message)
			if err != nil {
				panic(errors.Wrapf(err, "unmarshal json failed from %s", r.Method))
			}
		}

		o, handlerError := fn(r, message, now)
		writeProtoHandlerResult(w, o, handlerError)
	})
}

func writeProtoHandlerResult(w http.ResponseWriter, o proto.Message, err error) {
	w.Header().Set(ContentType, ContentTypeJSON)

	if err != nil {
		writeErrorToResponseAndLog(w, err)
		return
	}

	if o != nil {
		m, err := protojson.MarshalOptions{UseProtoNames: true}.Marshal(o)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if _, err := w.Write(m); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
}

func writeJSONHandlerResult(w http.ResponseWriter, o interface{}, err error) {
	w.Header().Set(ContentType, ContentTypeJSON)
	if err != nil {
		writeErrorToResponseAndLog(w, err)
		return
	}

	json, err := json.Marshal(o)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	_, err = w.Write(json)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

// Process unauthenticated routes using protobuff
func UnauthenticatedProto(
	factory func() proto.Message,
	fn func(UnauthenticatedParams) (proto.Message, error),
	db *sqlx.DB,
	env env.Env,
) http.HandlerFunc {
	return loggedProtoWithRequest(db, factory, func(r *http.Request, m proto.Message, now time.Time) (proto.Message, error) {
		return fn(UnauthenticatedParams{
			Message:       m,
			DB:            db,
			RequestHeader: r.Header,
			Request:       r,
			Env:           env,
		})
	}).ServeHTTP
}
