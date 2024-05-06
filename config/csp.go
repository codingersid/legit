package config

import (
	"fmt"
	"strings"

	legitConfig "github.com/codingersid/legit-cli/config"
)

// Untuk mengaktifkan CSP, anda ubah status SEC_CSP pada .env menjadi true
// Grab Semua CSP
func ConfigCSP() string {
	env := legitConfig.LoadEnv()
	envCSO := env["SEC_CSP"]
	var joined string

	if envCSO == "true" {
		joined = fmt.Sprintf("default-src %s; script-src %s; style-src %s; frame-src %s; img-src %s; media-src %s; connect-src %s; manifest-src %s; worker-src %s; font-src %s;", defaultSrc(), scriptSrc(), styleSrc(), frameSrc(), imgSrc(), mediaSrc(), connectSrc(), manifestSrc(), workerSrc(), fontSrc())
	} else {
		joined = ""
	}
	return joined
}

// default-src
func defaultSrc() string {
	slice := []string{
		"'self'",
	}
	result := strings.Join(slice, " ")
	return result
}

// script-src
func scriptSrc() string {
	slice := []string{
		"'self'",
		"'unsafe-inline'",
		"'unsafe-eval'",
		"*.googleapis.com",
		"fonts.gstatic.com",
		"*.cloudflare.com",
		"*.youtube.com",
		"*.tau.ac.id",
	}
	result := strings.Join(slice, " ")
	return result
}

// style-src
func styleSrc() string {
	slice := []string{
		"'self'",
		"'unsafe-inline'",
		"'unsafe-eval'",
		"*.googleapis.com",
		"fonts.gstatic.com",
		"*.cloudflare.com",
		"*.youtube.com",
		"*.tau.ac.id",
	}
	result := strings.Join(slice, " ")
	return result
}

// frame-src
func frameSrc() string {
	slice := []string{
		"'self'",
		"'unsafe-inline'",
		"'unsafe-eval'",
		"*.googleapis.com",
		"fonts.gstatic.com",
		"*.cloudflare.com",
		"*.youtube.com",
		"*.tau.ac.id",
	}
	result := strings.Join(slice, " ")
	return result
}

// imgs-rc
func imgSrc() string {
	slice := []string{
		"'self'",
		"'unsafe-inline'",
		"'unsafe-eval'",
		"*.googleapis.com",
		"fonts.gstatic.com",
		"*.cloudflare.com",
		"*.youtube.com",
		"*.tau.ac.id",
	}
	result := strings.Join(slice, " ")
	return result
}

// media-src
func mediaSrc() string {
	slice := []string{
		"'self'",
		"'unsafe-inline'",
		"'unsafe-eval'",
		"*.googleapis.com",
		"fonts.gstatic.com",
		"*.cloudflare.com",
		"*.youtube.com",
		"*.tau.ac.id",
	}
	result := strings.Join(slice, " ")
	return result
}

// connect-src
func connectSrc() string {
	slice := []string{
		"'self'",
		"'unsafe-inline'",
		"'unsafe-eval'",
		"*.googleapis.com",
		"fonts.gstatic.com",
		"*.cloudflare.com",
		"*.youtube.com",
		"*.tau.ac.id",
	}
	result := strings.Join(slice, " ")
	return result
}

// manifest-src
func manifestSrc() string {
	slice := []string{
		"'self'",
		"'unsafe-inline'",
		"'unsafe-eval'",
		"*.googleapis.com",
		"fonts.gstatic.com",
		"*.cloudflare.com",
		"*.youtube.com",
		"*.tau.ac.id",
	}
	result := strings.Join(slice, " ")
	return result
}

// worker-src
func workerSrc() string {
	slice := []string{
		"'self'",
		"'unsafe-inline'",
		"'unsafe-eval'",
		"*.googleapis.com",
		"fonts.gstatic.com",
		"*.cloudflare.com",
		"*.youtube.com",
		"*.tau.ac.id",
	}
	result := strings.Join(slice, " ")
	return result
}

// font-src
func fontSrc() string {
	slice := []string{
		"'self'",
		"'unsafe-inline'",
		"'unsafe-eval'",
		"*.googleapis.com",
		"fonts.gstatic.com",
		"*.cloudflare.com",
		"*.youtube.com",
		"*.tau.ac.id",
	}
	result := strings.Join(slice, " ")
	return result
}
