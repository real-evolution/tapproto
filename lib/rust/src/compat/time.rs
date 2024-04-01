use time::{Date, OffsetDateTime, Time};

use crate::gen::google::protobuf::Timestamp;

impl From<Timestamp> for OffsetDateTime {
    fn from(value: Timestamp) -> Self {
        const MAX_TIME: Time = unsafe { Time::__from_hms_nanos_unchecked(23, 59, 59, 999_999_999) };

        let seconds = value.seconds.clamp(
            Date::MIN.midnight().assume_utc().unix_timestamp(),
            Date::MAX.with_time(MAX_TIME).assume_utc().unix_timestamp(),
        );
        let nanos = value.nanos.clamp(0, 999_999_999) as u32;

        OffsetDateTime::from_unix_timestamp(seconds)
            .unwrap()
            .replace_nanosecond(nanos)
            .unwrap()
    }
}

impl From<OffsetDateTime> for Timestamp {
    fn from(value: OffsetDateTime) -> Self {
        let seconds = value.unix_timestamp();
        let nanos = value.nanosecond() as i32;

        Self { seconds, nanos }
    }
}
