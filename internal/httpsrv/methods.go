package httpsrv

import (
	"context"
	"go.uber.org/zap"
	"net/http"
	"sync"

	"github.com/romanchechyotkin/effective-mobile-test-task/internal/httpsrv/middleware"

	"github.com/gin-gonic/gin"
)

func (s *Server) RegisterRoutes() {
	s.router.Use(middleware.CORSMiddleware())

	s.router.GET("/status", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "ok\n")
	})

	usersGroup := s.router.Group("/users")
	usersGroup.GET("/")
	usersGroup.POST("/", s.createUser)
	usersGroup.GET("/:id")
	usersGroup.PATCH("/:id")
	usersGroup.DELETE("/:id")
}

// @Summary Create user
// @Description Endpoint for creating and saving user to database
// @Produce application/json
// @Success 201 {object} UserResponseDto
// @Router /users [post]
func (s *Server) createUser(ctx *gin.Context) {
	var userDto UserRequest
	var wg sync.WaitGroup

	err := ctx.ShouldBindJSON(&userDto)
	if err != nil {
		s.log.Error("failed to read json", zap.Error(err))
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	s.log.Info("decoded user dto", zap.Any("user", userDto))

	var userAge int
	wg.Add(1)
	go func() {
		defer wg.Done()

		age, err := s.client.GetAge(context.Background(), userDto.FirstName)
		if err != nil {
			s.log.Error("failed to get age", zap.String("name", userDto.FirstName), zap.Error(err))
			return
		}

		s.log.Info("got age", zap.String("name", userDto.FirstName), zap.Int("age", age.Age))
		userAge = age.Age
	}()

	var userGender string
	wg.Add(1)
	go func() {
		defer wg.Done()

		gender, err := s.client.GetGender(context.Background(), userDto.FirstName)
		if err != nil {
			s.log.Error("failed to get age", zap.String("name", userDto.FirstName), zap.Error(err))
			return
		}

		s.log.Info("got age", zap.String("name", userDto.FirstName), zap.String("gender", gender.Gender))
		userGender = gender.Gender
	}()

	var userNationality string
	wg.Add(1)
	go func() {
		defer wg.Done()

		nationality, err := s.client.GetNationality(context.Background(), userDto.FirstName)
		if err != nil {
			s.log.Error("failed to get age", zap.String("name", userDto.FirstName), zap.Error(err))
			return
		}

		nation := nationality.Country[0].CountryID
		s.log.Info("got age", zap.String("name", userDto.FirstName), zap.String("nationality", nation))
		userNationality = nation
	}()

	wg.Wait()

	response := &UserResponse{
		LastName:    userDto.LastName,
		FirstName:   userDto.FirstName,
		SecondName:  userDto.SecondName,
		Age:         userAge,
		Gender:      userGender,
		Nationality: userNationality,
	}

	//id, err := h.repository.saveUser(ctx, response)
	//if err != nil {
	//	logger.Error(h.log, "error during saving user to database", err)
	//	ctx.JSON(http.StatusInternalServerError, gin.H{
	//		"error": err.Error(),
	//	})
	//	return
	//}
	//
	//response.ID = id
	//
	//h.log.Info("user created", slog.Any("user", response))
	ctx.JSON(http.StatusCreated, response)
}

// @Summary All users
// @Description Endpoint for getting all users
// @Produce application/json
// @Success 200 {object} []UserResponseDto{}
// @Router /users [get]
//func (s *Server) getAllUsers(ctx *gin.Context) {
//	sort := ctx.Query("sort")
//	limit := ctx.Query("limit")
//	h.log.Debug("got sort query value", slog.String("sort", sort))
//	h.log.Debug("got limit query value", slog.String("limit", limit))
//
//	var users []*UserResponseDto
//	var err error
//
//	switch sort {
//	case SORT_BY_ASC_AGE:
//		users, err = h.repository.getAllUsers(ctx, sort, limit)
//	case SORT_BY_DESC_AGE:
//		users, err = h.repository.getAllUsers(ctx, sort, limit)
//	default:
//		users, err = h.repository.getAllUsers(ctx, sort, limit)
//	}
//
//	if err != nil {
//		if errors.Is(err, ErrNotFound) {
//			ctx.JSON(http.StatusNotFound, gin.H{
//				"error": err.Error(),
//			})
//		} else {
//			ctx.JSON(http.StatusInternalServerError, gin.H{
//				"error": err.Error(),
//			})
//		}
//		return
//	}
//
//	ctx.JSON(http.StatusOK, users)
//}

// @Summary Users Endpoint Health Check
// @Description Checking health of users endpoint
// @Produce application/json
// @Success 200 {string} nasa
// @Router /users/health [get]
//func (s *Server) index(ctx *gin.Context) {
//	ctx.String(http.StatusOK, "users")
//}

// @Summary Get exact user
// @Description Endpoint for getting user with exact id
// @Produce application/json
// @Success 200 {object} UserResponseDto
// @Param id path string true "id"
// @Router /users/{id} [get]
//func (s *Server) getUser(ctx *gin.Context) {
//	id := ctx.Param("id")
//	h.log.Debug("got id param", slog.String("id", id))
//
//	user, err := h.repository.getUser(ctx, id)
//	if err != nil {
//		logger.Error(h.log, "error during db query", err)
//		if errors.Is(err, ErrNotFound) {
//			ctx.JSON(http.StatusNotFound, gin.H{
//				"error": err.Error(),
//			})
//		} else {
//			ctx.JSON(http.StatusInternalServerError, gin.H{
//				"error": err.Error(),
//			})
//		}
//		return
//	}
//
//	ctx.JSON(http.StatusOK, user)
//}

// @Summary Update exact user
// @Description Endpoint for updating user with exact id
// @Produce application/json
// @Success 204 {object} UserResponseDto
// @Param id path string true "id"
// @Router /users/{id} [patch]
//func (s *Server) updateUser(ctx *gin.Context) {
//	id := ctx.Param("id")
//	h.log.Debug("got id param", slog.String("id", id))
//
//	var dto map[string]any
//	err := ctx.ShouldBindJSON(&dto)
//	if err != nil {
//		logger.Error(h.log, "error during decoding", err)
//		ctx.JSON(http.StatusBadRequest, gin.H{
//			"error": err,
//		})
//		return
//	}
//
//	h.log.Debug("decoded update user dto", slog.Any("dto", dto))
//
//	for k, v := range dto {
//		err := h.repository.updateUser(ctx, id, k, v)
//		if err != nil {
//			logger.Error(h.log, "error during updating in database", err)
//			if errors.Is(err, ErrNotFound) {
//				ctx.JSON(http.StatusNotFound, gin.H{
//					"error": err.Error(),
//				})
//			} else {
//				ctx.JSON(http.StatusInternalServerError, gin.H{
//					"error": err.Error(),
//				})
//			}
//			return
//		}
//	}
//
//	ctx.JSON(http.StatusNoContent, gin.H{
//		"message": "updated successfully",
//	})
//}

// @Summary Delete exact user
// @Description Endpoint for deleting user with exact id
// @Produce application/json
// @Success 204 {object} UserResponseDto
// @Param id path string true "id"
// @Router /users/{id} [delete]
//func (s *Server) deleteUser(ctx *gin.Context) {
//	id := ctx.Param("id")
//	h.log.Debug("got id param", slog.String("id", id))
//
//	err := h.repository.deleteUser(ctx, id)
//	if err != nil {
//		logger.Error(h.log, "error during db query", err)
//		if errors.Is(err, ErrNotFound) {
//			ctx.JSON(http.StatusNotFound, gin.H{
//				"error": err.Error(),
//			})
//		} else {
//			ctx.JSON(http.StatusInternalServerError, gin.H{
//				"error": err.Error(),
//			})
//		}
//		return
//	}
//
//	ctx.JSON(http.StatusNoContent, gin.H{
//		"message": "deleted successfully",
//	})
//}