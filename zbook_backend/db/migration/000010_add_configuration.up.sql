CREATE TABLE configurations (
    "config_name" VARCHAR(255) PRIMARY KEY,
    "config_value" BOOLEAN NOT NULL,
    "created_at" TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- 允许注册&登录
INSERT INTO "configurations" ("config_name", "config_value")
VALUES ('allow_registration', true)
ON CONFLICT ("config_name") DO NOTHING;

INSERT INTO "configurations" ("config_name", "config_value")
VALUES ('allow_login', true)
ON CONFLICT ("config_name") DO NOTHING;

INSERT INTO "configurations" ("config_name", "config_value")
VALUES ('allow_invitation', true)
ON CONFLICT ("config_name") DO NOTHING;