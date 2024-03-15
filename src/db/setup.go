package db

import "os"

// setup db credentials with environment variables
var db_user = os.Getenv("VISTA_VERSE_DB_USER")
var db_dbname = os.Getenv("VISTA_VERSE_DB_DBNAME")
var db_password = os.Getenv("VISTA_VERSE_DB_PASSWORD")
