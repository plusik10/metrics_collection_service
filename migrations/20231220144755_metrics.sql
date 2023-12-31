-- +goose Up
CREATE TABLE IF NOT EXISTS event
(
    id              BIGSERIAL Primary Key,
    event_type      VARCHAR(255) NOT NULL,
    screen_name     VARCHAR(255) NOT NULL,
    action          VARCHAR(255) NOT NULL,
    event_timestamp timestamp    NOT NULL default now()
);
-- +goose Down
DROP TABLE IF EXISTS event;