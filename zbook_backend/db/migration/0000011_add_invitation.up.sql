CREATE TABLE invitations (
    "invitation_id" BIGSERIAL PRIMARY KEY,
    "email" VARCHAR(255) NOT NULL, -- 邀请邮箱
    "invitation_url" VARCHAR(255) NOT NULL UNIQUE, -- 随机字符串 URL
    "is_used" boolean NOT NULL DEFAULT FALSE,
    "created_at" timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "expired_at" timestamptz NOT NULL DEFAULT (CURRENT_TIMESTAMP + INTERVAL '10 minutes')
);

CREATE INDEX ON "invitations" ("invitation_url");
