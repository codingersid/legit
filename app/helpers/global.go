package helpers

import (
	"golang.org/x/crypto/bcrypt"
)

// HashedPassword menghasilkan hash dari password yang diberikan
func HashedPassword(password string) (string, error) {
	// Enkripsi password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		// Jika terjadi kesalahan, kembalikan error
		return "", err
	}

	// Ubah hashedPassword menjadi string sebelum dikirim
	return string(hashedPassword), nil
}

// CheckPassword memeriksa apakah password sesuai dengan hashed password
// Mengembalikan nil jika cocok, atau error jika tidak cocok
func CheckPassword(hashedPassword string, password string) error {
	// Membandingkan password asli dengan password yang di-hash
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		// Password tidak cocok
		return err
	}

	// Password cocok
	return nil
}
