package internal_api

import (
	"context"
	"f1betting/betting_system"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/jackc/pgx/v5"
)

// RateLimiter for basic rate limiting
type RateLimiter struct {
	requests map[string][]time.Time
	mutex    sync.Mutex
}

func NewRateLimiter() *RateLimiter {
	return &RateLimiter{
		requests: make(map[string][]time.Time),
	}
}

var limiter = NewRateLimiter()

// AuthMiddleware wraps an http.HandlerFunc and adds security checks
func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Enforce HTTPS
		// if r.TLS == nil {
		// 	http.Redirect(w, r, "https://"+r.Host+r.RequestURI, http.StatusMovedPermanently)
		// 	return
		// }

		// Rate limiting
		limiter.mutex.Lock()
		now := time.Now()
		clientIP := r.Header.Get("X-Real-IP")
		if clientIP == "" {
			clientIP = r.RemoteAddr
		}

		// Clean old requests
		reqs := limiter.requests[clientIP]
		for i, t := range reqs {
			if now.Sub(t) > time.Minute {
				reqs = reqs[i:]
				break
			}
		}

		// Check rate limit (5 requests per minute)
		if len(reqs) >= 5 {
			limiter.mutex.Unlock()
			http.Error(w, "Rate limit exceeded", http.StatusTooManyRequests)
			return
		}

		limiter.requests[clientIP] = append(reqs, now)
		limiter.mutex.Unlock()

		// API key check
		apiKey := r.Header.Get("X-API-Key")
		expectedAPIKey := "surya" //os.Getenv("API_KEY")
		if apiKey == "" || apiKey != expectedAPIKey {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(`{"error": "unauthorized"}`))
			return
		}

		// Check for origin IP
		// originIP := r.Header.Get("X-Forwarded-For")
		// if originIP == "" {
		// 	originIP = r.RemoteAddr
		// }
		// if originIP != "expected_origin_ip" {
		// 	w.Header().Set("Content-Type", "application/json")
		// 	w.WriteHeader(http.StatusForbidden)
		// 	w.Write([]byte(`{"error": "forbidden"}`))
		// 	return
		// }

		// Security headers
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.Header().Set("X-Frame-Options", "DENY")
		w.Header().Set("X-XSS-Protection", "1; mode=block")
		w.Header().Set("Strict-Transport-Security", "max-age=31536000; includeSubDomains")
		w.Header().Set("Content-Security-Policy", "default-src 'self'")
		w.Header().Set("Referrer-Policy", "strict-origin-when-cross-origin")
		w.Header().Set("Cache-Control", "no-store")
		w.Header().Set("Pragma", "no-cache")

		next(w, r)
	}
}

type InternalAPI struct {
	conn *pgx.Conn
	ctx  context.Context
}

func (I *InternalAPI) settleFastesLapBetsHandler(w http.ResponseWriter, r *http.Request) {
	// Settle fastest lap bets

	err := betting_system.SettleBetsAndUpdateBalanceFastestLap(I.ctx, I.conn, 9165)

	if err != nil {
		log.Fatalf("Failed to settle bets: %v", err)
	}

}

func RegisterInternalAPIs(conn *pgx.Conn) {

	// Settle fastest lap bets
	api := &InternalAPI{
		conn: conn,
		ctx:  context.Background(),
	}
	http.HandleFunc("/internal/settle_fastest_lap_bets", AuthMiddleware(api.settleFastesLapBetsHandler))

}
