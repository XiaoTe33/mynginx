package logic

import (
	"crypto/sha256"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

type Balancer interface {
	Index() int
}

func NewBalancer(router Router) Balancer {
	switch router.BalanceMethod {
	case "WeightOrder":
		return NewWeightOrderBalancer(router)
	case "WeightRandom":
		return NewWeightRandomBalancer(router)
	case "ConsistentHash":
		return NewConsistentHashBalancer(router)
	default:
		panic("no such balance method: " + router.BalanceMethod)

	}
}

type WeightOrderBalancer struct {
	t       int
	weights []int
}

func (b *WeightOrderBalancer) Index() int {
	sum := 0
	for idx, i := range b.weights {
		sum += i
		if b.t < sum {
			b.t++
			return idx
		}
	}
	if b.t == sum {
		b.t = 1
		return 0
	}
	return -1
}

// NewWeightOrderBalancer 权重轮询
func NewWeightOrderBalancer(router Router) Balancer {
	var weights []int
	for _, p := range router.Proxy {
		weights = append(weights, p.Weight)
	}
	return &WeightOrderBalancer{
		t:       0,
		weights: weights,
	}
}

type WeightRandomBalancer struct {
	sum     int
	weights []int
}

// NewWeightRandomBalancer 权重随机
func NewWeightRandomBalancer(router Router) Balancer {
	var weights []int
	sum := 0
	for _, p := range router.Proxy {
		weights = append(weights, p.Weight)
		sum += p.Weight
	}
	return &WeightRandomBalancer{
		sum:     sum,
		weights: weights,
	}
}

func (b *WeightRandomBalancer) Index() int {
	sum := 0
	rand.Seed(time.Now().UnixNano())
	randomN := rand.Intn(b.sum)
	for idx, i := range b.weights {
		sum += i
		if randomN < sum {
			return idx
		}
	}
	return -1
}

type ConsistentHashBalancer struct {
	length int
	hash0  int64
}

// NewConsistentHashBalancer 一致哈希
func NewConsistentHashBalancer(router Router) Balancer {
	return &ConsistentHashBalancer{
		length: len(router.Proxy),
		hash0:  time.Now().UnixNano(),
	}
}

func (b *ConsistentHashBalancer) Index() int {
	h := sha256.New()
	bytes := h.Sum([]byte(strconv.FormatInt(b.hash0, 64)))
	i16 := fmt.Sprintf("%X", bytes[:6])
	i, err := strconv.ParseInt(i16, 16, 32)
	if err != nil {
		fmt.Println(err)
	}
	return int(i) % b.length
}
