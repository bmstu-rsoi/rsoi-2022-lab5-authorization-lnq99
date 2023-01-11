package service

import (
	"context"

	"flight/repository"

	"github.com/lnq99/rsoi-2022-lab3-fault-tolerance-lnq99/src/pkg/model"
)

type Service interface {
	GetFlight(ctx context.Context, flightNumber string) *model.FlightResponse
	ListFlights(ctx context.Context, page, size int32) *model.PaginationResponse
}

type serviceImpl struct {
	repo repository.Repo
}

var airportMap = make(map[int32]model.Airport)

func NewService(repo repository.Repo) Service {
	return &serviceImpl{repo: repo}
}

func (s *serviceImpl) ListFlights(ctx context.Context, page, size int32) *model.PaginationResponse {
	res := model.PaginationResponse{
		TotalElements: 0,
		Page:          page,
		PageSize:      size,
		Items:         []model.FlightResponse{},
	}

	flights, err := s.repo.ListFlightsWithOffsetLimit(ctx,
		repository.ListFlightsWithOffsetLimitParams{size, (page - 1) * size})

	if err != nil {
		return &res
	}

	res.TotalElements = int32(len(flights))

	for _, f := range flights {
		fromAirport := s.loadAirport(ctx, f.FromAirportID)
		toAirport := s.loadAirport(ctx, f.ToAirportID)
		if fromAirport == nil || toAirport == nil {
			continue
		}
		res.Items = append(res.Items, model.FlightResponse{
			FlightNumber: f.FlightNumber,
			FromAirport:  fromAirport.City + " " + fromAirport.Name,
			ToAirport:    toAirport.City + " " + toAirport.Name,
			Date:         f.Datetime,
			Price:        f.Price,
		})
	}

	return &res
}

func (s *serviceImpl) GetFlight(ctx context.Context, flightNumber string) *model.FlightResponse {
	f, err := s.repo.GetFlight(ctx, flightNumber)
	if err != nil {
		return nil
	}

	fromAirport := s.loadAirport(ctx, f.FromAirportID)
	toAirport := s.loadAirport(ctx, f.ToAirportID)

	return &model.FlightResponse{
		FlightNumber: f.FlightNumber,
		FromAirport:  fromAirport.City + " " + fromAirport.Name,
		ToAirport:    toAirport.City + " " + toAirport.Name,
		Date:         f.Datetime,
		Price:        f.Price,
	}
}

func (s *serviceImpl) loadAirport(ctx context.Context, id int32) *model.Airport {
	airport, ok := airportMap[id]
	if !ok {
		newAirport, err := s.repo.GetAirport(ctx, id)
		if err != nil {
			return nil
		}
		airport = model.Airport(newAirport)
		airportMap[id] = airport
	}
	return &airport
}
