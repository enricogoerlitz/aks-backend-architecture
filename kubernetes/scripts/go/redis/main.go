package main

import (
	"context"
	"fmt"
	"log"

	"github.com/redis/go-redis/v9"
)

func main() {
	log.Println("Starte Redis Cluster Client...")
	ctx := context.Background()
	// Redis Cluster-Adressen (anpassen je nach Setup)
	addrs := []string{
		"172.211.203.195:6379",
	}

	// Cluster-Client erstellen
	rdb := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:    addrs,
		Password: "my-password",
	})

	// Verbindung testen
	if err := rdb.Ping(ctx).Err(); err != nil {
		log.Fatalf("Fehler beim Verbinden mit dem Cluster: %v", err)
	}

	// INFO von einem Knoten abrufen
	info, err := rdb.Info(ctx, "all").Result()
	if err != nil {
		log.Fatalf("Fehler beim Abrufen von INFO: %v", err)
	}

	fmt.Println("===== Redis Cluster INFO =====")
	fmt.Println(info)
}
