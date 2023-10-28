# Makefile for PacAlias.
default:
	@echo 'Building PacAlias Version 0.1-alpha (HitoriGotoh)...'
	@go build -o pacalias
	@echo 'Done! Output binary is "pacalias".'
install:
	@echo 'Building PacAlias Version 0.1-alpha (HitoriGotoh)...'
	@go build -o pacalias
	@echo installing pacalias...
	@sudo cp pacalias /usr/bin
	@sudo chmod 755 /usr/bin/pacalias
	@sudo cp pacalias.conf /etc
	@sudo cp pacalias.service /etc/systemd/system
	@echo 'Done! Use "systemctl enable --now pacalias" to start and enable it.'
clean:
	@echo 'Cleaning...'
	@rm pacalias
	@echo 'Done!'