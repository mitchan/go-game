# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

This is a 2D game written in Go using the Ebiten game engine. The game features a player-controlled character, enemies (skeletons and pigs), tilemap-based world rendering, and a camera system that follows the player.

## Commands

### Build and Run
```bash
go build -v ./...
go run .
```

### Testing
```bash
go test -v ./...
```

### Pre-commit Hooks
The project uses lefthook for pre-commit hooks. Pre-commit runs:
- `go vet ./...` - Static analysis
- `go test -v ./...` - All tests
- `go build -v ./...` - Build verification

Install hooks: `lefthook install`
Run hooks manually: `lefthook run pre-commit`

## Architecture

### Game Loop Structure
The game follows Ebiten's standard game loop pattern implemented in `game.go`:
- `Update()` - Game logic, entity updates, camera positioning (60 TPS)
- `Draw()` - Render everything to screen
- `Layout()` - Define virtual screen dimensions with zoom factor

### Entity System
All game entities (Player, Skeleton, Pig) share a common `Sprite` base in `entity/sprite.go` that contains image data and position. Each entity type extends this with:
- Type-specific behavior in Update()
- Animation state machines mapping `entityState` (idle, up, down, left, right) to animations
- Custom Draw() methods that handle sprite animation frames and horizontal flipping

**Important**: All entities use `Vector2D` for velocity and follow the same animation pattern. Player has a `speed` multiplier applied to its velocity vector.

### Animation System (`animation/`)
Animations are frame-based and tick-driven:
- Define start/end frames, step size, and speed in TPS (ticks per second)
- Support horizontal/vertical flipping for directional sprites
- Each entity maintains a map of `entityState` to `Animation` instances
- The active animation is determined by movement direction

### Spritesheet System (`spritesheet/`)
The `SpriteSheet` struct defines sprite layout on tilesheets:
- `WidthInTiles`, `HeightInTiles` - Grid dimensions
- `TileSize` - Pixels per tile
- `Rect(frameIndex)` - Returns the image.Rectangle for a given frame

### Camera System (`camera.go`)
The camera implements a follow-cam with constraints:
- `FollowTarget()` - Centers camera on target position (typically player center)
- `Constrain()` - Prevents camera from showing areas outside the tilemap bounds
- Camera position is applied as a translation offset during Draw() calls

### Tilemap System (`tilemap.go`)
Tilemaps are loaded from JSON files (Tiled editor format):
- Supports multiple layers rendered in order
- Tile IDs in the `data` array map to tileset coordinates
- Assumes a 10-tile-wide tileset for coordinate calculation
- Map files located in `assets/maps/`

### Math Utilities (`math/`)
Contains `Vector2D` type for 2D coordinates. Note: Camera is cast to Vector2D when passing to Draw methods.

### Constants (`constants/`)
Centralized configuration:
- Window dimensions and zoom level
- `CellSize` - Standard tile/sprite size (32px)
- Asset paths for sprites

## Development Guidelines

### Adding New Entities
1. Create struct in `entity/` that embeds `*Sprite`
2. Define animation map with `entityState` keys
3. Implement Update() with movement/AI logic
4. Implement Draw() that uses active animation frame
5. Add to Game struct and initialize in NewGame()
6. Call Update() and Draw() in game loop

### Adding Animations
Sprite animations are frame indices on a spritesheet:
- Count frames left-to-right, top-to-bottom starting at 0
- For directional sprites, use `flipH: true` for left-facing animations
- Speed is in ticks before advancing frame (higher = slower)

### Working with the Map
- Edit maps in Tiled map editor, export as JSON to `assets/maps/`
- Map coordinates multiply by `CellSize` to get pixel coordinates
- Player bounds checking uses map dimensions: `mapWidth * CellSize`, `mapHeight * CellSize`
