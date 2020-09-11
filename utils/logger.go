package utils

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
)

var Log *log.Logger

func init() {
	os.Setenv("TZ", "Asia/Kolkata")
	Log = log.New()
	Log.SetLevel(log.TraceLevel)
	Log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})
}

type responseLogger struct {
	rw     http.ResponseWriter
	start  time.Time
	status int
	size   int
}

func (rl *responseLogger) Header() http.Header {
	return rl.rw.Header()
}

func (rl *responseLogger) Write(bytes []byte) (int, error) {
	if rl.status == 0 {
		rl.status = http.StatusOK
	}

	size, err := rl.rw.Write(bytes)

	rl.size += size

	return size, err
}

func (rl *responseLogger) WriteHeader(status int) {
	rl.status = status

	rl.rw.WriteHeader(status)
}

func (rl *responseLogger) Flush() {
	f, ok := rl.rw.(http.Flusher)

	if ok {
		f.Flush()
	}
}

//LogRequestHandler ...
type LogRequestHandler struct {
	HTTPHandler http.Handler
	Log         *log.Logger
}

//New ...
func New(h http.Handler, log *log.Logger) http.Handler {
	return LogRequestHandler{
		HTTPHandler: h,
		Log:         log,
	}
}

func (rh LogRequestHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	rl := &responseLogger{rw: res, start: time.Now()}

	rh.HTTPHandler.ServeHTTP(rl, req)

	rh.write(rl, req)
}

func (rh LogRequestHandler) write(rl *responseLogger, r *http.Request) {
	rh.Log.WithFields(logrus.Fields{
		"uri":         r.URL.String(),
		"method":      r.Method,
		"remoteAddr":  r.RemoteAddr,
		"porto":       r.Proto,
		"status":      rl.status,
		"size":        rl.size,
		"referer":     r.Referer(),
		"userAgent":   r.UserAgent(),
		"reponseTime": parseResponseTime(rl.start),
		"requestTime": rl.start.UTC(),
	}).Info("Processing Request")
}

func parseResponseTime(start time.Time) string {
	return fmt.Sprintf("%.4f ms", time.Now().Sub(start).Seconds()/1e6)
}
