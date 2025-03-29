CREATE TYPE "like_value" AS ENUM (
  'like',
  'dislike'
);

CREATE TABLE "users" (
  "id" uuid PRIMARY KEY,
  "login" varchar NOT NULL,
  "hashed_password" varchar NOT NULL,
  "username" varchar UNIQUE NOT NULL,
  "updated_at" timestamptz NOT NULL DEFAULT (now()),
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "sessions" (
  "id" uuid PRIMARY KEY,
  "user_id" uuid NOT NULL,
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
  "user_id" uuid NOT NULL,
  "likes_amount" int NOT NULL DEFAULT 0,
  "dislikes_amount" int NOT NULL DEFAULT 0,
  "updated_at" timestamptz NOT NULL DEFAULT (now()),
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "posts_likes" (
  "user_id" uuid NOT NULL,
  "post_id" uuid NOT NULL,
  "value" like_value NOT NULL,
  "updated_at" timestamptz NOT NULL DEFAULT (now()),
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "comments" (
  "id" uuid PRIMARY KEY,
  "text" varchar NOT NULL,
  "user_id" uuid NOT NULL,
  "post_id" uuid NOT NULL,
  "updated_at" timestamptz NOT NULL DEFAULT (now()),
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE INDEX ON "sessions" ("user_id");

CREATE INDEX ON "posts" ("user_id");

CREATE INDEX ON "posts_likes" ("user_id");

CREATE INDEX ON "posts_likes" ("post_id");

CREATE UNIQUE INDEX ON "posts_likes" ("user_id", "post_id");

CREATE INDEX ON "comments" ("user_id");

CREATE INDEX ON "comments" ("post_id");

CREATE INDEX ON "comments" ("user_id", "post_id");

ALTER TABLE "sessions" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON DELETE CASCADE;

ALTER TABLE "posts" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON DELETE CASCADE;

ALTER TABLE "posts_likes" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON DELETE CASCADE;

ALTER TABLE "posts_likes" ADD FOREIGN KEY ("post_id") REFERENCES "posts" ("id") ON DELETE CASCADE;

ALTER TABLE "comments" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON DELETE CASCADE;

ALTER TABLE "comments" ADD FOREIGN KEY ("post_id") REFERENCES "posts" ("id") ON DELETE CASCADE;