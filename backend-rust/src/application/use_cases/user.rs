use std::sync::Arc;

use argon2::{
    Argon2, Params, PasswordHash, PasswordHasher, PasswordVerifier,
    password_hash::{SaltString, rand_core::OsRng},
};
use uuid::Uuid;
use validator::Validate;

use crate::{
    application::models::user::{LoginRequest, RegisterRequest, UserResponse},
    domain::{models::user::NewUser, repository::user::UserRepository},
    errors::Error,
};

pub struct UserUseCase {
    pub user_repository: Arc<dyn UserRepository>,
}

impl UserUseCase {
    pub async fn get_user_by_id(&self, user_id: String) -> Result<Option<UserResponse>, Error> {
        let user = match self
            .user_repository
            .get_by_id(user_id)
            .await
            .map_err(Error::Database)?
        {
            Some(user) => user,
            None => return Ok(None),
        };

        let user_response = UserResponse::from_user_entity(&user);

        Ok(Some(user_response))
    }

    pub async fn check_user(&self, request: &LoginRequest) -> Result<Option<UserResponse>, Error> {
        request.validate()?;

        let user = match self
            .user_repository
            .get_by_email(request.email.clone())
            .await
            .map_err(Error::Database)?
        {
            Some(user) => user,
            None => return Ok(None),
        };

        self.verify_password(request.password.clone(), user.password.clone())?;

        let user_response = UserResponse::from_user_entity(&user);

        Ok(Some(user_response))
    }

    pub async fn create_user(&self, request: &RegisterRequest) -> Result<UserResponse, Error> {
        request.validate()?;

        let user = self
            .user_repository
            .get_by_email(request.email.clone())
            .await
            .map_err(Error::Database)?;

        if user.is_some() {
            return Err(Error::RegisterConflict(anyhow::anyhow!(
                "Email already exists"
            )));
        }

        let password_hash = self.calculate_password_hash(request.password.clone())?;
        let new_user = NewUser {
            id: Uuid::now_v7(),
            name: request.name.clone(),
            email: request.email.clone(),
            password: password_hash,
        };

        // Create user
        self.user_repository
            .create(new_user)
            .await
            .map_err(Error::Database)?;

        let user = self
            .user_repository
            .get_by_email(request.email.clone())
            .await
            .map_err(Error::Database)?
            .unwrap(); // use unwrap because user must be already in database

        let user_response = UserResponse::from_user_entity(&user);

        Ok(user_response)
    }

    fn verify_password(&self, password: String, password_hashed: String) -> Result<(), Error> {
        let parsed_password_hash = PasswordHash::new(&password_hashed)
            .map_err(|_| Error::Other(anyhow::anyhow!("Failed parse password hash")))?;

        Argon2::default()
            .verify_password(password.as_bytes(), &parsed_password_hash)
            .map_err(|_| Error::Unauthorized(anyhow::anyhow!("Failed verify password")))
    }

    fn calculate_password_hash(&self, password: String) -> Result<String, Error> {
        let salt = SaltString::generate(&mut OsRng);
        let password_hash = Argon2::new(
            argon2::Algorithm::Argon2id,
            argon2::Version::V0x13,
            Params::new(15000, 2, 1, None).unwrap(),
        )
        .hash_password(password.as_bytes(), &salt)
        .map_err(|_| Error::Other(anyhow::anyhow!("Failed hashing password")))?
        .to_string();

        Ok(password_hash)
    }
}
