package app

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
	"net/http"
	"products/internal/logger"
	"products/internal/models"
	"products/internal/services"
	"time"
)

type ProductHandlers struct {
	service services.ProductService
}

func (ch *ProductHandlers) save(ginCtx *gin.Context) {
	var (
		productRequest models.Product
		err            error
		ctx, cancel    = getContext()
	)

	defer cancel()
	err = json.NewDecoder(ginCtx.Request.Body).Decode(&productRequest)

	if err != nil {
		writeResponse(ginCtx.Writer, http.StatusBadRequest, err.Error())
	} else {
		job, appError := ch.service.Save(ctx, productRequest)

		if appError != nil {
			writeResponse(ginCtx.Writer, appError.Code, appError.Message)
		} else {
			writeResponse(ginCtx.Writer, http.StatusOK, job)
		}
	}
}

func writeResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		logger.Error("Data could not be encoded")
	}
}

func getContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 30*time.Second)
}
