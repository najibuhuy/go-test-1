package container

import (
	"fmt"
	"go-test-1/config"
	bankAccount_feature "go-test-1/domain/bank_account/feature"
	bankAccount_repository "go-test-1/domain/bank_account/repository"
	"go-test-1/infrastructure/database"
	"go-test-1/infrastructure/logger"
	"go-test-1/infrastructure/shared/constant"
	"log"
)

type Container struct {
	EnvironmentConfig  config.EnvironmentConfig
	BankAccountFeature bankAccount_feature.BankAccountFeature
}

func SetupContainer() Container {
	fmt.Println("Starting new container...")

	fmt.Println("Loading config...")
	config, err := config.LoadENVConfig()
	if err != nil {
		log.Panic(err)
	}

	logger.InitializeLogger(constant.LOGRUS) // choose which log, ZAP or LOGRUS. Default: LOGRUS

	fmt.Println("Loading database...")
	db, err := database.LoadDatabase()
	if err != nil {
		log.Panic(err)
	}

	fmt.Println("Loading repository's...")
	bankAccountRepository := bankAccount_repository.NewBankAccountRepository(db)
	bankAccountFeature := bankAccount_feature.NewBankAccountFeature(bankAccountRepository)

	return Container{
		EnvironmentConfig:  config,
		BankAccountFeature: bankAccountFeature,
	}
}
