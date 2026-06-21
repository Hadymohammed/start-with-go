package repository

import (
	"context"

	"start-with-go/internal/db"
)

type NoteRepository struct {
	queries *db.Queries
}

func NewNoteRepository(queries *db.Queries) *NoteRepository {
	return &NoteRepository{queries: queries}
}

func (r *NoteRepository) GetAll(ctx context.Context) ([]db.Note, error) {
	return r.queries.ListNotes(ctx)
}

func (r *NoteRepository) GetByID(ctx context.Context, id int64) (db.Note, error) {
	return r.queries.GetNote(ctx, id)
}

func (r *NoteRepository) Create(ctx context.Context, title, content string) (db.Note, error) {
	return r.queries.CreateNote(ctx, db.CreateNoteParams{
		Title:   title,
		Content: content,
	})
}

func (r *NoteRepository) Update(ctx context.Context, id int64, title, content string) (db.Note, error) {
	return r.queries.UpdateNote(ctx, db.UpdateNoteParams{
		ID:      id,
		Title:   title,
		Content: content,
	})
}

func (r *NoteRepository) Delete(ctx context.Context, id int64) error {
	return r.queries.DeleteNote(ctx, id)
}
