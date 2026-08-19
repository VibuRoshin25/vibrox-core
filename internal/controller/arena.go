package controller

import (
	"net/http"

	"vibrox-core/internal/config"
	arenapb "vibrox-core/internal/proto/arena"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type playMoveRequest struct {
	HumanMark string   `json:"humanMark" binding:"omitempty,oneof=X O"`
	Board     []string `json:"board" binding:"required,len=9"`
	Position  int32    `json:"position" binding:"min=0,max=8"`
}

// PlayArenaMove handles the POST request to play a move in the arena game. It validates the request, sends it to the Arena service, and returns the response.
func PlayArenaMove(c *gin.Context) {
	var request playMoveRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	response, err := config.ArenaClient.PlayMove(c, &arenapb.PlayMoveRequest{
		Board:     request.Board,
		Position:  request.Position,
		HumanMark: request.HumanMark,
	})
	if err != nil {
		if status.Code(err) == codes.InvalidArgument {
			c.JSON(http.StatusBadRequest, gin.H{"error": status.Convert(err).Message()})
			return
		}
		c.JSON(http.StatusBadGateway, gin.H{"error": "arena service unavailable"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"board":              response.GetBoard(),
		"botPosition":        response.GetBotPosition(),
		"outcome":            response.GetOutcome(),
		"strategy":           response.GetStrategy(),
		"score":              response.GetScore(),
		"nodesEvaluated":     response.GetNodesEvaluated(),
		"searchDepth":        response.GetSearchDepth(),
		"decisionTimeMicros": response.GetDecisionTimeMicros(),
	})
}
