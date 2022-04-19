package migrations

import modelSql "anya-day/models/sql"

var Models = []any{
	&modelSql.User{},
	&modelSql.UserCredential{},
}
