CREATE TABLE IF NOT EXISTS Documents (
    Id          UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    Title       VARCHAR(255)    NOT NULL,
    Content     TEXT            NOT NULL,
    CreatedAt   TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
