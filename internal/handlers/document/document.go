package document

import (
	"errors"
	"net/http"

	"Gin-Api/config"
	"Gin-Api/internal/dto"
	"Gin-Api/internal/model"
	service "Gin-Api/internal/service/document"

	"github.com/gin-gonic/gin"
)

type DocumentClient struct {
	docServ *service.DocumentService
}

func NewDocumentClient(service *service.DocumentService) *DocumentClient {
	return &DocumentClient{
		docServ: service,
	}
}

func (dc *DocumentClient) Routes(r *gin.Engine, cfg *config.Config) {
	r.GET("getdoc", dc.GetDocument)
	r.POST("documents", dc.CreateDocument)
}

func (dc *DocumentClient) CreateDocument(ctx *gin.Context) {
	var request dto.CreateDocumentRequest
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = dc.docServ.CreateDocument(ctx, &request)
	if err != nil {
		if errors.Is(err, model.ErrDuplicateTitle) {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": "a title already exists"})
			return
		} else {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "document successfully created"})
}

func (dc *DocumentClient) GetDocument(ctx *gin.Context) {
	var request dto.GetDocumentRequest
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	document, err := dc.docServ.GetDocument(&request)
	if err != nil {
		if errors.Is(err, model.ErrDocumentNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Document not found"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"1": gin.H{"message": "document successfully read"},
		"2": gin.H{
			"ID":    document.ID,
			"title": document.Title, "content": document.Content,
		},
	})
}
