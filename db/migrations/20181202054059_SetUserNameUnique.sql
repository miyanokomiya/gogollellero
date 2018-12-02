
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
alter table `users` add unique (`name`);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
alter table `users` drop index `name`;
