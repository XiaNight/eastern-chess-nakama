package main

import (
	"context"
	"database/sql"
	"encoding/json"

	"github.com/heroiclabs/nakama-common/runtime"
)

//noinspection GoUnusedExportedFunction
func InitModule(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, initializer runtime.Initializer) error {
	
	// if err := initializer.RegisterMatch("lobby", RegisterMatch); err != nil {
	// 	logger.Error("unable to register: %v", err)
	// 	return err
	// }

	err := initializer.RegisterRpc("healthcheck", RpcHealthcheck)
	if err != nil{
		return err
	}

	// err2 := initializer.RegisterRpc("healthcheck2", RpcHealthcheck)
	// if err2 != nil{
	// 	return err2
	// }

	// err3 := initializer.RegisterBeforeRt("MatchmakerAdd", MatchmakerAdd)
	// if err3 != nil{
	// 	return err3
	// }

	logger.Info("Eastern Chess engine loaded.")

	return nil
}

type HealthcheckResponse struct {
	Success bool `json: "success"`
}

func RpcHealthcheck(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, payload string) (string, error){
	logger.Debug("Healthcheck RPC called")
	response := &HealthcheckResponse{ Success: true }

	out, err := json.Marshal(response)
	if err != nil{
		logger.Error("Error marshalling response type to JSON: %v", err)
		return "", runtime.NewError("Cannon marshal type", 13)
	}

	return string(out), nil
}