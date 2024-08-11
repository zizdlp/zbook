import maxminddb
import psycopg2
from ipaddress import ip_network

# 连接到 PostgreSQL 数据库
conn = psycopg2.connect(
    dbname="zbook", user="root", password="secret", host="localhost"
)
cur = conn.cursor()

# 打开 MMDB 文件
with maxminddb.open_database('GeoLite2-City.mmdb') as reader:
    # 遍历所有 IP 地址段
    for ip_range, data in reader:
        # 将 IP 范围转换为字符串（CIDR 格式）
        ip_range_cidr = str(ip_network(ip_range))
        
        # 提取各语言的城市名称
        city_names = data.get('city', {}).get('names', {})
        city_name_en = city_names.get('en', None)
        city_name_zh_cn = city_names.get('zh-CN', None)
        
        # 提取地理位置信息
        latitude = data.get('location', {}).get('latitude', None)
        longitude = data.get('location', {}).get('longitude', None)

        # 将数据插入 PostgreSQL
        cur.execute("""
            INSERT INTO geoip (
                ip_range_cidr, city_name_en, city_name_zh_cn,
                latitude, longitude
            )
            VALUES (%s, %s, %s, %s, %s)
            ON CONFLICT (ip_range_cidr) DO NOTHING;
        """, (
            ip_range_cidr, city_name_en, city_name_zh_cn,latitude, longitude
        ))

# 提交并关闭数据库连接
conn.commit()
cur.close()
conn.close()
