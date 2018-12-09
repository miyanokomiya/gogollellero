
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE IF NOT EXISTS `tags` (
    `id` INT NOT NULL AUTO_INCREMENT,
    `title` VARCHAR(255) NOT NULL UNIQUE,
    `created_at` timestamp not null default current_timestamp,
    `updated_at` timestamp not null default current_timestamp on update current_timestamp,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE IF NOT EXISTS `post_tags` (
    `id` INT NOT NULL AUTO_INCREMENT,
    `post_id` INT NOT NULL,
    `tag_id` INT NOT NULL,
    `created_at` timestamp not null default current_timestamp,
    `updated_at` timestamp not null default current_timestamp on update current_timestamp,
    PRIMARY KEY (`id`),
    UNIQUE KEY uq_post_tag (post_id, tag_id),
    CONSTRAINT fk_post_id
      FOREIGN KEY (post_id) 
      REFERENCES posts (id)
      ON DELETE CASCADE ON UPDATE CASCADE,
    CONSTRAINT fk_tag_id
      FOREIGN KEY (tag_id) 
      REFERENCES tags (id)
      ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE = InnoDB DEFAULT CHARSET=utf8;

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE `post_tags`;
DROP TABLE `tags`;
