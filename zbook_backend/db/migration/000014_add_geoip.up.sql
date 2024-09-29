CREATE TABLE IF NOT EXISTS "geoip" (
    geoip_id bigserial PRIMARY KEY,
    ip_range_cidr cidr,                    -- IP 地址范围（CIDR 格式）
    city_name_en TEXT,                     -- 城市名称（英语）
    city_name_zh_cn TEXT,                  -- 城市名称（中文-简体）
    latitude DOUBLE PRECISION,             -- 纬度
    longitude DOUBLE PRECISION,            -- 经度
    UNIQUE (ip_range_cidr)                 -- 在创建唯一约束时，指定 ip_range_cidr 列为唯一
);

-- 如果需要 GIST 索引来提高查询性能，可以创建非唯一 GIST 索引
CREATE INDEX geoip_cidr_gist_idx ON geoip USING GIST (ip_range_cidr inet_ops);