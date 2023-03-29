/*
Create: 2023/3/30
Project: FuShin
Github: https://github.com/landers1037
Copyright Renj
*/

package autocert

import "github.com/JJApplication/fushin/cron"

// 自动刷新任务

func AutoCert(c *Cert, err error) {
	gr := cron.NewGroup(cron.Monthly)
	_, err = gr.AddFunc(func() {
		res, e := c.start()
		if e != nil {
			err = e
			return
		}
		err = SaveCert(c.certPath, res)
	})
	gr.Start()
}
