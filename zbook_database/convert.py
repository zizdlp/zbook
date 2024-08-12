import maxminddb
import psycopg2
from ipaddress import ip_network, AddressValueError
import logging
from concurrent.futures import ThreadPoolExecutor, as_completed

# 设置日志记录
logging.basicConfig(level=logging.INFO)
logger = logging.getLogger(__name__)

def process_batch(batch_data, conn_params):
    try:
        with psycopg2.connect(**conn_params) as conn:
            with conn.cursor() as cur:
                cur.executemany("""
                    INSERT INTO geoip (
                        ip_range_cidr, city_name_en, city_name_zh_cn,
                        latitude, longitude
                    )
                    VALUES (%s, %s, %s, %s, %s)
                    ON CONFLICT (ip_range_cidr) DO NOTHING;
                """, batch_data)
            conn.commit()
    except psycopg2.DatabaseError as e:
        logger.error(f"DatabaseError: {e}")
    except Exception as e:
        logger.error(f"Unexpected error during batch processing: {e}")

def main():
    conn_params = {
        'dbname': 'zbook',
        'user': 'root',
        'password': 'secret',
        'host': 'localhost'
    }

    batch_size = 1000  # 每批次插入1000条数据
    batch_data = []

    try:
        # 打开 MMDB 文件
        with maxminddb.open_database('GeoLite2-City.mmdb') as reader:
            with ThreadPoolExecutor() as executor:
                futures = []
                
                # 遍历所有 IP 地址段
                for ip_range, data in reader:
                    try:
                        # 将 IP 范围转换为字符串（CIDR 格式）
                        ip_range_cidr = str(ip_network(ip_range))

                        # 提取各语言的城市名称
                        city_names = data.get('city', {}).get('names', {})
                        city_name_en = city_names.get('en', None)
                        city_name_zh_cn = city_names.get('zh-CN', None)

                        # 提取地理位置信息
                        latitude = data.get('location', {}).get('latitude', None)
                        longitude = data.get('location', {}).get('longitude', None)

                        # 将数据添加到批处理列表中
                        batch_data.append((
                            ip_range_cidr, city_name_en, city_name_zh_cn,
                            latitude, longitude
                        ))

                        # 如果批处理列表达到指定大小，则提交线程池任务
                        if len(batch_data) >= batch_size:
                            futures.append(executor.submit(process_batch, batch_data, conn_params))
                            batch_data = []  # 清空批处理列表

                    except AddressValueError as e:
                        logger.error(f"AddressValueError: {e} for IP range: {ip_range}")
                    except Exception as e:
                        logger.error(f"Error processing IP range {ip_range}: {e}")

                # 插入剩余的数据
                if batch_data:
                    futures.append(executor.submit(process_batch, batch_data, conn_params))

                # 等待所有线程完成
                for future in as_completed(futures):
                    future.result()

    except Exception as e:
        logger.error(f"Unexpected error: {e}")

if __name__ == "__main__":
    main()