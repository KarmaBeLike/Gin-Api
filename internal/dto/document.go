package dto

type CreateDocumentRequest struct { // для запроса на создание документа
	Title   string `json:"title"`
	Content string `json:"content"`
}

type GetDocumentRequest struct {
	Title string `json:"title"`
}
