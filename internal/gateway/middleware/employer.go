package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/ahdaan67/jobportal/internal/gateway/response"
	"github.com/ahdaan67/jobportal/utils/token"
	"github.com/gin-gonic/gin"
)


func EmployerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenHeader := c.GetHeader("Authorization")

		if tokenHeader == "" {
			response := response.ErrorResponse{
				Status:  "error",
				Code:    http.StatusUnauthorized,
				Message: "No auth header provided",
			}
			c.JSON(http.StatusUnauthorized, response)
			c.Abort()
			return
		}

		splitted := strings.Split(tokenHeader, " ")
		if len(splitted) != 2 || splitted[0] != "Bearer" {
			response := response.ErrorResponse{
				Status:  "error",
				Code:    http.StatusUnauthorized,
				Message: "Invalid Token Format",
			}
			c.JSON(http.StatusUnauthorized, response)
			c.Abort()
			return
		}

		tokenPart := splitted[1]
		tokenClaims, err := token.ValidateToken(tokenPart)
		if err != nil {
			response := response.ErrorResponse{
				Status:  "error",
				Code:    http.StatusUnauthorized,
				Message: "Invalid Token",
				Errors:  map[string]string{"token_error": err.Error()},
			}
			c.JSON(http.StatusUnauthorized, response)
			c.Abort()
			return
		}

		fmt.Printf("job-seeker Token Claims: %+v\n", tokenClaims)

		if tokenClaims.Role != "employer" {
			response := response.ErrorResponse{
				Status:  "error",
				Code:    http.StatusForbidden,
				Message: "Forbidden: Insufficient Role",
			}
			c.JSON(http.StatusForbidden, response)
			c.Abort()
			return
		}

		jobseekerID := tokenClaims.Id
		c.Set("id", jobseekerID)

		c.Next()
	}
}