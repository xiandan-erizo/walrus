package convertor

import (
	"fmt"

	"k8s.io/apimachinery/pkg/util/sets"

	"github.com/seal-io/seal/pkg/dao/model"
	"github.com/seal-io/seal/pkg/platformtf/block"
)

const (
	TypeKubernetes = "kubernetes"
	TypeHelm       = "helm"
	TypeKubectl    = "kubectl"
	// Add more convertor type.
)

type ConvertOptions struct {
	SecretMountPath string
	ConnSeparator   string
	Providers       []string
}

// LoadConvertor loads the convertor by the provider type.
func LoadConvertor(provider string) Convertor {
	switch provider {
	case TypeKubernetes:
		return K8sConvertor(provider)
	case TypeHelm:
		return HelmConvertor(provider)
	case TypeKubectl:
		return KubectlConvertor(provider)
	default:
		return DefaultConvertor(provider)
	}
}

// ToProvidersBlocks converts the connectors to provider blocks with required providers.
func ToProvidersBlocks(providers []string, connectors model.Connectors, opts ConvertOptions) (blocks block.Blocks, err error) {
	var builtinProviders = []string{
		TypeKubernetes,
		TypeHelm,
	}

	for _, p := range providers {
		var convertBlocks block.Blocks
		convertBlocks, err = ToProviderBlocks(p, connectors, opts)
		if err != nil {
			return nil, err
		}
		if convertBlocks == nil {
			continue
		}
		blocks = append(blocks, convertBlocks...)
	}

	//  required custom provider that could convert from connectors is validated when converting.
	// otherwise it will return error. Here we only validate the builtin providers.
	ok, err := validateRequiredProviders(providers, builtinProviders, blocks)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, fmt.Errorf("failed to validate providers: %v, current blockProviders: %v", providers, blocks)
	}

	return blocks, nil
}

// ToProviderBlocks converts the connectors to blocks with provider name.
func ToProviderBlocks(provider string, connectors model.Connectors, opts ConvertOptions) (block.Blocks, error) {
	var toBlockOpts Options
	switch provider {
	case TypeKubernetes, TypeHelm, TypeKubectl:
		toBlockOpts = K8sConvertorOptions{
			ConfigPath:    opts.SecretMountPath,
			ConnSeparator: opts.ConnSeparator,
		}
	}

	return LoadConvertor(provider).ToBlocks(connectors, toBlockOpts)
}

// validateRequiredProviders the providers both in providers and supportedProviders
// need to be checked in the generated blocks.
func validateRequiredProviders(providers, supportedProviders []string, blocks block.Blocks) (bool, error) {
	// blockProviders is the providers of the generated blocks.
	blockProviders, err := blocks.GetProviderNames()
	if err != nil {
		return false, err
	}
	actual := sets.NewString(blockProviders...)

	// supported is the providers which convertors can generate with the given connectors.
	// custom providers are the providers which are generated by the primary provider.
	supported := sets.NewString(supportedProviders...)

	// get the intersection of the required and supported providers to validate.
	required := sets.NewString(providers...)
	expected := required.Intersection(supported)

	return actual.IsSuperset(expected), nil
}
