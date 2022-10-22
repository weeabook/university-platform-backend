-- Tables

CREATE TABLE users (
    id            SERIAL NOT NULL UNIQUE,
    username      TEXT NOT NULL,
    surname       TEXT NOT NULL,
    patronymic    TEXT,
    email         TEXT NOT NULL,
    password_hash TEXT NOT NULL,
    recordbook    INT
);

CREATE TABLE roles (
    id        SERIAL NOT NULL UNIQUE,
    role_name TEXT NOT NULL UNIQUE
);

CREATE TABLE news (
    id         SERIAL NOT NULL UNIQUE,
    img_url    TEXT NOT NULL,
    title      TEXT NOT NULL,
    content    TEXT NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

CREATE TABLE comments (
    id      SERIAL NOT NULL UNIQUE,
    title   TEXT NOT NULL,
    content TEXT NOT NULL
);

CREATE TABLE categories (
    id      SERIAL NOT NULL UNIQUE,
    content TEXT NOT NULL UNIQUE
);

CREATE TABLE groups (
    id      SERIAL NOT NULL UNIQUE,
    content TEXT NOT NULL UNIQUE
);

CREATE TABLE lessons (
    id         SERIAL NOT NULL UNIQUE,
    group_id   INT REFERENCES groups(id) ON DELETE CASCADE NOT NULL UNIQUE,
    up_week    BOOLEAN NOT NULL,
    week_day   INT NOT NULL,
    time_start TIME NOT NULL,
    time_end   TIME NOT NULL,
    classroom  TEXT NOT NULL,
    subject    TEXT NOT NULL
);

-- Link tables

CREATE TABLE user_role (
    id      SERIAL NOT NULL UNIQUE,
    user_id INT REFERENCES users(id) ON DELETE CASCADE NOT NULL UNIQUE,
    role_id INT REFERENCES roles(id) ON DELETE CASCADE NOT NULL UNIQUE
);

CREATE TABLE user_group (
    id       SERIAL NOT NULL UNIQUE,
    user_id  INT REFERENCES users(id) ON DELETE CASCADE NOT NULL UNIQUE,
    group_id INT REFERENCES users(id) ON DELETE CASCADE NOT NULL
);

CREATE TABLE news_category (
    id          SERIAL NOT NULL UNIQUE,
    news_id     INT REFERENCES news(id) ON DELETE CASCADE NOT NULL UNIQUE,
    category_id INT REFERENCES categories(id) ON DELETE CASCADE NOT NULL
);

CREATE TABLE user_comments (
    id         SERIAL NOT NULL UNIQUE,
    user_id    INT REFERENCES users(id) ON DELETE CASCADE NOT NULL UNIQUE,
    comment_id INT REFERENCES comments(id) ON DELETE CASCADE NOT NULL UNIQUE
);
