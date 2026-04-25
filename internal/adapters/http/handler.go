package http

import (
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/snufkin23/lusay/internal/core/domain"
	"github.com/snufkin23/lusay/internal/core/service"
)

type Handler struct {
	aiSvc  *service.AIService
	logger *slog.Logger
}

func NewHandler(aiSvc *service.AIService, logger *slog.Logger) *Handler {
	return &Handler{aiSvc: aiSvc, logger: logger}
}

// Router builds the fully configured Chi router.
func (h *Handler) Router() http.Handler {
	r := chi.NewRouter()

	// ─── Global middleware ────────────────────────────────────────────────
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)
	r.Use(requestLogger(h.logger))
	r.Use(middleware.Timeout(30 * time.Second))

	// ─── Routes ───────────────────────────────────────────────────────────
	r.Get("/health", h.handleHealth)

	r.Route("/api/v1", func(r chi.Router) {
		r.Post("/generate", h.handleGenerate)
	})

	return r
}

func (h *Handler) handleHealth(w http.ResponseWriter, r *http.Request) {
	respond(w, http.StatusOK, map[string]string{
		"status": "ok",
	})
}

func (h *Handler) handleGenerate(w http.ResponseWriter, r *http.Request) {
	// 1. Decode
	var req struct {
		Prompt string `json:"prompt"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondErr(w, http.StatusBadRequest, "invalid request body")
		return
	}

	// 2. Call service
	resp, err := h.aiSvc.GenerateResponse(req.Prompt)
	if err != nil {
		h.mapErr(w, err)
		return
	}

	// 3. Encode
	respond(w, http.StatusOK, map[string]interface{}{
		"content": resp.Text,
		"mood":    resp.Mood,
	})
}

func (h *Handler) mapErr(w http.ResponseWriter, err error) {
	switch {
	case errors.Is(err, domain.ErrInvalidInput), errors.Is(err, domain.ErrHarmfulContent):
		respondErr(w, http.StatusBadRequest, err.Error())
	case errors.Is(err, domain.ErrContentFiltered):
		respondErr(w, http.StatusForbidden, "content filtered by safety settings")
	case errors.Is(err, domain.ErrRateLimitExceeded):
		respondErr(w, http.StatusTooManyRequests, "ai provider rate limit exceeded")
	case errors.Is(err, domain.ErrNetworkFailure):
		respondErr(w, http.StatusBadGateway, "failed to reach ai provider")
	case errors.Is(err, domain.ErrEmptyResponse):
		respondErr(w, http.StatusInternalServerError, "ai provider returned an empty response")
	default:
		h.logger.ErrorContext(nil, "unhandled service error", slog.String("err", err.Error()))
		respondErr(w, http.StatusInternalServerError, "internal server error")
	}
}
