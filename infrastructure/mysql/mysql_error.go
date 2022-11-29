package mysql

import (
	coreerror "daijoubuteam.xyz/se214-library-management/core/error"
	"github.com/go-sql-driver/mysql"
)

func DriverErrorHandling(err *mysql.MySQLError) error {

	if err.Number == ER_ROW_IS_REFERENCED || err.Number == ER_ROW_IS_REFERENCED_2 {
		return coreerror.NewConflictError("Unable to delete because of conflict with server state", err)
	}
	return coreerror.NewInternalServerError("mysql error"+err.Message, err)
}
