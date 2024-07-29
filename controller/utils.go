package controller

// import (
// 	"context"
// 	"encoding/json"
// 	"errors"
// 	"fmt"
// 	"net/http"

// 	"github.com/sirupsen/logrus"
// )

// func RespondWithError(w http.ResponseWriter, r *http.Request, defaultCode int, err error) {
// 	var payload interface{}
// 	code := defaultCode
// 	action := getActionNameFromCtx(r.Context())

// 	logger := logctx.From(r.Context()).
// 		WithField("log_type", "request_err").
// 		WithField("action", fmt.Sprintf("[%s]", action))

// 	var reqError common.RequestError
// 	if errors.As(err, &reqError) {
// 		code = reqError.StatusCode
// 		payload = reqError.GetPayload()
// 		logger = logger.WithField("err_payload", payload)
// 	} else {
// 		reqId := common.GetRequestIdFromContext(r.Context())
// 		if code == 500 {
// 			payload = map[string]interface{}{"message": "Internal Server Error", "correlation_id": reqId}
// 		} else {
// 			payload = map[string]interface{}{"message": err.Error(), "correlation_id": reqId}
// 		}
// 		logger = logger.WithError(err)
// 	}

// 	logLevel := logrus.InfoLevel
// 	if code >= 500 {
// 		logLevel = logrus.ErrorLevel
// 	}
// 	logger.WithField("status", code).Logf(logLevel, "%s %s", r.Method, r.URL.String())

// 	RespondWithJSON(w, code, payload)
// }

// func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
// 	t := "data"
// 	if code > 300 {
// 		t = "errors"
// 	}
// 	response, _ := json.Marshal(map[string]interface{}{
// 		t: payload,
// 	})
// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(code)
// 	_, _ = w.Write(response)
// }

// type actionNameKey struct{}

// func getActionNameFromCtx(ctx context.Context) string {
// 	action, ok := ctx.Value(actionNameKey{}).(string)
// 	if ok {
// 		return action
// 	}
// 	return "UNKNOWN"
// }
