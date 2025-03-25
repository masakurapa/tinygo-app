init:
	brew tap tinygo-org/tools
	brew install go tinygo minicom
	go version
	tinygo version

flash_lchika:
	cd cmd/lchika && tinygo flash --target wioterminal

# games
flash_marumarugame:
	cd cmd/game/marumaru && tinygo flash --target wioterminal
flash_timergame:
	cd cmd/game/timer && tinygo flash --target wioterminal
