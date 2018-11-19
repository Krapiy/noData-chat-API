package fixtures

import (
	"github.com/Krapiy/noData-chat-API/db"
	"github.com/pkg/errors"
	"gopkg.in/testfixtures.v2"
)

// Load write fixtures form './data' to database
func Load(db *db.MysqlDB) error {
	fixtures, err := testfixtures.NewFolder(db.Conn.DB, &testfixtures.MySQL{}, "db/fixtures/data")
	if err != nil {
		return errors.Wrap(err, "cannot prepare fixtures")
	}

	testfixtures.SkipDatabaseNameCheck(true)

	err = fixtures.Load()
	if err != nil {
		return errors.Wrap(err, "cannot load fixtures")
	}

	return nil
}
