package service

import (
	"context"

	"github.com/VicFunas/cms-wikium/internal/domain"
	"github.com/VicFunas/cms-wikium/internal/repository"
	pb "github.com/VicFunas/cms-wikium/proto"
)

type ModService struct {
	modRepo *repository.ModRepository
}

func NewModService(repo *repository.ModRepository) *ModService {
	return &ModService{
		modRepo: repo,
	}
}

// GetMod just passes the call through to the repository.
func (s *ModService) GetMod(ctx context.Context, id string) (domain.Mod, error) {
	return s.modRepo.GetModByID(ctx, id)
}

// CreateMod creates the domain model and calls the repository to save it.
func (s *ModService) CreateMod(ctx context.Context, req *pb.CreateModRequest) (domain.Mod, error) {
	mod := domain.Mod{
		Name: req.Name,
	}
	return s.modRepo.CreateMod(ctx, mod)
}
