package internal

import "os"

type Endpoint string

var (
	API   = getEnv("OMISE_API", "https://api.omise.co")
	Vault = getEnv("OMISE_VAULT", "https://vault.omise.co")
)

func getEnv(key, fallback string) Endpoint {
	value, exists := os.LookupEnv(key)
	if !exists {
		value = fallback
	}
	return Endpoint(value)
}
