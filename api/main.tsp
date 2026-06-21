import "@typespec/http";

using Http;

@service(#{ title: "Notes API" })
@useAuth(ApiKeyAuth<ApiKeyLocation.header, "X-API-Key">)
namespace NotesService;

model Note {
  id: int64;
  title: string;
  content: string;
  created_at: utcDateTime;
  updated_at: utcDateTime;
}

model NoteRequest {
  title: string;
  content: string;
}

@error
model ApiError {
  error: string;
}

@route("/api/v1/notes")
@tag("Notes")
interface Notes {
  @get list(): Note[] | ApiError;
  @get read(@path id: int64): Note | ApiError;
  @post create(@body body: NoteRequest): Note | ApiError;
  @put update(@path id: int64, @body body: NoteRequest): Note | ApiError;
  @delete delete(@path id: int64): void | ApiError;
}
