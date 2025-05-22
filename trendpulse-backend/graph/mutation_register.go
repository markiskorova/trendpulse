func (r *mutationResolver) Register(ctx context.Context, email string, password string) (*model.AuthPayload, error) {
    hashed, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    user := models.User{Email: email, Password: string(hashed)}
    if err := r.DB.Create(&user).Error; err != nil {
        return nil, err
    }

    token := generateJWT(user.ID)
    return &model.AuthPayload{Token: token, User: user}, nil
}
