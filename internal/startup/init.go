package startup

import (
	"context"
	"fmt"
	"gitee.com/unitedrhino/share/domain/application"
	"gitee.com/unitedrhino/share/events/topics"
	"gitee.com/unitedrhino/share/utils"
	"gitee.com/unitedrhino/things/service/dmsvr/pb/dm"
	"github.com/zeromicro/go-zero/core/logx"
	"lightsvr/internal/svc"
	"time"
)

func InitEventBus(svcCtx *svc.ServiceContext) {
	err := svcCtx.FastEvent.QueueSubscribe(topics.ApplicationDeviceReportThingPropertyAllDevice, func(ctx context.Context, t time.Time, body []byte) error {
		if t.Before(time.Now().Add(-time.Second * 2)) { //2秒之前的跳过
			return nil
		}
		var stu application.PropertyReport
		err := utils.Unmarshal(body, &stu)
		if err != nil {
			logx.WithContext(ctx).Errorf("Subscribe.QueueSubscribe.Unmarshal body:%v err:%v", string(body), err)
			return err
		}
		logx.WithContext(ctx).Info(string(body))
		ret, err := svcCtx.DeviceInteract.PropertyControlSend(ctx, &dm.PropertyControlSendReq{
			ProductID:  stu.Device.ProductID,
			DeviceName: stu.Device.DeviceName,
			Data:       fmt.Sprintf(`{"load":%v}`, stu.Param),
		})
		logx.WithContext(ctx).Info(ret, err)
		return nil
	})
	logx.Must(err)

	err = svcCtx.FastEvent.Start()
	logx.Must(err)
}
