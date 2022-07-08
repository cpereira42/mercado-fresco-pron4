package inbound_orders

import "time"

type Service interface {
	GetAll() ([]ReportInboundOrders, error)
	GetByID(id int) (ReportInboundOrders, error)
	Create(inboundOrders InboundOrdersCreate) (InboundOrdersResponse, error)
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s service) GetAll() ([]ReportInboundOrders, error) {
	reportInboundOrders, err := s.repository.GetAll()
	if err != nil {
		return []ReportInboundOrders{}, err
	}
	return reportInboundOrders, nil
}

func (s service) GetByID(id int) (ReportInboundOrders, error) {
	reportInboundOrder, err := s.repository.GetByID(id)
	if err != nil {
		return ReportInboundOrders{}, err
	}
	return reportInboundOrder, nil
}
func (s service) Create(inboundOrders InboundOrdersCreate) (InboundOrdersResponse, error) {
	currentTime := time.Now()

	orderDate := time.Date(currentTime.Year(),
		currentTime.Month(),
		currentTime.Day(),
		currentTime.Hour(),
		currentTime.Minute(),
		currentTime.Second(),
		100,
		time.Local).Format("2006-01-02 15:04:05")

	inboundOrder, err := s.repository.Create(orderDate, inboundOrders)

	if err != nil {
		return InboundOrdersResponse{}, err
	}

	return inboundOrder, nil
}
