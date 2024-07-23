CREATE TABLE configurations (
    config_name VARCHAR(255) PRIMARY KEY,
    config_value BOOLEAN NOT NULL
);

-- 允许注册&登录
INSERT INTO configurations (config_name, config_value) VALUES ('allow_registration', true);
INSERT INTO configurations (config_name, config_value) VALUES ('allow_login', true);