package adminDb

type AdminDB struct {
}

// func (AdminDB) CreatePgParams(name string, username string, phoneNu int32, gender int16) string {
func (AdminDB) CreatePgParams() string {

	return `INSERT INTO pg_basic_models( name , user_name , phone_number , gender)
	VALUES( $1 , $2 , $3 , $4 )`
}
