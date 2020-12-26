.PHONY: bench, memstat

bench:
	go test -bench . -benchmem

memstat:
	@echo 'MEMSTAT STD MAP'
	@go run cmd/map/map.go
	@echo '-----------------'
	@echo 'MEMSTAT VELLUMFST'
	@go run cmd/vellumfst/vellumfst.go
	@echo '-----------------'
	@echo 'MEMSTAT PATRICIA'
	@go run cmd/patricia/patricia.go
	@echo '-----------------'
	@echo 'MEMSTAT SMHANOV_DAWG'
	@go run cmd/smhanov_dawg/smhanov_dawg.go
	@echo '-----------------'
	@echo 'MEMSTAT HASHICORP_IRADIX'
	@go run cmd/hashicorp_iradix/hashicorp_iradix.go
	@echo '-----------------'
	@echo 'MEMSTAT ARMON_RADIXTRIE'
	@go run cmd/armon_radixtrie/armon_radixtrie.go
	
	