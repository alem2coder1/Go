package Model

type Cache struct {
	CacheKey       string `gorm:"primaryKey;column:cacheKey"`
	CacheValue     string `gorm:"column:cacheValue"`
	ExpirationTime int    `gorm:"column:expirationTime"`
}
