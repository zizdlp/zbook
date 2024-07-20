---- 系统通知，通知指定用户
CREATE TABLE "system_notifications" (
  "noti_id" bigserial PRIMARY KEY,
  "user_id" bigint NOT NULL,
  "title" varchar NOT NULL,
  "contents" varchar NOT NULL,
  "redirect_url" varchar DEFAULT '',
  "readed" boolean NOT NULL DEFAULT 'false',
  "created_at"  timestamptz NOT NULL DEFAULT (now())
);

-- 添加外键约束和索引
ALTER TABLE "system_notifications"
 ADD FOREIGN KEY ("user_id") REFERENCES "users" ("user_id")  ON DELETE CASCADE;

CREATE INDEX ON "system_notifications" ("noti_id", "user_id");  --- 添加索引，加快检索

---- new follower 通知
CREATE TABLE "follower_notifications" (
  "noti_id" bigserial PRIMARY KEY,
  "user_id" bigint NOT NULL,
  "follower_id" bigint NOT NULL,
  "readed" boolean NOT NULL DEFAULT 'false',
  "created_at"  timestamptz NOT NULL DEFAULT (now())
);

-- 添加外键约束和索引
ALTER TABLE "follower_notifications"
  ADD FOREIGN KEY ("user_id") REFERENCES "users" ("user_id") ON DELETE CASCADE,
  ADD FOREIGN KEY ("follower_id") REFERENCES "users" ("user_id") ON DELETE CASCADE;

CREATE UNIQUE INDEX ON "follower_notifications" ("user_id","follower_id"); --- 添加唯一索引，防止重复
CREATE INDEX ON "follower_notifications" ("noti_id","user_id");  --- 添加索引，加快检索

---- following's repo create
CREATE TABLE "repo_notifications" (
  "noti_id" bigserial PRIMARY KEY,
  "user_id" bigint NOT NULL,
  "repo_id" bigint NOT NULL,
  "readed" boolean NOT NULL DEFAULT 'false',
  "created_at"  timestamptz NOT NULL DEFAULT (now())
);

-- 添加外键约束和索引
ALTER TABLE "repo_notifications"
  ADD FOREIGN KEY ("user_id") REFERENCES "users" ("user_id") ON DELETE CASCADE,
  ADD FOREIGN KEY ("repo_id") REFERENCES "repos" ("repo_id") ON DELETE CASCADE;

CREATE UNIQUE INDEX ON "repo_notifications" ("user_id","repo_id"); --- 添加唯一索引，防止重复
CREATE INDEX ON "repo_notifications" ("noti_id","user_id");  --- 添加索引，加快检索

---- new comment 通知
CREATE TABLE "comment_notifications" (
  "noti_id" bigserial PRIMARY KEY,
  "user_id" bigint NOT NULL,
  "comment_id" bigint Not NULL,
  "readed" boolean NOT NULL DEFAULT 'false',
  "created_at"  timestamptz NOT NULL DEFAULT (now())
);

-- 添加外键约束和索引
ALTER TABLE "comment_notifications"
  ADD FOREIGN KEY ("user_id") REFERENCES "users" ("user_id") ON DELETE CASCADE,
  ADD FOREIGN KEY ("comment_id") REFERENCES "comments" ("comment_id") ON DELETE CASCADE;

CREATE UNIQUE INDEX ON "comment_notifications" ("user_id","comment_id"); --- 添加唯一索引，防止重复
CREATE INDEX ON "comment_notifications" ("noti_id","user_id");  --- 添加索引，加快检索