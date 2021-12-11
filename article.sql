DROP TABLE IF EXISTS articles;
create table articles (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    author varchar (255) not null,
    title varchar (255) null,
    body varchar (255) not null,
    created_at timestamp,
    updated_at timestamp,
    deleted_at timestamp
);