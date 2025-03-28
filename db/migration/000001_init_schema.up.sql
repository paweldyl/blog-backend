CREATE TYPE "like_value" AS ENUM (
  'like',
  'dislike'
);

CREATE TABLE "users" (
  "login" varchar PRIMARY KEY,
  "hashed_password" varchar NOT NULL,
  "username" varchar UNIQUE NOT NULL,
  "edited_at" timestamptz NOT NULL DEFAULT (now()),
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "sessions" (
  "id" uuid PRIMARY KEY,
  "login" varchar NOT NULL,
  "refresh_token" varchar NOT NULL,
  "is_blocked" boolean NOT NULL DEFAULT false,
  "expires_at" timestamptz NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "posts" (
  "id" uuid PRIMARY KEY,
  "title" varchar NOT NULL,
  "short_desc" varchar NOT NULL,
  "description" varchar NOT NULL,
  "author_login" varchar NOT NULL,
  "likes_amount" int NOT NULL DEFAULT 0,
  "dislikes_amount" int NOT NULL DEFAULT 0,
  "edited_at" timestamptz NOT NULL DEFAULT (now()),
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "posts_likes" (
  "user_login" varchar NOT NULL,
  "post_id" uuid NOT NULL,
  "value" like_value NOT NULL
);

CREATE TABLE "comments" (
  "id" uuid PRIMARY KEY,
  "text" varchar NOT NULL,
  "author_login" varchar NOT NULL,
  "post_id" uuid NOT NULL
);

CREATE INDEX ON "sessions" ("login");

CREATE INDEX ON "posts" ("author_login");

CREATE INDEX ON "posts_likes" ("user_login");

CREATE INDEX ON "posts_likes" ("post_id");

CREATE UNIQUE INDEX ON "posts_likes" ("user_login", "post_id");

CREATE INDEX ON "comments" ("author_login");

CREATE INDEX ON "comments" ("post_id");

CREATE INDEX ON "comments" ("author_login", "post_id");

ALTER TABLE "sessions" ADD FOREIGN KEY ("login") REFERENCES "users" ("login") ON DELETE CASCADE;

ALTER TABLE "posts" ADD FOREIGN KEY ("author_login") REFERENCES "users" ("login") ON DELETE CASCADE;

ALTER TABLE "posts_likes" ADD FOREIGN KEY ("user_login") REFERENCES "users" ("login") ON DELETE CASCADE;

ALTER TABLE "posts_likes" ADD FOREIGN KEY ("post_id") REFERENCES "posts" ("id") ON DELETE CASCADE;

ALTER TABLE "comments" ADD FOREIGN KEY ("author_login") REFERENCES "users" ("login") ON DELETE CASCADE;

ALTER TABLE "comments" ADD FOREIGN KEY ("post_id") REFERENCES "posts" ("id") ON DELETE CASCADE;
