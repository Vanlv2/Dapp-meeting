package handlers

import (
	"net/http"

	"Dapp-meeting/bandwidthSaving/services"

	"github.com/labstack/echo/v4"
)

type MeetingHandler struct {
	Cloudflare *services.CloudflareService
}

func NewMeetingHandler(cfService *services.CloudflareService) *MeetingHandler {
	return &MeetingHandler{
		Cloudflare: cfService,
	}
}

// OptimizeMeetingHandler nhận payload từ client, gọi Cloudflare để tạo live input và trả về URL tối ưu.
func (h *MeetingHandler) OptimizeMeetingHandler(c echo.Context) error {
	var req struct {
		Url string `json:"meeting_url"`
	}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "dữ liệu không hợp lệ"})
	}

	url := req.Url
	if url == "" {
		url = "default meeting stream"
	}

	optimizedURL, err := h.Cloudflare.CreateLiveInput(url)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"optimized_url": optimizedURL,
	})
}
