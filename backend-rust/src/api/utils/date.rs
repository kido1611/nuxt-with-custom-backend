use chrono::NaiveDateTime;
use time::OffsetDateTime;
use tower_cookies::cookie::Expiration;

pub fn naivedatetime_to_expiration(datetime: &NaiveDateTime) -> Expiration {
    let result = OffsetDateTime::from_unix_timestamp(datetime.and_utc().timestamp()).unwrap();

    Expiration::from(result)
}
