package usecase

import (
    "errors"
    "restfull_gin_gorm/internal/domain"
    "restfull_gin_gorm/internal/model"
    "restfull_gin_gorm/pkg/middleware"
    "restfull_gin_gorm/pkg/utils"
)

type UserUseCase struct {
    repo domain.UserRepository
}

func NewUserUseCase(repo domain.UserRepository) *UserUseCase {
    return &UserUseCase{repo: repo}
}

func (uc *UserUseCase) Register(req model.RegisterRequest) error {
    hash, err := utils.HashPassword(req.Password)
    if err != nil {
        return err
    }
    user := &domain.User{
        Name:     req.Name,
        Email:    req.Email,
        Password: hash,
        Role:     "user",
    }
    return uc.repo.Create(user)
}

func (uc *UserUseCase) Login(req model.LoginRequest) (string, string, error) {
    user, err := uc.repo.FindByEmail(req.Email)
    if err != nil {
        return "", "", errors.New("invalid credentials")
    }
    if !utils.CheckPasswordHash(req.Password, user.Password) {
        return "", "", errors.New("invalid credentials")
    }
    access, err := middleware.GenerateToken(user.ID, user.Email, user.Role)
    if err != nil {
        return "", "", err
    }
    refresh, err := middleware.GenerateRefreshToken(user.ID, user.Email, user.Role)
    if err != nil {
        return "", "", err
    }
    return access, refresh, nil
}

func (uc *UserUseCase) RefreshToken(refresh string) (string, error) {
    claims, err := middleware.ParseToken(refresh)
    if err != nil {
        return "", err
    }
    role, ok := claims["role"].(string)
    if !ok {
        return "", errors.New("role claim missing or invalid")
    }
    access, err := middleware.GenerateToken(
        uint(claims["user_id"].(float64)),
        claims["email"].(string),
        role,
    )
    if err != nil {
        return "", err
    }
    return access, nil
}
