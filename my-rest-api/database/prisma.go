package database

import (
	"my-rest-api/prisma/db"
	"log"
)

var PrismaClient *db.PrismaClient

func InitializePrisma() {
	PrismaClient = db.NewClient()
	err := PrismaClient.Connect()
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
}