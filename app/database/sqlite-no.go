// +build aix android dragonfly freebsd js illumos linux,ppc64 linux,riscv64 linux,mips linux,mips64 linux,s390x netbsd openbsd plan9 solaris windows,arm

package database

import (
	"go.uber.org/zap"
)

const SQLiteEnabled = false

func OpenSQLiteDatabase(params *SQLiteParams, logger *zap.SugaredLogger) (*Service, error) {
	return nil, nil
}