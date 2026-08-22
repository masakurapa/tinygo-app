# CLAUDE.md - Guide for TinyGo App Development

## Build Commands
- `make init` - Set up environment (install tinygo, dependencies)
- `make build_wasm_timer` - Build timer game as WebAssembly
- `make flash_timer` - Flash timer game to Wio Terminal
- `make build_wasm_dropcircle` - Build dropcircle game as WebAssembly
- `make flash_dropcircle` - Flash dropcircle game to Wio Terminal

## Code Style Guidelines
- **Imports**: Group standard library first, then external packages
- **Formatting**: Use gofmt/go fmt
- **Types**: Declare at package level with const blocks for related constants
- **Naming**: 
  - camelCase for variables and functions
  - structs use lowercase names
  - constants use camelCase with descriptive prefixes
- **Error Handling**: Simple return nil pattern
- **Project Structure**:
  - `games/[game_name]/game/game.go` - Game logic implementation
  - `games/[game_name]/main.go` - Entry points
  - `internal/` - Shared utilities

## Technology Stack
- Go 1.24.1
- TinyGo 0.37.0
- koebiten framework for game development
- Targets Wio Terminal hardware