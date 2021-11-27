package freegames

import (
	"context"
	"errors"
	"testing"
	"time"

	freegames "github.com/arkiant/freegames/internal"
	"github.com/arkiant/freegames/internal/platform/platform/mockplatform"
	"github.com/arkiant/freegames/internal/platform/storage/mockrepository"
	"github.com/arkiant/freegames/kit/clock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetFreeGames(t *testing.T) {

	tests := []struct {
		name       string
		repository freegames.GameRepository
		platforms  []freegames.Platform
		response   interface{}
		err        error
	}{
		{
			name: "given a mock platform with a game and a empty repository when we get free games then we get the new games and persist the games.",
			repository: func() freegames.GameRepository {
				repository := mockrepository.GameRepository{}
				repository.On("GetGames", mock.Anything).Return(freegames.FreeGames{}, nil)
				repository.On("Exists", mock.Anything).Return(false)
				repository.On("Store", mock.Anything).Return(nil)
				return &repository
			}(),
			platforms: func() []freegames.Platform {
				mockPlatform := mockplatform.Platform{}
				mockPlatform.On("GetName").Return(func() string {
					return "mockplatform"
				})
				mockPlatform.On("Run").Return(
					func() freegames.FreeGames {
						return freegames.FreeGames{
							freegames.Game{
								Name:        "test_game",
								Platform:    "mockplatform",
								CreatedAt:   clock.NewParseTime("2021-11-26T00:00:00Z00:00"),
								UpdatedAt:   clock.NewParseTime("2021-11-26T00:00:00Z00:00"),
								AvailableTo: clock.NewParseTime("2021-11-26T00:00:00Z00:00").Add(5 * 24 * time.Hour),
							},
						}
					}, nil)
				return []freegames.Platform{&mockPlatform}
			}(),
			response: freegames.FreeGames{
				freegames.Game{
					Name:        "test_game",
					Platform:    "mockplatform",
					CreatedAt:   clock.NewParseTime("2021-11-26T00:00:00Z00:00"),
					UpdatedAt:   clock.NewParseTime("2021-11-26T00:00:00Z00:00"),
					AvailableTo: clock.NewParseTime("2021-11-26T00:00:00Z00:00").Add(5 * 24 * time.Hour),
				},
			},
		},
		{
			name: "given a mock platform without games and a repository with games when we get free games then we get the games from repository.",
			repository: func() freegames.GameRepository {
				repository := mockrepository.GameRepository{}
				repository.On("GetGames", mock.Anything).Return(freegames.FreeGames{
					freegames.Game{
						Name:        "test_game",
						Platform:    "mockplatform",
						CreatedAt:   clock.NewParseTime("2021-11-26T00:00:00Z00:00"),
						UpdatedAt:   clock.NewParseTime("2021-11-26T00:00:00Z00:00"),
						AvailableTo: clock.NewParseTime("2021-11-26T00:00:00Z00:00").Add(5 * 24 * time.Hour),
					},
				}, nil)
				return &repository
			}(),
			platforms: func() []freegames.Platform {
				mockPlatform := mockplatform.Platform{}
				return []freegames.Platform{&mockPlatform}
			}(),
			response: freegames.FreeGames{
				freegames.Game{
					Name:        "test_game",
					Platform:    "mockplatform",
					CreatedAt:   clock.NewParseTime("2021-11-26T00:00:00Z00:00"),
					UpdatedAt:   clock.NewParseTime("2021-11-26T00:00:00Z00:00"),
					AvailableTo: clock.NewParseTime("2021-11-26T00:00:00Z00:00").Add(5 * 24 * time.Hour),
				},
			},
		},
		{
			name: "given a repository with error when we get free games then the service throws the error",
			repository: func() freegames.GameRepository {
				repository := mockrepository.GameRepository{}
				repository.On("GetGames", mock.Anything).Return(nil, errors.New("repository error"))
				return &repository
			}(),
			platforms: func() []freegames.Platform {
				mockPlatform := mockplatform.Platform{}
				return []freegames.Platform{&mockPlatform}
			}(),
			response: freegames.FreeGames{
				freegames.Game{
					Name:        "test_game",
					Platform:    "mockplatform",
					CreatedAt:   clock.NewParseTime("2021-11-26T00:00:00Z00:00"),
					UpdatedAt:   clock.NewParseTime("2021-11-26T00:00:00Z00:00"),
					AvailableTo: clock.NewParseTime("2021-11-26T00:00:00Z00:00").Add(5 * 24 * time.Hour),
				},
			},
			err: errors.New("repository error"),
		},
		{
			name: "given an empty repository and no platforms when we get free games then we don't get games.",
			repository: func() freegames.GameRepository {
				repository := mockrepository.GameRepository{}
				return &repository
			}(),
			platforms: []freegames.Platform{},
			response:  freegames.FreeGames(nil),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			service := NewService(tc.repository, tc.platforms...)
			response, err := service.GetFreeGames(context.Background())
			if tc.err != nil {
				assert.Equal(t, tc.err, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.response, response)
			}
		})
	}
}
