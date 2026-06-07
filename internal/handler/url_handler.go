package handler

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/soorajsomans/url-shortener/internal/dto"
	"github.com/soorajsomans/url-shortener/internal/service"
)

type URLHandler struct {
	urlService service.URLService
}

func NewURLHandler(
	urlService service.URLService,
) *URLHandler {
	return &URLHandler{
		urlService: urlService,
	}
}

func (h *URLHandler) RegisterRoutes(
	mux *http.ServeMux,
) {
	mux.HandleFunc(
		"/shorten",
		h.Shorten,
	)

	mux.HandleFunc(
		"/",
		h.Redirect,
	)
}

func writeJSON(
	w http.ResponseWriter,
	statusCode int,
	payload any,
) {
	w.Header().Set(
		"Content-Type", "application/json",
	)
	w.WriteHeader(statusCode)

	_ = json.NewEncoder(w).Encode(payload)
}

func writeError(
	w http.ResponseWriter,
	statusCode int,
	message string,
) {
	writeJSON(
		w,
		statusCode,
		dto.ErrorResponse{
			Message: message,
		},
	)
}

// Shorten godoc
//
// @Summary Create short URL
// @Description Create a short URL from a long URL
// @Tags URLs
// @Accept json
// @Produce json
// @Param request body dto.ShortenRequest true "URL payload"
// @Success 201 {object} dto.ShortenResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 405 {object} dto.ErrorResponse
// @Router /shorten [post]
func (h *URLHandler) Shorten(
	w http.ResponseWriter,
	r *http.Request,
) {
	if r.Method != http.MethodPost {
		writeError(
			w,
			http.StatusMethodNotAllowed,
			"method not allowed",
		)
		return
	}

	var req dto.ShortenRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(
			w,
			http.StatusBadRequest,
			"invalid request body",
		)
		return
	}

	urlEntity, err := h.urlService.Shorten(
		r.Context(),
		req.URL,
	)
	if err != nil {
		writeError(
			w,
			http.StatusBadRequest,
			err.Error(),
		)
		return
	}

	writeJSON(
		w,
		http.StatusCreated,
		dto.ShortenResponse{
			ShortCode: urlEntity.ShortCode,
		},
	)
}

// Redirect godoc
//
// @Summary Redirect to original URL
// @Description Redirects using short code
// @Tags URLs
// @Produce plain
// @Param code path string true "Short code"
// @Success 302
// @Failure 404 {object} dto.ErrorResponse
// @Router /{code} [get]
func (h *URLHandler) Redirect(
	w http.ResponseWriter,
	r *http.Request,
) {
	code := strings.TrimPrefix(
		r.URL.Path,
		"/",
	)

	urlEntity, err := h.urlService.Resolve(
		r.Context(),
		code,
	)

	if err != nil {
		writeError(
			w,
			http.StatusNotFound,
			err.Error(),
		)
		return
	}

	http.Redirect(
		w,
		r,
		urlEntity.LongURL,
		http.StatusFound,
	)
}
