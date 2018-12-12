
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
ALTER TABLE users CONVERT TO CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;
ALTER TABLE posts CONVERT TO CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;
ALTER TABLE tags CONVERT TO CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;
ALTER TABLE post_tags CONVERT TO CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
ALTER TABLE post_tags CONVERT TO CHARACTER SET utf8 COLLATE utf8_general_ci;
ALTER TABLE tags CONVERT TO CHARACTER SET utf8 COLLATE utf8_general_ci;
ALTER TABLE posts CONVERT TO CHARACTER SET utf8 COLLATE utf8_general_ci;
ALTER TABLE users CONVERT TO CHARACTER SET utf8 COLLATE utf8_general_ci;

