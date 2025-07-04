use std::sync::Arc;

use uuid::Uuid;
use validator::Validate;

use crate::{
    application::models::note::{CreateNoteRequest, NoteResponse},
    domain::{models::note::NewNote, repository::note::NoteRepository},
    errors::Error,
};

pub struct NoteUseCase {
    pub note_repository: Arc<dyn NoteRepository>,
}

impl NoteUseCase {
    pub async fn list_notes(&self, user_id: String) -> Result<Vec<NoteResponse>, Error> {
        let notes = self
            .note_repository
            .list(user_id)
            .await
            .map_err(Error::Database)?;

        let note_responses: Vec<NoteResponse> =
            notes.iter().map(NoteResponse::from_note_entity).collect();

        Ok(note_responses)
    }

    pub async fn create_note(
        &self,
        user_id: String,
        request: CreateNoteRequest,
    ) -> Result<NoteResponse, Error> {
        request.validate()?;

        let note_id = Uuid::now_v7();
        let note = NewNote {
            id: note_id,
            user_id: Uuid::try_parse(&user_id).unwrap(),
            title: request.title,
            description: request.description,
        };
        self.note_repository.create(note).await?;

        let note = self
            .note_repository
            .get(user_id, note_id.to_string())
            .await?
            .unwrap(); // use unwrap because note must be exist after create

        let note_response = NoteResponse::from_note_entity(&note);

        Ok(note_response)
    }

    pub async fn get_note(
        &self,
        user_id: String,
        note_id: String,
    ) -> Result<Option<NoteResponse>, Error> {
        let note = match self.note_repository.get(user_id, note_id).await? {
            Some(note) => note,
            None => return Ok(None),
        };

        let note_response = NoteResponse::from_note_entity(&note);

        Ok(Some(note_response))
    }

    pub async fn delete_note(&self, user_id: String, note_id: String) -> Result<(), Error> {
        self.note_repository.delete(user_id, note_id).await?;

        Ok(())
    }
}
