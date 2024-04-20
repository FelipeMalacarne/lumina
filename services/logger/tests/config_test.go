package config

import (
	"os"
	"testing"

	"github.com/felipemalacarne/lumina/logger/api/config"
)

func TestLoadConfig(t *testing.T) {
	// Set up test environment
	os.Setenv("MONGO_URI", "test-mongo-uri")
	os.Setenv("MONGO_DB_NAME", "test-mongo-db-name")
	os.Setenv("MONGO_ROOT_USERNAME", "test-mongo-root-username")
	os.Setenv("MONGO_ROOT_PASSWORD", "test-mongo-root-password")
	os.Setenv("RPC_PORT", ":5001")
	os.Setenv("GRPC_PORT", ":50051")

	// Call LoadConfig function
	config.LoadConfig()

	// Check that the configuration values were set correctly
	if config.MongoURI != "test-mongo-uri" {
		t.Errorf("MongoURI is not set correctly, got: %s", config.MongoURI)
	}
	if config.MongoDBName != "test-mongo-db-name" {
		t.Errorf("MongoDBName is not set correctly, got: %s", config.MongoDBName)
	}
	if config.MongoRootUsername != "test-mongo-root-username" {
		t.Errorf("MongoRootUsername is not set correctly, got: %s", config.MongoRootUsername)
	}
	if config.MongoRootPassword != "test-mongo-root-password" {
		t.Errorf("MongoRootPassword is not set correctly, got: %s", config.MongoRootPassword)
	}
	if config.RpcPort != "5001" {
		t.Errorf("RpcPort is not set correctly, got: %s", config.RpcPort)
	}
	if config.GrpcPort != ":50051" {
		t.Errorf("GrpcPort is not set correctly, got: %s", config.GrpcPort)
	}

	// Clean up test environment
	os.Unsetenv("MONGO_URI")
	os.Unsetenv("MONGO_DB_NAME")
	os.Unsetenv("MONGO_ROOT_USERNAME")
	os.Unsetenv("MONGO_ROOT_PASSWORD")
	os.Unsetenv("RPC_PORT")
	os.Unsetenv("GRPC_PORT")
}
