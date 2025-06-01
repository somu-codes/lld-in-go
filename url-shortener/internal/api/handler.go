package api

import (
	"net/http"
	"url-shortener/internal/service"
)
import "github.com/gin-gonic/gin"

type Handler struct {
	urlSvc *service.URLService
}

func NewHandler(s *service.URLService) *Handler {
	return &Handler{urlSvc: s}
}

type ShortenRequest struct {
	URL string `json:"url" binding:"required,url"`
}

type ResolveRequest struct {
	Code string `json:"code" binding:"required"`
}

func (h *Handler) HandleShorten(c *gin.Context) {
	var req ShortenRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "invalid request",
			"details": err.Error(),
		})
	}

	code, err := h.urlSvc.Shorten(req.URL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not shorten url", "details": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{
		"short_url": code,
	})
}

func (h *Handler) HandleResolve(c *gin.Context) {
	code := c.Param("code")

	url, err := h.urlSvc.Resolve(code)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not resolve"})
		return
	}
	c.Redirect(http.StatusFound, url)
}
