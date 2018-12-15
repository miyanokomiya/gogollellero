-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE IF NOT EXISTS `post_parents` (
    `id` INT NOT NULL AUTO_INCREMENT,
    `view_count` INT DEFAULT 0,
    `current_id` INT NOT NULL,
    `created_at` timestamp not null default current_timestamp,
    `updated_at` timestamp not null default current_timestamp on update current_timestamp,
    PRIMARY KEY (`id`),
    UNIQUE KEY uq_post_current (current_id)
) ENGINE = InnoDB DEFAULT CHARSET=utf8mb4 COLLATE utf8mb4_bin;
INSERT INTO post_parents (current_id) SELECT id FROM posts;
ALTER TABLE posts ADD `type` INT NOT NULL DEFAULT 1 AFTER `lesson`;
ALTER TABLE posts ADD `post_parent_id` INT AFTER `lesson`;
UPDATE posts, post_parents SET posts.post_parent_id = post_parents.id WHERE posts.id = post_parents.current_id;
ALTER TABLE posts ADD CONSTRAINT fk_post_parent_id FOREIGN KEY (post_parent_id) REFERENCES post_parents(id) ON DELETE CASCADE ON UPDATE CASCADE;
ALTER TABLE `post_parents` DROP COLUMN `current_id`;

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled backo
ALTER TABLE `posts` DROP FOREIGN KEY fk_post_parent_id;
ALTER TABLE `posts` DROP COLUMN `post_parent_id`;
ALTER TABLE `posts` DROP COLUMN `type`;
DROP TABLE `post_parents`;