package relay

import (
	relaycommon "github.com/QuantumNous/tabbycat/relay/common"
	"github.com/QuantumNous/tabbycat/types"
)

func TabbyCatErrorFromParamOverride(err error) *types.TabbyCatError {
	if fixedErr, ok := relaycommon.AsParamOverrideReturnError(err); ok {
		return relaycommon.TabbyCatErrorFromParamOverride(fixedErr)
	}
	return types.NewError(err, types.ErrorCodeChannelParamOverrideInvalid, types.ErrOptionWithSkipRetry())
}
