create table basic_info (
    id BIGSERIAL PRIMARY KEY,
    first_name VARCHAR(50) NOT NULL,
    last_name VARCHAR(50) NOT NULL,
    additional_name VARCHAR(50),
    pronouns VARCHAR(50),
    head_line VARCHAR(50)
);