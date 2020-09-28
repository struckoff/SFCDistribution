package curvetest

import (
	"github.com/struckoff/sfcdistribution/report"
	balancer "github.com/struckoff/sfcframework"
)

func CurveReport(conf CurveTestConfig, rpt report.Report) report.ReportFunc {
	if conf.Stateless {
		return curveReportStateless(conf, rpt)
	}
	return curveReportStatefull(conf, rpt)

}

func curveReportStateless(conf CurveTestConfig, rpt report.Report) report.ReportFunc {
	bal, err := balancer.NewBalancer(conf.Curve, conf.Dims, 1<<conf.Bits, conf.Tf, conf.Of, conf.Nodes)
	if err != nil {
		panic(err)
	}
	if err := bal.Optimize(); err != nil {
		panic(err)
	}

	return func(rCh chan<- report.Report) {
		for _, di := range conf.Dis {
			if n, _, err := bal.LocateData(di); err != nil {
				panic(err)
			} else {
				if err := rpt.Add(n, di); err != nil {
					panic(err)
				}
			}
		}
		rCh <- rpt
	}
}
func curveReportStatefull(conf CurveTestConfig, rpt report.Report) report.ReportFunc {
	bal, err := balancer.NewBalancer(conf.Curve, conf.Dims, 1<<conf.Bits, conf.Tf, conf.Of, conf.Nodes)
	if err != nil {
		panic(err)
	}
	if err := bal.Optimize(); err != nil {
		panic(err)
	}

	return func(rCh chan<- report.Report) {
		for _, di := range conf.Dis {
			if _, _, err := bal.LocateData(di); err != nil {
				panic(err)
			}
		}
		if err := bal.Optimize(); err != nil {
			panic(err)
		}
		for _, di := range conf.Dis {
			if n, _, err := bal.LocateData(di); err != nil {
				panic(err)
			} else {
				if err := rpt.Add(n, di); err != nil {
					panic(err)
				}
			}
		}
		rCh <- rpt
	}
}
