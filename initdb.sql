CREATE TABLE books (
		id   TEXT,
		name         TEXT,
		author         TEXT,
		translator     TEXT,
        publisher TEXT,
        print_year TEXT,
		date_updated TIMESTAMP with time zone DEFAULT now(),
	
		PRIMARY KEY (id)
);

INSERT INTO books (
	id,
	name,
	author,
	translator,
	publisher,
	print_year,
	date_updated
	)
VALUES(
	'978-616-553-903-6',
	'อินเทอร์เน็ตเพื่องานธุรกิจ',
	'สุนทรีย์ โพธิ์อิ่ม, ไมตรี ฉลาดธรรม',
	'',
	'สำนักพิมพ์ศูนย์ส่งเสริมอาชีวะ',
	'2562',
	TIMESTAMP '2020-01-28 09:12:00'
	);
