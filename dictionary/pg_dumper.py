import psycopg2 # noqa
import csv
import os

from dotenv import load_dotenv

load_dotenv()
DB_HOST = os.getenv("DB_HOST", "localhost")
DB_PORT = os.getenv("DB_PORT", 5432)
DB_NAME = os.getenv("DB_NAME", "dictionary")
DB_USER = os.getenv("DB_USER", "your_username")
DB_PASSWORD = os.getenv("DB_PASSWORD", "your_password")

CSV_DIRS = "."
FILES = os.listdir(CSV_DIRS)
# Sort the files to ensure that the CSV files are imported in order
CSVS = sorted([f for f in FILES if f.endswith(".csv")])[:-1]

def import_csv_to_postgres():
    cursor, connection = None, None
    try:
        connection = psycopg2.connect(
            host=DB_HOST,
            port=DB_PORT,
            database=DB_NAME,
            user=DB_USER,
            password=DB_PASSWORD
        )
        cursor = connection.cursor()

        for CSV_FILE_PATH in CSVS:
            TABLE_NAME = CSV_FILE_PATH.split(".")[0]
            cursor.execute(f"DROP TABLE IF EXISTS {TABLE_NAME}")
            cursor.execute(
                f"CREATE TABLE IF NOT EXISTS {TABLE_NAME} (id SERIAL PRIMARY KEY, word TEXT, definition TEXT)"
            )
            with open(CSV_FILE_PATH, 'r') as csv_file:
                reader = csv.reader(csv_file)
                next(reader)  # Skip header
                row_count = 0
                for row in reader:
                    try:
                        insert_query = f"INSERT INTO {TABLE_NAME} (Word, Definition) VALUES (%s, %s)"
                        cursor.execute(insert_query, row)
                        row_count += 1
                        if row_count % 100 == 0:
                            connection.commit()
                            print(f"{row_count} rows inserted so far...")
                    except Exception as e:
                        print(f"Error with row {row_count}: {e}")
                connection.commit()
                print(f"Total rows inserted from {CSV_FILE_PATH}: {row_count}")
    except Exception as e:
        print(f"Error: {e}")
    finally:
        if cursor:
            cursor.close()
        if connection:
            connection.close()


if __name__ == '__main__':
    import_csv_to_postgres()