package handler

import (
	"context"

	"start-with-go/internal/api"
	"start-with-go/internal/repository"
)

type NoteHandler struct {
	repo *repository.NoteRepository
}

func NewNoteHandler(repo *repository.NoteRepository) *NoteHandler {
	return &NoteHandler{repo: repo}
}

func (h *NoteHandler) NotesList(ctx context.Context, request api.NotesListRequestObject) (api.NotesListResponseObject, error) {
	notes, err := h.repo.GetAll(ctx)
	if err != nil {
		return api.NotesListdefaultJSONResponse{Body: api.ApiError{Error: err.Error()}, StatusCode: 500}, nil
	}

	response := make(api.NotesList200JSONResponse, len(notes))
	for i, n := range notes {
		response[i] = api.Note{
			Id:        n.ID,
			Title:     n.Title,
			Content:   n.Content,
			CreatedAt: n.CreatedAt,
			UpdatedAt: n.UpdatedAt,
		}
	}
	return response, nil
}

func (h *NoteHandler) NotesRead(ctx context.Context, request api.NotesReadRequestObject) (api.NotesReadResponseObject, error) {
	note, err := h.repo.GetByID(ctx, request.Id)
	if err != nil {
		return api.NotesReaddefaultJSONResponse{Body: api.ApiError{Error: err.Error()}, StatusCode: 404}, nil
	}

	return api.NotesRead200JSONResponse{
		Id:        note.ID,
		Title:     note.Title,
		Content:   note.Content,
		CreatedAt: note.CreatedAt,
		UpdatedAt: note.UpdatedAt,
	}, nil
}

func (h *NoteHandler) NotesCreate(ctx context.Context, request api.NotesCreateRequestObject) (api.NotesCreateResponseObject, error) {
	note, err := h.repo.Create(ctx, request.Body.Title, request.Body.Content)
	if err != nil {
		return api.NotesCreatedefaultJSONResponse{Body: api.ApiError{Error: err.Error()}, StatusCode: 500}, nil
	}

	return api.NotesCreate200JSONResponse{
		Id:        note.ID,
		Title:     note.Title,
		Content:   note.Content,
		CreatedAt: note.CreatedAt,
		UpdatedAt: note.UpdatedAt,
	}, nil
}

func (h *NoteHandler) NotesUpdate(ctx context.Context, request api.NotesUpdateRequestObject) (api.NotesUpdateResponseObject, error) {
	note, err := h.repo.Update(ctx, request.Id, request.Body.Title, request.Body.Content)
	if err != nil {
		return api.NotesUpdatedefaultJSONResponse{Body: api.ApiError{Error: err.Error()}, StatusCode: 500}, nil
	}

	return api.NotesUpdate200JSONResponse{
		Id:        note.ID,
		Title:     note.Title,
		Content:   note.Content,
		CreatedAt: note.CreatedAt,
		UpdatedAt: note.UpdatedAt,
	}, nil
}

func (h *NoteHandler) NotesDelete(ctx context.Context, request api.NotesDeleteRequestObject) (api.NotesDeleteResponseObject, error) {
	if err := h.repo.Delete(ctx, request.Id); err != nil {
		return api.NotesDeletedefaultJSONResponse{Body: api.ApiError{Error: err.Error()}, StatusCode: 500}, nil
	}

	return api.NotesDelete204Response{}, nil
}
