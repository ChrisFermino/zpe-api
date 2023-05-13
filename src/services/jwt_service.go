package services

import (
    "fmt"
    "github.com/dgrijalva/jwt-go"
    "time"
    "zpeTest/src/utils"
)

type JwtService struct {
    secretKey string
    issure    string
}

type Claim struct {
    Sum      uint64 `json:"sum"`
    FullName string `json:"fullName"`
    Email    string `json:"email"`
    Role     string `json:"role"`
    jwt.StandardClaims
}

func NewJWTService() *JwtService {
    return &JwtService{
        secretKey: utils.EnvVariable("SECRET_KEY"),
        issure:    "zpe-api",
    }
}

func (s *JwtService) GenerateToken(id uint64, name string, email string, Role string) (string, error) {
    claim := &Claim{
        id,
        name,
        email,
        Role,
        jwt.StandardClaims{
            ExpiresAt: time.Now().Add(time.Minute * 15).Unix(),
            Issuer:    s.issure,
            IssuedAt:  time.Now().Unix(),
        },
    }
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
    t, err := token.SignedString([]byte(s.secretKey))
    if err != nil {
        return "", err
    }
    return t, nil
}

func (s *JwtService) ValidateToken(token string) (*Claim, bool) {
    ParseToken, err := jwt.ParseWithClaims(token, &Claim{}, func(token *jwt.Token) (interface{}, error) {
        if _, isValid := token.Method.(*jwt.SigningMethodHMAC); !isValid {
            return nil, fmt.Errorf("invalid token: %v", token)
        }
        return []byte(s.secretKey), nil
    })
    return ParseToken.Claims.(*Claim), err == nil
}
