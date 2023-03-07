package server

import (
	"context"
	"fmt"

	"github.com/seal-io/seal/pkg/dao/model"
)

type initOptions struct {
	ModelClient *model.Client
}

func (r *Server) init(ctx context.Context, opts initOptions) error {
	// create model schema.
	var err = opts.ModelClient.Schema.Create(ctx)
	if err != nil {
		return fmt.Errorf("error creating model schemas: %w", err)
	}

	// initialize critical resources.
	type initor struct {
		name string
		init func(context.Context, initOptions) error
	}

	var inits = []initor{
		{name: "settings", init: r.initSettings},
		{name: "subscribers", init: r.initSubscribers},
	}
	inits = append(inits,
		initor{name: "modules", init: r.initModules},
		initor{name: "roles", init: r.initRoles},
		initor{name: "subjects", init: r.initSubjects},
		initor{name: "perspective", init: r.initPerspectives},
		initor{name: "cronjobs", init: r.initCronJobs},
	)
	if r.EnableAuthn {
		inits = append(inits,
			initor{name: "casdoor", init: r.initCasdoor},
		)
	}

	for i := range inits {
		if err = inits[i].init(ctx, opts); err != nil {
			return fmt.Errorf("%s: %w", inits[i].name, err)
		}
	}
	return nil
}
