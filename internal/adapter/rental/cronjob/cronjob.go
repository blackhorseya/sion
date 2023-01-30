package cronjob

import (
	"time"

	"github.com/blackhorseya/irent/pkg/adapters"
	"github.com/google/wire"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var RentalSet = wire.NewSet(NewOptions, NewImpl)

type Options struct {
	Enabled  bool          `json:"enabled" yaml:"enabled"`
	Interval time.Duration `json:"interval" yaml:"interval"`
}

func NewOptions(v *viper.Viper, logger *zap.Logger) (*Options, error) {
	ret := new(Options)

	err := v.UnmarshalKey("cronjob", &ret)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to load cronjob options")
	}

	logger.Info("Loaded cronjob options success", zap.Any("options", ret))

	return ret, nil
}

type impl struct {
	opts   *Options
	logger *zap.Logger

	taskC chan time.Time
	done  chan struct{}
}

func NewImpl(opts *Options, logger *zap.Logger) adapters.Cronjob {
	return &impl{
		opts:   opts,
		logger: logger,
		taskC:  make(chan time.Time, 1),
		done:   make(chan struct{}),
	}
}

func (i *impl) Start() error {
	if !i.opts.Enabled {
		i.logger.Info("Not enabled cronjob service")
		return nil
	}

	i.logger.Info("starting cronjob service...")

	go i.worker()

	i.logger.Info("started cronjob service")

	return nil
}

func (i *impl) Stop() error {
	if !i.opts.Enabled {
		return nil
	}

	i.logger.Info("stopping cronjob service...")

	i.done <- struct{}{}

	i.logger.Info("sent stop signal")

	return nil
}

func (i *impl) worker() {
	ticker := time.NewTicker(i.opts.Interval)

	for {
		select {
		case <-i.done:
			return
		case t := <-ticker.C:
			i.addTask(t)
		case <-i.taskC:
			// todo: 2023/1/30|sean|impl me
			i.logger.Debug("executing task")
		}
	}
}

func (i *impl) addTask(t time.Time) {
	select {
	case i.taskC <- t:
	case <-time.After(50 * time.Millisecond):
		return
	}
}
