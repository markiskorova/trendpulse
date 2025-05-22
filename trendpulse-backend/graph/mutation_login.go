func (r *mutationResolver) Login(ctx context.Context, email string, password string) (*model.AuthPayload, error) {
	var user models.User
	if err := r.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, errors.New("invalid credentials")
	}

	token := generateJWT(user.ID)
	return &model.AuthPayload{Token: token, User: user}, nil
}
