package http

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
	"social-network/internal/domain"
)

type Endpoints struct {
	Auth   *AuthEndpoints
	Social *SocialEndpoints
}

func MakeEndpoints(auth domain.AuthService, social domain.SocialService) *Endpoints {
	return &Endpoints{
		Auth: &AuthEndpoints{
			SignUp:       makeSignUpEndpoint(auth),
			SignIn:       makeSignInEndpoint(auth),
			RefreshToken: makeRefreshTokenEndpoint(auth),
		},
		Social: &SocialEndpoints{
			Questionnaires: makeQuestionnairesEndpoint(social)},
	}
}

type AuthEndpoints struct {
	SignUp       gin.HandlerFunc
	SignIn       gin.HandlerFunc
	RefreshToken gin.HandlerFunc
}

func makeSignUpEndpoint(svc domain.AuthService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var request SignUpRequest
		if err := c.Bind(&request); err != nil {
			c.JSON(http.StatusBadRequest, ErrorResponse{
				Message: err.Error(),
			})

			return
		}

		credentials, err := domain.NewCredentials(request.Email, request.Password)
		if err != nil {
			c.JSON(http.StatusInternalServerError, ErrorResponse{
				Message: err.Error(),
			})

			return
		}

		if err = svc.SignUp(c, &domain.User{
			Credentials: *credentials,
			Name:        request.Name,
			Surname:     request.Surname,
			Birthday:    request.Birthday,
			Sex:         request.Sex,
			City:        request.City,
			Interests:   request.Interests,
		}); err != nil {
			c.JSON(http.StatusInternalServerError, ErrorResponse{
				Message: err.Error(),
			})

			return
		}

		c.JSON(http.StatusOK, EmptyResponse{})
	}
}

func makeSignInEndpoint(svc domain.AuthService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var request SignInRequest
		if err := c.Bind(&request); err != nil {
			c.JSON(http.StatusBadRequest, ErrorResponse{
				Message: err.Error(),
			})

			return
		}

		tokenPair, err := svc.SignIn(c, &domain.Credentials{
			Login:    request.Login,
			Password: request.Password,
		})
		if err != nil {
			c.JSON(http.StatusInternalServerError, ErrorResponse{
				Message: err.Error(),
			})

			return
		}

		c.JSON(http.StatusOK, SignInResponse{
			AccessToken:  tokenPair.AccessToken,
			RefreshToken: tokenPair.RefreshToken,
		})
	}
}

func makeRefreshTokenEndpoint(svc domain.AuthService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var request RefreshTokenRequest
		if err := c.Bind(&request); err != nil {
			c.JSON(http.StatusBadRequest, ErrorResponse{
				Message: err.Error(),
			})

			return
		}

		tokenPair, err := svc.RefreshToken(c, request.RefreshToken)

		switch err {
		case nil:
			c.JSON(http.StatusOK, SignInResponse{
				AccessToken:  tokenPair.AccessToken,
				RefreshToken: tokenPair.RefreshToken,
			})
		case sql.ErrNoRows:
			c.JSON(http.StatusUnauthorized, ErrorResponse{
				Message: "token is expired",
			})
		default:
			c.JSON(http.StatusInternalServerError, ErrorResponse{
				Message: err.Error(),
			})
		}
	}
}

type SocialEndpoints struct {
	Questionnaires gin.HandlerFunc
}

func makeQuestionnairesEndpoint(svc domain.SocialService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var header AuthorizationHeader
		if err := c.ShouldBindHeader(&header); err != nil {
			c.JSON(http.StatusUnauthorized, ErrorResponse{
				Message: err.Error(),
			})

			return
		}

		var request QuestionnairesRequest
		if err := c.Bind(&request); err != nil {
			c.JSON(http.StatusBadRequest, ErrorResponse{
				Message: err.Error(),
			})

			return
		}

		// todo: make authorize action

		quest, count, err := svc.GetQuestionnaires(c, "080b2d4b-096e-11eb-b4fd-0242c0a80002", *request.Limit, request.Offset)
		if err != nil {
			c.JSON(http.StatusInternalServerError, ErrorResponse{
				Message: err.Error(),
			})

			return
		}

		c.JSON(http.StatusOK, QuestionnairesResponse{
			Questionnaires: quest,
			Count:          count,
		})
	}
}