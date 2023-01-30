package cronjob

import (
	"time"

	"github.com/blackhorseya/irent/pkg/adapters"
	"github.com/blackhorseya/irent/pkg/contextx"
	rb "github.com/blackhorseya/irent/pkg/entity/domain/rental/biz"
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

	biz rb.IBiz

	taskC chan time.Time
	done  chan struct{}
}

func NewImpl(opts *Options, logger *zap.Logger, biz rb.IBiz) adapters.Cronjob {
	return &impl{
		opts:   opts,
		logger: logger,
		biz:    biz,
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
		case <-ticker.C:
			i.addTask()
		case <-i.taskC:
			err := i.do()
			if err != nil {
				i.logger.Error("execute task got error", zap.Error(err))
			}
		}
	}
}

func (i *impl) addTask() {
	select {
	case i.taskC <- time.Now():
	case <-time.After(50 * time.Millisecond):
		return
	}
}

func (i *impl) do() error {
	ctx := contextx.BackgroundWithLogger(i.logger)
	defer ctx.Elapsed("[do]")()

	ctx.Debug("executing task...")

	_, err := i.biz.UpdateInfoCars(ctx)
	if err != nil {
		return err
	}

	return nil
}
