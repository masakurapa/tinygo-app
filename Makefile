init:
	brew tap tinygo-org/tools
	brew install go tinygo minicom
	go version
	tinygo version

flash_lchika:
	cd cmd/lchika && tinygo flash --target wioterminal
flash_marumarugame:
	cd cmd/marumarugame && tinygo flash --target wioterminal
flash_timergame:
	cd cmd/timergame && tinygo flash --target wioterminal
