package auth

import "os"

var jwt_secret = os.Getenv("VISTA_VERSE_JWT_SECRET")
