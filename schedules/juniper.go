package schedules

import (
	"github.com/astaxie/beego"
	"time"
)

func (s *Schedule)JuniperConfigTask() error{

	beego.Info(time.Now(), s.T.Name)
	return nil
}