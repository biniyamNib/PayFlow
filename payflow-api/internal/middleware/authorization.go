package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)


func RequireRole(allowedRoles ...string) gin.HandlerFunc {
    return func(c *gin.Context) {
        roleValue, exists := c.Get("role")
        if !exists {
            c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
                "error": "Forbidden",
            })
            return
        }

        role, ok := roleValue.(string)
        if !ok || role == "" {
            c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
                "error": "Forbidden",
            })
            return
        }

        for _, allowedRole := range allowedRoles {
            if allowedRole == role {
                c.Next()
                return
            }
        }

        c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
            "error": "Forbidden",
        })
    }
}
