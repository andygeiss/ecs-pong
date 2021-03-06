package systems

import (
	"github.com/andygeiss/ecs"
	"github.com/andygeiss/ecs-pong/internal/app/components"
)

// collision ...
type collision struct {
	windowHeight int32
	windowWidth  int32
}

// NewCollision ...
func NewCollision(windowWidth, windowHeight int32) ecs.System {
	return &collision{
		windowHeight: windowHeight,
		windowWidth:  windowWidth,
	}
}

// Process ...
func (s *collision) Process(entityManager *ecs.EntityManager) (state int) {
	for _, e := range entityManager.FilterByMask(components.MaskPosition | components.MaskSize | components.MaskVelocity) {
		switch e.ID() {
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
func (s *collision) Setup() {}

// Teardown ...
func (s *collision) Teardown() {}

func (s *collision) blockWindowBottom(entity *ecs.Entity) {
	position := entity.Get(components.MaskPosition).(*components.Position)
	size := entity.Get(components.MaskSize).(*components.Size)
	velocity := entity.Get(components.MaskVelocity).(*components.Velocity)
	if position.Y+velocity.Y+size.Height >= float32(s.windowHeight) {
		velocity.Y = 0
	}
}

func (s *collision) blockWindowTop(entity *ecs.Entity) {
	position := entity.Get(components.MaskPosition).(*components.Position)
	velocity := entity.Get(components.MaskVelocity).(*components.Velocity)
	if position.Y+velocity.Y <= 0 {
		velocity.Y = 0
	}
}

func (s *collision) handleCollisionSoundIfPresent(ball *ecs.Entity) {
	sound := ball.Get(components.MaskSound)
	if sound == nil {
		return
	}
	snd := sound.(*components.Sound)
	snd.Filename = snd.EventFilename["collision"]
}

func (s *collision) hasCollisionWithEnemy(ball, enemy *ecs.Entity) (hasCollision bool) {
	ballPos := ball.Get(components.MaskPosition).(*components.Position)
	ballSize := ball.Get(components.MaskSize).(*components.Size)
	ballVelocity := ball.Get(components.MaskVelocity).(*components.Velocity)
	enemyPos := enemy.Get(components.MaskPosition).(*components.Position)
	enemySize := enemy.Get(components.MaskSize).(*components.Size)
	enemyAI := enemy.Get(components.MaskAI).(*components.AI)
	if ballPos.X+ballSize.Width >= enemyPos.X &&
		ballPos.Y >= enemyPos.Y &&
		ballPos.Y+ballSize.Height <= enemyPos.Y+enemySize.Height {
		ballVelocity.X *= -1
		if enemyAI.Down && ballVelocity.Y > 0 {
			ballVelocity.Y *= 1.5
		} else if enemyAI.Down && ballVelocity.Y < 0 {
			ballVelocity.Y *= -0.75
			ballVelocity.X *= 1.25
		} else if enemyAI.Up && ballVelocity.Y < 0 {
			ballVelocity.Y *= 1.5
		} else if enemyAI.Up && ballVelocity.Y > 0 {
			ballVelocity.Y *= -0.75
			ballVelocity.X *= 1.25
		}
		return true
	}
	return false
}

func (s *collision) hasCollisionWithPlayer(ball, player *ecs.Entity) (hasCollision bool) {
	ballPos := ball.Get(components.MaskPosition).(*components.Position)
	ballSize := ball.Get(components.MaskSize).(*components.Size)
	ballVelocity := ball.Get(components.MaskVelocity).(*components.Velocity)
	playerPos := player.Get(components.MaskPosition).(*components.Position)
	playerSize := player.Get(components.MaskSize).(*components.Size)
	playerInput := player.Get(components.MaskInput).(*components.Input)
	if ballPos.X <= playerPos.X+playerSize.Width &&
		ballPos.Y >= playerPos.Y &&
		ballPos.Y+ballSize.Height <= playerPos.Y+playerSize.Height {
		ballVelocity.X *= -1
		if playerInput.Down && ballVelocity.Y > 0 {
			ballVelocity.Y *= 1.5
		} else if playerInput.Down && ballVelocity.Y < 0 {
			ballVelocity.Y *= -0.75
			ballVelocity.X *= 1.25
		} else if playerInput.Up && ballVelocity.Y < 0 {
			ballVelocity.Y *= 1.5
		} else if playerInput.Up && ballVelocity.Y > 0 {
			ballVelocity.Y *= -0.75
			ballVelocity.X *= 1.25
		}
		return true
	}
	return false
}

func (s *collision) hasCollisionWithWindowBottom(entity *ecs.Entity) (hasCollision bool) {
	position := entity.Get(components.MaskPosition).(*components.Position)
	size := entity.Get(components.MaskSize).(*components.Size)
	velocity := entity.Get(components.MaskVelocity).(*components.Velocity)
	if position.Y+velocity.Y+size.Height >= float32(s.windowHeight) {
		velocity.Y *= -1
		return true
	}
	return false
}

func (s *collision) hasCollisionWithWindowTop(entity *ecs.Entity) (hasCollision bool) {
	position := entity.Get(components.MaskPosition).(*components.Position)
	velocity := entity.Get(components.MaskVelocity).(*components.Velocity)
	if position.Y+velocity.Y <= 0 {
		velocity.Y *= -1
		return true
	}
	return false
}

func (s *collision) handleEnemyScore(ball, enemy, scoreboard *ecs.Entity) {
	position := ball.Get(components.MaskPosition).(*components.Position)
	velocity := ball.Get(components.MaskVelocity).(*components.Velocity)
	score := scoreboard.Get(components.MaskScore).(*components.Score)
	if position.X+velocity.X <= 0 {
		score.Enemy++
		velocity.X = -3
		velocity.Y = 2
		position.X = float32(s.windowWidth) / 2
		position.Y = float32(s.windowHeight) / 2
	}
}

func (s *collision) handlePlayerScore(ball, player, scoreboard *ecs.Entity) {
	position := ball.Get(components.MaskPosition).(*components.Position)
	velocity := ball.Get(components.MaskVelocity).(*components.Velocity)
	score := scoreboard.Get(components.MaskScore).(*components.Score)
	if position.X+velocity.X >= float32(s.windowWidth) {
		score.Player++
		velocity.X = 3
		velocity.Y = -2
		position.X = float32(s.windowWidth) / 2
		position.Y = float32(s.windowHeight) / 2
	}
}
