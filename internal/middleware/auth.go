package middleware

import (
	"github.com/gin-gonic/gin"
)

// UserData hard-coded data untuk simulasi SSO
type UserData struct {
	NIK        string
	CostCenter string
	Levels     []int
}

// AuthMiddleware mengganti SSO auth dari Python
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Hard-coded user data untuk simulasi
		userData := UserData{
			NIK:        "1025020026",
			CostCenter: "GTIK01",
			Levels:     []int{1000}, // 0 = SUPER_ADMIN, 100 = MANAGER, 1000 = USER
		}

		// Set data ke context untuk digunakan handler
		c.Set("nik", userData.NIK)
		c.Set("costCenter", userData.CostCenter)
		c.Set("levels", userData.Levels)

		// Lanjutkan ke handler berikutnya
		c.Next()
	}
}

// RoleMiddleware middleware untuk role-based access
func RoleMiddleware(validRoles ...int) gin.HandlerFunc {
	return func(c *gin.Context) {
		levels, exists := c.Get("levels")
		if !exists {
			c.AbortWithStatusJSON(401, gin.H{
				"response_code": "99",
				"message":      "Unauthorized: No role data",
				"data":         []interface{}{},
			})
			return
		}

		userLevels, ok := levels.([]int)
		if !ok {
			c.AbortWithStatusJSON(401, gin.H{
				"response_code": "99",
				"message":      "Unauthorized: Invalid role data format",
				"data":         []interface{}{},
			})
			return
		}

		// Check if user has any of the valid roles
		hasRole := false
		for _, userLevel := range userLevels {
			for _, validRole := range validRoles {
				if userLevel == validRole {
					hasRole = true
					break
				}
			}
		}

		if !hasRole {
			c.AbortWithStatusJSON(403, gin.H{
				"response_code": "99",
				"message":      "Forbidden: Insufficient privileges",
				"data":         []interface{}{},
			})
			return
		}

		c.Next()
	}
}