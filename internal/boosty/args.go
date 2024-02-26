package boosty

import (
	"net/url"
	"strconv"
)

const (
	defaultLimit = 5
)

type Args struct {
	Limit  int
	Offset string
}

func (args Args) QueryParams() url.Values {
	values := make(url.Values)

	if args.Limit == 0 {
		args.Limit = defaultLimit
	}
	values.Add("limit", strconv.Itoa(args.Limit))

	if args.Offset != "" {
		values.Add("offset", args.Offset)
	}

	return values
}
