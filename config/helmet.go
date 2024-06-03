package config

import "github.com/gofiber/fiber/v2/middleware/helmet"

// Inisialisasi penyimpanan helmet
var ConfigHelmet = helmet.Config{
	ContentSecurityPolicy:     ConfigCSP(),                          // Sebaiknya konfigurasikan CSP secara ketat.
	XSSProtection:             "1; mode=block",                      // Mengaktifkan dan memblokir serangan XSS.
	ContentTypeNosniff:        "nosniff",                            // Mencegah browser menebak tipe MIME.
	XFrameOptions:             "DENY",                               // Mencegah halaman ditampilkan dalam frame.
	ReferrerPolicy:            "no-referrer",                        // Tidak mengirimkan header Referrer.
	CrossOriginEmbedderPolicy: "require-corp; report-to='default';", // (require-corp: Hanya izinkan dokumen dengan kebijakan yang benar). (report-to='default';: Jika ada pelanggaran, akan dilaporkan ke endpoint default). (unsafe-none: Untuk mengizinkan dokumen dan sumber daya dari berbagai asal)
	CrossOriginOpenerPolicy:   "same-origin-allow-popups",           // (same-origin: Isolasi data dengan kebijakan asal yang sama). (same-origin-allow-popups: Untuk mengizinkan pop-up dari asal yang sama)
	CrossOriginResourcePolicy: "cross-origin",                       // (same-origin: Batasi sumber daya hanya untuk dokumen dengan asal yang sama). (cross-origin: Untuk mengizinkan sumber daya dari berbagai asal)
	OriginAgentCluster:        "?1",                                 // Mengaktifkan Origin-Agent-Cluster.
	XDNSPrefetchControl:       "off",                                // Mencegah prefetching DNS.
	XDownloadOptions:          "noopen",                             // Mencegah file dibuka langsung di browser.
	XPermittedCrossDomain:     "none",                               // Batasi berbagi data lintas domain.
}
