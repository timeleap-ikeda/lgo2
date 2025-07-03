package main

import (
	"context"
	"log/slog"
	"os"
	"time"
	"fmt"
)

func main() {
	slog.Debug("debug log message")  //liststart1
	slog.Info("info log message")
	slog.Warn("warning log message")
	slog.Error("error log message")  //listend1

	fmt.Println()
	
	userID := "fred"  //liststart2
	loginCount := 20
	slog.Info("user login",
		"id", userID,
		"login_count", loginCount)  //listend2

	fmt.Println()	
	
	options := &slog.HandlerOptions{Level: slog.LevelDebug}  //liststart3
	handler := slog.NewJSONHandler(os.Stderr, options)
	mySlog := slog.New(handler)
	lastLogin := time.Date(2025, 01, 01, 11, 50, 00, 00, time.UTC)
	mySlog.Debug("debug message",
		"id", userID,
		"last_login", lastLogin)  //listend3

	ctx := context.Background()  //liststart4
	mySlog.LogAttrs(ctx, slog.LevelInfo, "faster logging", slog.String("id", userID), slog.Time("last_login", lastLogin))  //listend4

	myLog := slog.NewLogLogger(mySlog.Handler(), slog.LevelDebug)  //liststart5
	myLog.Println("using the mySlog Handler")  //listend5
}
