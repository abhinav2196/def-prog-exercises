package legacyconversions

import (
	"blocksize-capital/golab-conference/testing-app/def-prog-exercises/safesql"
	"blocksize-capital/golab-conference/testing-app/def-prog-exercises/safesql/internal/raw"
)

var trustedSQLCtor = raw.TrustedSQLCtor.(func(string) safesql.TrustedSQL)

func RiskilyAssumeTrustedSQL(trusted string) safesql.TrustedSQL {
	return trustedSQLCtor(trusted)
}
