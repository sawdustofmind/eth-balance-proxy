package server

import "github.com/gin-gonic/gin"

func (s *ServerImpl) GetEthBalanceAddress(c *gin.Context, address string) {
	balance, err := s.bg.GetBalance(c, address)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"balance": balance})
}
