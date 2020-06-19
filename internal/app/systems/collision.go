package systems

import (
	"github.com/andygeiss/ecs"
	"github.com/andygeiss/ecs-pong/internal/app/components"
	"github.com/gen2brain/raylib-go/raylib"
)

// Collision ...
type Collision struct {
	windowHeight int32
	windowWidth  int32
}

// NewCollision ...
func NewCollision(windowWidth, windowHeight int32) ecs.System {
	return &Collision{
		windowHeight: windowHeight,
		windowWidth:  windowWidth,
	}
}

// Process ...
func (s *Collision) Process(entityManager *ecs.EntityManager) (state int) {
	for _, e := range entityManager.FilterBy("position", "size", "velocity") {
		switch e.Id {
		case "ball":
			enemy := entityManager.Get("enemy")
			player := entityManager.Get("player")
			scoreboard := entityManager.Get("scoreboard")
			if s.hasCollisionWithEnemy(e, enemy) ||
				s.hasCollisionWithPlayer(e, player) ||
				s.hasCollisionWithWindowBottom(e) ||
				s.hasCollisionWithWindowTop(e) {
				s.handleCollisionSoundIfPresent(e)
			}
			s.handleEnemyScore(e, enemy, scoreboard)
			s.handlePlayerScore(e, player, scoreboard)
		case "enemy", "player":
			s.blockWindowBottom(e)
			s.blockWindowTop(e)
		}
	}
	return ecs.StateEngineContinue
}

// Setup ...
func (s *Collision) Setup() {}

// Teardown ...
func (s *Collision) Teardown() {}

func (s *Collision) blockWindowBottom(entity *ecs.Entity) {
	position := entity.Get("position").(*components.Position)
	size := entity.Get("size").(*components.Size)
	velocity := entity.Get("velocity").(*components.Velocity)
	if position.Y+velocity.Y+size.Height >= float32(s.windowHeight) {
		velocity.Y = 0
	}
}

func (s *Collision) blockWindowTop(entity *ecs.Entity) {
	position := entity.Get("position").(*components.Position)
	velocity := entity.Get("velocity").(*components.Velocity)
	if position.Y+velocity.Y <= 0 {
		velocity.Y = 0
	}
}

func (s *Collision) getEntityRect(entity *ecs.Entity) rl.Rectangle {
	position := entity.Get("position").(*components.Position)
	size := entity.Get("size").(*components.Size)
	return rl.NewRectangle(position.X, position.Y, size.Width, size.Height)
}

func (s *Collision) handleCollisionSoundIfPresent(ball *ecs.Entity) {
	sound := ball.Get("sound")
	if sound == nil {
		return
	}
	snd := sound.(*components.Sound)
	snd.Filename = snd.EventFilename["collision"]
}

func (s *Collision) hasCollisionWithEnemy(ball, enemy *ecs.Entity) (hasCollision bool) {
	ballRect := s.getEntityRect(ball)
	ballVelocity := ball.Get("velocity").(*components.Velocity)
	enemyRect := s.getEntityRect(enemy)
	enemyAI := enemy.Get("ai").(*components.AI)
	if rl.CheckCollisionRecs(ballRect, enemyRect) {
		ballVelocity.X *= -1
		if enemyAI.Down && ballVelocity.Y > 0 {
			ballVelocity.Y *= 2
		} else if enemyAI.Down && ballVelocity.Y < 0 {
			ballVelocity.Y *= 0.5
			ballVelocity.X *= 1.5
		} else if enemyAI.Up && ballVelocity.Y < 0 {
			ballVelocity.Y *= 2
		} else if enemyAI.Up && ballVelocity.Y > 0 {
			ballVelocity.Y *= 0.5
			ballVelocity.X *= 1.5
		}
		return true
	}
	return false
}

func (s *Collision) hasCollisionWithPlayer(ball, player *ecs.Entity) (hasCollision bool) {
	ballRect := s.getEntityRect(ball)
	ballVelocity := ball.Get("velocity").(*components.Velocity)
	playerRect := s.getEntityRect(player)
	playerInput := player.Get("input").(*components.Input)
	if rl.CheckCollisionRecs(ballRect, playerRect) {
		ballVelocity.X *= -1
		if playerInput.Down && ballVelocity.Y > 0 {
			ballVelocity.Y *= 2
		} else if playerInput.Down && ballVelocity.Y < 0 {
			ballVelocity.Y *= -0.5
			ballVelocity.X *= 1.5
		} else if playerInput.Up && ballVelocity.Y < 0 {
			ballVelocity.Y *= 2
		} else if playerInput.Up && ballVelocity.Y > 0 {
			ballVelocity.Y *= -0.5
			ballVelocity.X *= 1.5
		}
		return true
	}
	return false
}

func (s *Collision) hasCollisionWithWindowBottom(entity *ecs.Entity) (hasCollision bool) {
	position := entity.Get("position").(*components.Position)
	size := entity.Get("size").(*components.Size)
	velocity := entity.Get("velocity").(*components.Velocity)
	if position.Y+velocity.Y+size.Height >= float32(s.windowHeight) {
		velocity.Y *= -1
		return true
	}
	return false
}

func (s *Collision) hasCollisionWithWindowTop(entity *ecs.Entity) (hasCollision bool) {
	position := entity.Get("position").(*components.Position)
	velocity := entity.Get("velocity").(*components.Velocity)
	if position.Y+velocity.Y <= 0 {
		velocity.Y *= -1
		return true
	}
	return false
}

func (s *Collision) handleEnemyScore(ball, enemy, scoreboard *ecs.Entity) {
	position := ball.Get("position").(*components.Position)
	velocity := ball.Get("velocity").(*components.Velocity)
	score := scoreboard.Get("score").(*components.Score)
	if position.X+velocity.X <= 0 {
		score.Enemy++
		velocity.X = -3
		velocity.Y = 2
		position.X = float32(s.windowWidth) / 2
		position.Y = float32(s.windowHeight) / 2
	}
}

func (s *Collision) handlePlayerScore(ball, player, scoreboard *ecs.Entity) {
	position := ball.Get("position").(*components.Position)
	velocity := ball.Get("velocity").(*components.Velocity)
	score := scoreboard.Get("score").(*components.Score)
	if position.X+velocity.X >= float32(s.windowWidth) {
		score.Player++
		velocity.X = 3
		velocity.Y = -2
		position.X = float32(s.windowWidth) / 2
		position.Y = float32(s.windowHeight) / 2
	}
}
