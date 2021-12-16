package params

type Login struct {
	Email    string `form:"email" binding:"required"`
	Password string `form:"password" binding:"required"`
}

type Register struct {
	Email    string `form:"email" binding:"required"`
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}

type ChangePassword struct {
	Password string `form:"password" binding:"required"`
}

type Verify struct {
	Token string `form:"token"`
	Sign  string `form:"sign" binding:"required"`
	Ts    int64  `form:"ts" binding:"required"`
}

type UserInfo struct {
	Email    string `form:"email" binding:"required"`
	Username string `form:"username" binding:"required"`
	Mobile   string `form:"mobile" binding:"required"`
}
