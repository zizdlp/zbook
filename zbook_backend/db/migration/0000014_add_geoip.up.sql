CREATE TABLE IF NOT EXISTS "geoip" (
    ip_range_cidr cidr PRIMARY KEY,        -- IP 地址范围（CIDR 格式）
    city_name_en TEXT,                     -- 城市名称（英语）
    city_name_zh_cn TEXT,                  -- 城市名称（中文-简体）
    latitude DOUBLE PRECISION,             -- 纬度
    longitude DOUBLE PRECISION             -- 经度
);