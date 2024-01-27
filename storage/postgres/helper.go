package postgres

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/Masterminds/squirrel"
	pb "github.com/golanguzb70/go_subscription_service/genproto/subscription_service"
	"github.com/golanguzb70/go_subscription_service/pkg/logger"
	"github.com/lib/pq"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func HandleDatabaseError(err error, log logger.Logger, message string) error {
	if err == nil {
		return nil
	}
	log.Error(message + ": " + err.Error())
	switch err {
	case sql.ErrNoRows:
		return status.Error(codes.NotFound, "This information is not exists.")
	case sql.ErrConnDone:
		return err
	case sql.ErrTxDone:
		return err
	}

	switch e := err.(type) {
	case *pq.Error:
		// Handle Postgres-specific errors
		switch e.Code.Name() {
		case "unique_violation":
			return status.Error(codes.AlreadyExists, "Already exists")
		case "foreign_key_violation":
			return status.Error(codes.InvalidArgument, "Oops something went wrong")
		default:
			return err
		}
	default:
		// Handle all other errors
		return err
	}
}

func PrepareWhere(filters []*pb.Filters) squirrel.And {
	res := squirrel.And{}

	for _, e := range filters {
		switch e.Type {
		case "search":
			res = append(res, squirrel.ILike{e.Field: "%" + e.Value + "%"})
		case "=":
			res = append(res, squirrel.Eq{e.Field: e.Value})
		case "<=":
			res = append(res, squirrel.LtOrEq{e.Field: e.Value})
		case "<":
			res = append(res, squirrel.Lt{e.Field: e.Value})
		case ">=":
			res = append(res, squirrel.GtOrEq{e.Field: e.Value})
		case ">":
			res = append(res, squirrel.Gt{e.Field: e.Value})
		}
	}

	return res
}

func PrepareOrder(orders []*pb.SortBy) string {
	res := []string{}
	for _, e := range orders {
		switch e.Type {
		case "desc", "asc":
			res = append(res, fmt.Sprintf("%s %s", e.Field, e.Type))
		}
	}

	return strings.Join(res, ", ")
}

func ParseTimeString(input string) (time.Time, error) {
	layout := "2006-01-02 15:04"
	return time.Parse(layout, input)
}
