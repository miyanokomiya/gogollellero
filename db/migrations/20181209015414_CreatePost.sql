-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE IF NOT EXISTS `posts` (
    `id` INT NOT NULL AUTO_INCREMENT,
    `user_id` INT NOT NULL,
    `title` VARCHAR(255) NOT NULL,
    `problem` TEXT NOT NULL,
    `solution` TEXT NOT NULL,
    `lesson` TEXT NOT NULL,
    `created_at` timestamp not null default current_timestamp,
    `updated_at` timestamp not null default current_timestamp on update current_timestamp,
    PRIMARY KEY (`id`),
    CONSTRAINT fk_user_id
      FOREIGN KEY (user_id) 
      REFERENCES users (id)
      ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE = InnoDB;

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE `posts`;