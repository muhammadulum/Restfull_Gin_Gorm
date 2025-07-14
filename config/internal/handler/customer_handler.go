package handler

import (
    "gin_restfull/internal/model"
    "gin_restfull/internal/usecase"
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
)

type CustomerHandler struct {
    usecase *usecase.CustomerUseCase
}

func NewCustomerHandler(uc *usecase.CustomerUseCase) *CustomerHandler {
    return &CustomerHandler{usecase: uc}
}

func (h *CustomerHandler) RegisterRoutes(rg *gin.RouterGroup) {
    customer := rg.Group("/customers")
    {
        customer.POST("", h.Create)
        customer.GET("", h.GetAll)
        customer.GET("/:id", h.GetByID)
        customer.PUT("/:id", h.Update)
        customer.DELETE("/:id", h.Delete)
    }
}

func (h *CustomerHandler) Create(c *gin.Context) {
    var req model.CustomerRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    if err := h.usecase.Create(req); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusCreated, gin.H{"message": "customer created"})
}

func (h *CustomerHandler) GetAll(c *gin.Context) {
    customers, err := h.usecase.GetAll()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, customers)
}

func (h *CustomerHandler) GetByID(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    customer, err := h.usecase.GetByID(uint(id))
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "customer not found"})
        return
    }
    c.JSON(http.StatusOK, customer)
}

func (h *CustomerHandler) Update(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    var req model.CustomerRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    if err := h.usecase.Update(uint(id), req); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "customer updated"})
}

func (h *CustomerHandler) Delete(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    if err := h.usecase.Delete(uint(id)); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "customer deleted"})
}
