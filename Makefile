generate: generateDB generateWire

generateDB:
	@go run cmd/gorm/generate.go
generateWire:
	@wire wires/wire.go
