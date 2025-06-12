package handler

import (
	"context"

	pb "github.com/VicFunas/cms-wikium/proto"

	"github.com/VicFunas/cms-wikium/internal/service"
)

type ModHandler struct {
	pb.UnimplementedModServiceServer
	modService *service.ModService
}

func NewModHandler(s *service.ModService) *ModHandler {
	return &ModHandler{
		modService: s,
	}
}

func (h *ModHandler) ListMods(ctx context.Context, req *pb.ListModsRequest) (*pb.ListModsResponse, error) {
	mods, err := h.modService.ListMods(ctx)
	if err != nil {
		// This will trigger for both invalid ID format and not found
		return nil, err
	}

	var modsResponse []*pb.Mod
	for _, mod := range mods {
		modsResponse = append(modsResponse, &pb.Mod{
			Id:   mod.ID.Hex(),
			Name: mod.Name,
		})
	}

	return &pb.ListModsResponse{Mods: modsResponse}, nil
}

func (h *ModHandler) GetMod(ctx context.Context, req *pb.GetModRequest) (*pb.GetModResponse, error) {
	mod, err := h.modService.GetMod(ctx, req.ModId)
	if err != nil {
		// This will trigger for both invalid ID format and not found
		return nil, err
	}

	return &pb.GetModResponse{Mod: &pb.Mod{
		Id:          mod.ID.Hex(),
		Name:        mod.Name,
		Description: mod.Description,
	}}, nil
}

func (h *ModHandler) CreateMod(ctx context.Context, req *pb.CreateModRequest) (*pb.CreateModResponse, error) {
	mod, err := h.modService.CreateMod(ctx, req)
	if err != nil {
		return nil, err
	}

	return &pb.CreateModResponse{Mod: &pb.Mod{
		Id:          mod.ID.Hex(),
		Name:        mod.Name,
		Description: mod.Description,
	}}, nil
}

// Implementation of the gRPC endpoint
func (h *ModHandler) UpdateMod(ctx context.Context, req *pb.UpdateModRequest) (*pb.UpdateModResponse, error) {
	mod, err := h.modService.UpdateMod(ctx, req)
	if err != nil {
		return nil, err
	}

	return &pb.UpdateModResponse{Mod: &pb.Mod{
		Id:          mod.ID.Hex(),
		Name:        mod.Name,
		Description: mod.Description,
	}}, nil
}
