BEGIN TRANSACTION;
CREATE TABLE IF NOT EXISTS "users" (
	"id"	INTEGER,
    "username"	TEXT,
	"first_name"	TEXT,
	"last_name"	TEXT,
	"password" TEXT,
	"role"	TEXT,
	PRIMARY KEY("id" AUTOINCREMENT)
);
COMMIT;
