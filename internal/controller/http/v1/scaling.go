package v1

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/nfnt/resize"
	"html"
	"net/http"
	"resize-server/internal/usecase"
	"resize-server/pkg/logger"
	"strconv"
	"strings"
)

type scalingRoutes struct {
	s *usecase.ScalingUseCase
	l logger.Interface
}

func newScalingRoutes(handler *gin.RouterGroup, s *usecase.ScalingUseCase, l logger.Interface) {
	r := scalingRoutes{s, l}

	h := handler.Group("/scale")
	{
		h.GET("/do", r.doScale)
	}
}

func (r *scalingRoutes) doScale(c *gin.Context) {
	pp, ok := c.GetQuery("path")
	if !ok {
		r.l.Error("missing path parameter", "http - v1 - doScale")
		errorResponse(c, http.StatusBadRequest, "missing path parameter")

		return
	}

	path := html.EscapeString(pp)
	r.s.Scale.Path = path

	size, ok := c.GetQuery("size")
	if !ok {
		r.l.Error("missing size parameter", "http - v1 - doScale")
		errorResponse(c, http.StatusBadRequest, "missing size parameter")

		return
	}

	dims := strings.Split(size, "x")
	if len(dims) == 1 {
		ok := r.s.SetParamsFromPreset(dims[0])
		if !ok {
			r.l.Error("unknown preset")
			errorResponse(c, http.StatusBadRequest, "unknown preset")

			return
		}
	} else {
		w, _ := strconv.ParseUint(dims[0], 10, 0)
		h, _ := strconv.ParseUint(dims[1], 10, 0)
		r.s.Scale.Width = uint(w)
		r.s.Scale.Height = uint(h)
	}

	s := strings.Split(path, "/")
	file := strings.Split(s[len(s)-1], ".")
	r.s.Scale.OutName = file[0] + "_" + strconv.Itoa(int(r.s.Scale.Width)) + "x" + strconv.Itoa(int(r.s.Scale.Height)) + "." + file[1]
	r.s.Scale.Ext = file[1]

	ib, _ := r.s.GetImageBytes(r.s.Scale)
	buf := bytes.NewReader(ib)

	img, err := r.s.Decode(buf, r.s.Scale)
	if err != nil {
		r.l.Error(err)
		errorResponse(c, http.StatusInternalServerError, "error while decoding image bytes")

		return
	}

	m := resize.Resize(r.s.Scale.Width, r.s.Scale.Height, img, r.s.Scale.InterFunc)
	ebuf, err := r.s.Encode(m, r.s.Scale)
	if err != nil {
		r.l.Error(err)
		errorResponse(c, http.StatusInternalServerError, "error while encoding image bytes")

		return
	}

	c.Header("Content-Type", r.s.Scale.MimeType())
	_, err = c.Writer.Write(ebuf.Bytes())
	if err != nil {
		r.l.Error(err)
		errorResponse(c, http.StatusInternalServerError, "error while writing response body")
	}
}
