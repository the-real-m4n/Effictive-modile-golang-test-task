package handler

import (
	"net/http"
	"strconv"
	models "subscriptions-service/internal/model"
	"subscriptions-service/internal/repository"

	"time"

	"log"

	"github.com/gin-gonic/gin"
)

type SubscriptionHandler struct {
	repo *repository.SubscriptionRepo
}

func NewSubscriprionHandler(repo *repository.SubscriptionRepo) *SubscriptionHandler {
	return &SubscriptionHandler{repo: repo}
}

type createRequest struct {
	ServiceName string `json:"service_name" binding:"required" example:"Yandex Plus"`
	Price       int    `json:"price" binding:"required" example:"400"`
	UserID      string `json:"user_id" binding:"required" example:"60601fee-2bf1-4721-ae6f-7636e79a0cba"`
	StartDate   string `json:"start_date" binding:"required" example:"2025-07"`
	EndDate     string `json:"end_date" example:"2025-12"`
}

// Create godoc
// @Summary Создать подписку
// @Description Добавляет новую подписку в систему
// @Tags subscriptions
// @Accept json
// @Produce json
// @Param subscription body createRequest true "Subscription"
// @Success 201 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /subscriptions [post]
func (h *SubscriptionHandler) Create(c *gin.Context) {
	var req createRequest
	// Читаем тело запроса (JSON)
	if err := c.ShouldBindBodyWithJSON(&req); err != nil { // сравниваем входные данные (с) с шаблонной структурой (req) createRequest
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}) // если входные кривые возвращаем ошибку
		return
	}
	// валидация начала подписки
	start, err := time.Parse("2006-01", req.StartDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid start_date format (expected YYYY-MM)"})
	}
	// валидация окончания подписки
	var end *time.Time
	if req.EndDate != "" {
		parsed, err := time.Parse("2006-01", req.EndDate)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid end_date format (expected YYYY-MM)"})
			return
		}
		end = &parsed

	}
	sub := models.Subscription{
		ServiceName: req.ServiceName,
		Price:       req.Price,
		UserID:      req.UserID,
		StartDate:   start,
		EndDate:     end,
	}
	// Вызваем репозиторий, чтобы записать данные в БД
	if err := h.repo.Create(c.Request.Context(), sub); err != nil {
		log.Printf("Ошибка на стороне сервера %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create subscription"}) // если выбивает ошибку то проблема в функции repository
		return
	}
	c.JSON(http.StatusCreated, gin.H{"status": "created"})
}

// Create godoc
// @Summary прочитать подписки
// @Description читает подписки
// @Tags subscriptions
// @Produce json
// @Success 201 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /subscriptions [get]
func (h *SubscriptionHandler) GetAll(c *gin.Context) { // хендлер для вывода всех подписок
	subs, err := h.repo.GetAll(c.Request.Context())

	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Ошибка чтения подписок"})
		return

	}
	c.JSON(http.StatusOK, subs)
}

// GetByID godoc
// @Summary Получить подписку по ID пользователя
// @Description Возвращает одну подписку
// @Tags subscriptions
// @Produce json
// @Param id path string true "Subscription ID"
// @Success 200 {object} models.Subscription
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /subscriptions/{id} [get]
func (h *SubscriptionHandler) GetByID(c *gin.Context) {
	idParam := c.Param("id")

	sub, err := h.repo.GetByID(c.Request.Context(), idParam)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "subscription not found"})
		return
	}

	c.JSON(http.StatusOK, sub)
}

// Update godoc
// @Summary Обновить подписку
// @Description Изменяет данные подписки по ID
// @Tags subscriptions
// @Accept json
// @Produce json
// @Param id path int true "Subscription ID"
// @Param subscription body createRequest true "Subscription data"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /subscriptions/{id} [put]
func (h *SubscriptionHandler) Update(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	var req createRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	// Валидация дат
	start, err := time.Parse("2006-01", req.StartDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid start_date format"})
		return
	}

	var end *time.Time
	if req.EndDate != "" {
		parsed, err := time.Parse("2006-01", req.EndDate)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid end_date format"})
			return
		}
		end = &parsed
	}

	sub := models.Subscription{
		ID:          id,
		ServiceName: req.ServiceName,
		Price:       req.Price,
		UserID:      req.UserID,
		StartDate:   start,
		EndDate:     end,
	}

	if err := h.repo.Update(c.Request.Context(), sub); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update subscription"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "updated"})
}

// Delete godoc
// @Summary Удалить подписку по ID
// @Description Возвращает статус
// @Tags subscriptions
// @Produce json
// @Param id path int true "Subscription ID"
// @Success 200 {object} models.Subscription
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /subscriptions/{id} [delete]
func (h *SubscriptionHandler) Delete(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	if err := h.repo.Delete(c.Request.Context(), id); err != nil {
		if err.Error() == "not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "subscription not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete subscription"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "deleted"})

}

// GetTotal godoc
// @Summary Подсчитать сумму подписок
// @Description Возвращает суммарную стоимость всех подписок за период с фильтрацией
// @Tags subscriptions
// @Produce json
// @Param user_id query string true "User ID"
// @Param service_name query string true "Service name"
// @Param from query string true "Start period (YYYY-MM)"
// @Param to query string true "End period (YYYY-MM)"
// @Success 200 {object} map[string]int
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /subscriptions/total [get]
func (h *SubscriptionHandler) GetTotalPrice(c *gin.Context) {
	userID := c.Query("user_id")
	serviceName := c.Query("service_name")
	fromStr := c.Query("from")
	toStr := c.Query("to")

	if userID == "" || serviceName == "" || fromStr == "" || toStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing required params"})
		return
	}

	from, err := time.Parse("2006-01", fromStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid from format"})
		return
	}

	to, err := time.Parse("2006-01", toStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid to format"})
		return
	}

	total, err := h.repo.GetTotalPrice(c.Request.Context(), userID, serviceName, from, to)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to calculate total"})
		return
	}
	log.Printf("Server respons %v", total)
	c.JSON(http.StatusOK, gin.H{"total": total})
}
