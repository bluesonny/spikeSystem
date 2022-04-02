package localSpike

import "log"

type LocalSpike struct {
	LocalInStock     int64
	LocalSalesVolume int64
}

//本地扣库存,返回bool值
func (spike *LocalSpike) LocalDeductionStock() bool {
	spike.LocalSalesVolume = spike.LocalSalesVolume + 1
	log.Printf("准备卖出第%v个", spike.LocalSalesVolume)
	log.Printf("本地是%v", spike.LocalInStock)
	ret := spike.LocalSalesVolume <= spike.LocalInStock
	log.Printf("本地卖出结果返回%v", ret)
	return ret
}
