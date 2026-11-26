package database

import (
	"ai-dekawaf/model"
	"fmt"
)

func RunMigration() {
	var err error

	db := PostsqlConn()

	err = db.AutoMigrate(&model.SecurityProfile{})
	if err != nil {
		fmt.Printf("failed to migrate table security_profiles: %v\n", err)
	}

}

func RunRollback() {
	var err error

	db := PostsqlConn()

	err = db.Migrator().DropTable(&model.SecurityProfile{})
	if err != nil {
		fmt.Printf("failed to drop table security_profiles: %v\n", err)
	}

}
