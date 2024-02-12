package constants

type Role uint8

const (
	NONE    Role = 0
	PATIENT Role = 1

	GOD   Role = 100
	ADMIN   Role = 101
	DOCTOR   Role = 102
	RECEPTIONIST Role = 103
	TECHNICIAN  Role = 104
	PHARMACY Role = 105
)
