
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
ALTER TABLE post_parents ADD `status` INT NOT NULL DEFAULT 1 AFTER `id`;

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
ALTER TABLE `post_parents` DROP COLUMN `status`;
