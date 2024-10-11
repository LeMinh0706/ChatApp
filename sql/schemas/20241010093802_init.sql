-- +goose Up
-- +goose StatementBegin
CREATE TABLE "users" (
  "id" bigserial PRIMARY KEY,
  "username" varchar UNIQUE NOT NULL,
  "password" varchar NOT NULL,
  "url_avatar" varchar NOT NULL
);

CREATE TABLE "message" (
  "id" bigserial PRIMARY KEY,
  "from_id" bigint NOT NULL,
  "to_id" bigint NOT NULL,
  "content" varchar NOT NULL,
  "date_created" timestamptz NOT NULL DEFAULT (now())
);

CREATE INDEX ON "message" ("from_id");

CREATE INDEX ON "message" ("to_id");

ALTER TABLE "message" ADD FOREIGN KEY ("from_id") REFERENCES "users" ("id");

ALTER TABLE "message" ADD FOREIGN KEY ("to_id") REFERENCES "users" ("id");

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE "message" DROP CONSTRAINT IF EXISTS "message_from_id_fkey";
ALTER TABLE "message" DROP CONSTRAINT IF EXISTS "message_to_id_fkey";

DROP TABLE IF EXISTS "users";
DROP TABLE IF EXISTS "message";
-- +goose StatementEnd
