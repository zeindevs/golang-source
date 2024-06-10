package api

import "net/http"

func jsonResponseWithMetadata(w http.ResponseWriter, status int, data any, metadata any) error {
	type envelope struct {
		Data     any `json:"data"`
		Metadata any `json:"metadata"`
	}

	return writeJSON(w, status, &envelope{Data: data, Metadata: metadata})

}

func rateLimiterExceededResponse(w http.ResponseWriter, r *http.Request, retryAfter string) {
	// app.logger.Warnw("rate limiter exceeded", "method", r.Method, "path", r.URL.Path)
	w.Header().Set("Retry-After", retryAfter)
	writeJSONError(w, http.StatusTooManyRequests, "rate limite exceeded, rety after: "+retryAfter)
}
