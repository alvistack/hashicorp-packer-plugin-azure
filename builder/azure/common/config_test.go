package common_test

import (
	"context"
	"testing"

	"github.com/hashicorp/packer-plugin-sdk/multistep"
	"github.com/stretchr/testify/assert"

	"github.com/hashicorp/packer-plugin-azure/builder/azure/common"
)

func TestSkipCreateImage(t *testing.T) {
	var said []string

	say := func(what string) {
		said = append(said, what)
	}

	config := common.Config{}
	message := "Capture Image"

	steps := config.CaptureSteps(say, common.NewStepNotify(message, say))
	state := &multistep.BasicStateBag{}

	ctx := context.Background()

	for _, step := range steps {
		step.Run(ctx, state)
	}

	assert.Equal(t, said, []string{message})

	said = nil
	config.SkipCreateImage = true

	steps = config.CaptureSteps(say, common.NewStepNotify(message, say))

	for _, step := range steps {
		step.Run(ctx, state)
	}

	assert.Equal(t, said, []string{common.SkippingImageCreation})
}
