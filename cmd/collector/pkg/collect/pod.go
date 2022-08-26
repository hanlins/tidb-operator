package collect

import (
	"context"

	_ "k8s.io/api/core/v1"
	corev1 "k8s.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type Pod struct {
	*BaseCollector
}

var _ Collector = (*Pod)(nil)

func (p *Pod) Objects() (<-chan client.Object, error) {
	list := &corev1.PodList{}
	err := p.Reader.List(context.Background(), list, p.opts...)
	if err != nil {
		return nil, err
	}

	ch := make(chan client.Object)
	go func() {
		for _, obj := range list.Items {
			ch <- &obj
		}
		close(ch)
	}()
	return ch, nil
}

func NewPodCollector(cli client.Reader) Collector {
	addCorev1Scheme()
	return &Pod{
		BaseCollector: NewBaseCollector(cli),
	}
}
