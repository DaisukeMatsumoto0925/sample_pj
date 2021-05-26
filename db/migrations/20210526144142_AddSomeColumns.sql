
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE samples (
  id int UNSIGNED AUTO_INCREMENT PRIMARY KEY,
  first_name varchar(255) NOT NULL,
  last_name varchar(255) NOT NULL,
  super_user int(8) UNSIGNED,
  created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE IF EXISTS samples;
