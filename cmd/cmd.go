package cmd

// Restful http server
type Restful interface {
	InitRouting() error
}

// Grpc is a grpc server
type Grpc interface {
	// Start start grpc server
	Start() error

	// Stop stop grpc server
	Stop() error
}

// Cronjob is a cronjob adapters
type Cronjob interface {
	// Start to run
	Start() error

	// Stop to end
	Stop() error
}
