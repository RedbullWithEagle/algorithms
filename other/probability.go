package other

import (
	"fmt"
	"math/rand"
	"time"
)

//sample  简单的随机算法
//每次随机一个数，看看落在哪个区间
//从结果来看，只是近似随机，当样本足够大，接近给出的区间
func sample(cdf []float32) int {
	r := rand.Float32()

	bucket := 0
	for r > cdf[bucket] {
		bucket++
	}
	return bucket
}

func TestSamplePro() {
	//使用随机函数，记得初始化时间作为种子
	rand.Seed(time.Now().Unix())
	// probability density function
	/*pdf := []float32{0.3, 0.4, 0.2, 0.1}

	// get cdf
	cdf := []float32{0.0, 0.0, 0.0, 0.0}
	cdf[0] = pdf[0]
	for i := 1; i < 4; i++ {
		cdf[i] = cdf[i-1] + pdf[i]
	}

	// test sampling with 100 samples
	samples := []float32{0.0, 0.0, 0.0, 0.0}

	count := 1000
	for i := 0; i < count; i++ {
		samples[sample(cdf)]++
	}

	// normalize
	for i := 0; i < 4; i++ {
		samples[i] /= 1000.
	}

	fmt.Println(samples)
	fmt.Println(pdf)*/

	result := make(map[uint32]uint32)
	for i := 0; i < 10000; i++ {
		GetNumber()
	}

	for i := 0; i < 10000; i++ {
		num := GetNumber()
		result[num] += 1
	}
	fmt.Println(result)
}

/******************************************************
*精准随机，每种奖励出现的次数是固定的
*比如果 10000抽奖，1出现的次数必定是600次
*2出现的概率必定是700次，.....
*该版本的优化版本，可以存储在redis的list中，取元素，直接lpop。或者rpop
*******************************************************/
var (
	// 随机概率表  总数是10000
	DefaultProbability = map[uint32]uint32{
		1:  600,
		2:  700,
		3:  1000,
		4:  1500,
		5:  1500,
		12: 900,
		16: 900,
		17: 900,
		18: 700,
		19: 700,
		20: 600,
	}

	tmp []uint32
)

// GenerateSlice 生成概率数组
func GenerateSlice() {
	tmp = make([]uint32, 0)
	for idx, val := range DefaultProbability {
		for i := uint32(0); i < val; i++ {
			tmp = append(tmp, idx)
		}
	}

	//随机洗牌算法
	lenSlice := len(tmp)
	for i := lenSlice - 1; i >= 0; i-- {
		idx := rand.Intn(i + 1)
		if i != idx {
			data := tmp[i]
			tmp[i] = tmp[idx]
			tmp[idx] = data
		}
	}
	return
}

// GetNumber 生成随机数
func GetNumber() uint32 {
	//这里需要考虑并发的问题
	if len(tmp) == 0 {
		GenerateSlice()
	}

	count := len(tmp)
	num := tmp[count-1]
	tmp = tmp[:count-1]
	fmt.Println(len(tmp))
	return num
}
