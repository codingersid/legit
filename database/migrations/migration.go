package migrations

import legitConfig "github.com/codingersid/legit-cli/config"

func RunMigration() {
	// inisialisasi database
	db := legitConfig.DB
	// call migration
	Sessions(db)
	Caches(db)
}
