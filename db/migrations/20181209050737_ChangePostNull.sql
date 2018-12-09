
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
ALTER TABLE posts MODIFY COLUMN problem TEXT;
ALTER TABLE posts MODIFY COLUMN solution TEXT;
ALTER TABLE posts MODIFY COLUMN lesson TEXT;

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
ALTER TABLE posts MODIFY COLUMN problem TEXT NOT NULL;
ALTER TABLE posts MODIFY COLUMN solution TEXT NOT NULL;
ALTER TABLE posts MODIFY COLUMN lesson TEXT NOT NULL;
