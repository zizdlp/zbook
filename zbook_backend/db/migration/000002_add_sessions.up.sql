CREATE TABLE "sessions" (
  "session_id" uuid PRIMARY KEY,
  "user_id" bigint NOT NULL,
  "refresh_token" varchar NOT NULL,
  "user_agent" varchar NOT NULL,
  "client_ip" varchar NOT NULL,
  "expires_at" timestamptz NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX ON "sessions" ("user_id");
CREATE INDEX ON "sessions" ("refresh_token");
CREATE INDEX ON "sessions" ("expires_at");
