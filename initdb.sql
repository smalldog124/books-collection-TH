CREATE TABLE books (
		id   SERIAL PRIMARY KEY,
		ISBN TEXT,
		name         TEXT,
		writer         TEXT,
		translator     TEXT,
        publisher TEXT,
        print_year TEXT,
		image_uri TEXT,
		date_updated TIMESTAMP with time zone DEFAULT now()
);

CREATE TABLE shelf (
	id SERIAL PRIMARY KEY,
	user_id integer,
	book_id integer,
	score integer,
	date_created TIMESTAMP with time zone DEFAULT now()
);

CREATE TABLE wish_list (
	book_id integer,
	user_id integer,
	date_created TIMESTAMP with time zone DEFAULT now()
);

CREATE TABLE review (
	book_id integer,
	score integer,
	date_created TIMESTAMP with time zone DEFAULT now()
);

INSERT INTO books (
	ISBN,
	name,
	writer,
	translator,
	publisher,
	print_year,
	date_updated
	)
VALUES(
	'978-616-18-2996-4',
	'ทำไม Netflix ถึงมีแต่คนโตครเก่ง',
	'แพตตี้ แมคคอร์ด',
	'วิกันดา จันทร์ทองสุข',
	'บริษัทอมรินทร์พริ้นติ้งแอนด์พับลิชซิ่ง จำกัด (มหาชน)',
	'2558',
	TIMESTAMP '2020-01-28 09:12:00'
);

INSERT INTO books (
	ISBN,
	name,
	writer,
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

INSERT INTO wish_list (
	book_id,
	user_id,
	date_created
	)
VALUES(
	2,
	137499732,
	TIMESTAMP '2020-01-28 09:12:00'
);