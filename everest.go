package everest_api

import "github.com/everestafrica/everest-api/internal/config"

func Main(cfg *config.Config) error {

	serv := server{
		cfg:            cfg,
		requestTimeout: 5,
	}

	return serv.Start()

}
