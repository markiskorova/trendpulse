func Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenStr := extractTokenFromHeader(c.Request)
		userID, err := parseJWT(tokenStr)
		if err == nil {
			ctx := context.WithValue(c.Request.Context(), "userID", userID)
			c.Request = c.Request.WithContext(ctx)
		}
		c.Next()
	}
}
