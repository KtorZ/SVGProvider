package ttnsvg

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
)

type svgHandler struct {
	svgSrc []byte
	regReq regexp.Regexp
	regSVG regexp.Regexp
}

func (h *svgHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "image/svg+xml")
	w.WriteHeader(http.StatusOK)

	city := h.regReq.FindSubmatch([]byte(r.URL.Path))

	if city == nil {
		w.Write(h.regSVG.ReplaceAll(h.svgSrc, *new([]byte)))
		return
	}

	w.Write(h.regSVG.ReplaceAll(h.svgSrc, city[1]))
}

func init() {
	svgSrc, err := ioutil.ReadFile("ttn_logo.svg")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	http.Handle("/", &svgHandler{
		[]byte(svgSrc),
		*regexp.MustCompile("^/([\\s\\p{L}]+)/?$"),
		*regexp.MustCompile("(@city@)"),
	})
}
