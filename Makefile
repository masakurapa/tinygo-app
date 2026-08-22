init:
	brew tap tinygo-org/tools
	brew install go tinygo minicom
	go version
	tinygo version

#
# games
#
_flash:
	tinygo flash --size short --target wioterminal ./apps/$(GAME_NAME)

flash_timer:
	GAME_NAME=timer $(MAKE) _flash

flash_dropcircle:
	GAME_NAME=dropcircle $(MAKE) _flash

flash_poyopoyo:
	GAME_NAME=poyopoyo $(MAKE) _flash
