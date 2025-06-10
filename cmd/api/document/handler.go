package document

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/marioanchevski/docu-reach/service/parser"
	"github.com/marioanchevski/docu-reach/types"
	"github.com/marioanchevski/docu-reach/utils"
)

type Handler struct {
	documentStore types.DocumentStore
}

func NewHandler(ds types.DocumentStore) *Handler {
	return &Handler{
		documentStore: ds,
	}
}

func (h *Handler) CreateDocumentHandler(w http.ResponseWriter, r *http.Request) {
	var payload types.CreateDocumentRequest

	if err := utils.ReadJSON(w, r, &payload); err != nil {
		utils.WriteErrorResponse[any](w, http.StatusBadRequest, err.Error())
		return
	}

	if err := payload.Validate(); err != nil {
		utils.WriteErrorResponse[any](w, http.StatusBadRequest, err.Error())
		return
	}

	response := h.documentStore.Create(payload)
	utils.WriteSuccessResponse(w, http.StatusCreated, "SUCCESS", response)
}

func (h *Handler) FindDocumetByIdHandler(w http.ResponseWriter, r *http.Request) {

	id, err := getIdRequestParam(r)

	if err != nil {
		utils.WriteErrorResponse[any](w, http.StatusBadRequest, "Invalid Id format") // ToDo maybe return 404 for security purposes
		return
	}

	document, err := h.documentStore.FindById(id)
	if err != nil {
		utils.WriteErrorResponse[any](w, http.StatusNotFound, err.Error())
		return
	}

	utils.WriteSuccessResponse(w, http.StatusOK, "SUCCESS", document)

}

func (h *Handler) FindAllDocumentsHandler(w http.ResponseWriter, r *http.Request) {
	utils.WriteSuccessResponse(w, http.StatusOK, "SUCCESS", h.documentStore.FindAll())
}

func (h *Handler) DeleteDocumentByIdHandler(w http.ResponseWriter, r *http.Request) {

	id, err := getIdRequestParam(r)
	if err != nil {
		utils.WriteErrorResponse[any](w, http.StatusBadRequest, "Invalid Id format") // ToDo maybe return 404 for security purposes
		return
	}

	if err := h.documentStore.DeleteById(id); err != nil {
		utils.WriteErrorResponse[any](w, http.StatusNotFound, err.Error())
		return
	}

	utils.WriteSuccessResponse[any](w, http.StatusOK, "Document deleted", nil)
}

func (h *Handler) FilterDocumentsHandler(w http.ResponseWriter, r *http.Request) {
	titleParam := r.URL.Query().Get("title")
	descParam := r.URL.Query().Get("desc")

	if titleParam == "" && descParam == "" {
		utils.WriteErrorResponse[any](w, http.StatusBadRequest, "Invalid search parameters")
	}

	op := r.URL.Query().Get("op")
	if op != "or" {
		op = "and"
	}

	descInclude, descExclude := parser.ParseSearchTerms(descParam)
	titleInclude, titleExclude := parser.ParseSearchTerms(titleParam)

	docFilter := types.DocumentFilter{
		TitleInclude: titleInclude,
		TitleExclude: titleExclude,
		DescInclude:  descInclude,
		DescExclude:  descExclude,
		Operator:     op,
	}

	utils.WriteSuccessResponse(w, http.StatusOK, "SUCCESS", h.documentStore.Filter(docFilter))
}

func getIdRequestParam(r *http.Request) (int, error) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		return id, fmt.Errorf("Invalid Id format")
	}

	return id, nil
}
