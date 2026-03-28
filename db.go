package main

import (
	"log"
	"os"
	"time"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	
)

var db *gorm.DB

func realignVikingSequence() error {
	query := `SELECT setval(
		pg_get_serial_sequence('"Viking"', 'viking_id'),
		COALESCE((SELECT MAX(viking_id) FROM "Viking"), 0) + 1,
		false
	)`

	return db.Exec(query).Error
}

func initDB() {

	
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal("La variable d'environnement DATABASE_URL est requise")
	}

	var err error
	db, err = gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true,
	}), &gorm.Config{PrepareStmt: false})
	
	if err != nil {
		log.Fatalf("Impossible de se connecter à la base de données : %v", err)
	}

	if err := db.SetupJoinTable(&Event{}, "Vikings", &EventViking{}); err != nil {
		log.Fatalf("Impossible de configurer la table de jointure EventViking : %v", err)
	}

	if err := realignVikingSequence(); err != nil {
		log.Printf("⚠️ Impossible de realigner la sequence Viking: %v", err)
	} else {
		log.Printf("✓ Sequence Viking realignee")
	}


	// Récupère la connexion SQL sous-jacente
    sqlDB, err := db.DB()
    if err != nil {
        log.Printf("Erreur lors de la récupération de la connexion SQL : %v", err)
		return
    } 

    go func() {
		duration := 30 * time.Second
		timer := time.NewTimer(duration)

		for {
			<-timer.C

			if err := sqlDB.Ping(); err != nil {
				log.Printf("❌ Ping DB échoué: %v", err)
			} else {
				stats := sqlDB.Stats()
				log.Printf("✓ Connexion OK")
				log.Printf("✓ Max open connections: %d", stats.MaxOpenConnections)
				log.Printf("✓ Connexions actives: %d", stats.OpenConnections)
			}
			log.Println("Base de données connectée et migrée avec succès")

			// 🔁 reset du timer
			timer.Reset(duration)
		}
	}()

	
}

