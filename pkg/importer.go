package pkg

import (
	corev1alpha1 "github.com/katanomi/core/pkg/apis/core/v1alpha1"
	corectrl "github.com/katanomi/core/pkg/controllers/core"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func GetImportersForPlugin() map[metav1.TypeMeta]corectrl.TypeDecoderMutator {
	return map[metav1.TypeMeta]corectrl.TypeDecoderMutator{
		// Core IntegrationClass
		{
			APIVersion: corev1alpha1.GroupVersion.String(),
			Kind:       corev1alpha1.IntegrationClassKind,
		}: {
			Type:                &corev1alpha1.IntegrationClass{},
			MutatorFunc:         corectrl.NoOpMutation,
			CheckDependencyFunc: corectrl.InvalidPayloadCheckDependency(&corev1alpha1.IntegrationClass{ObjectMeta: metav1.ObjectMeta{Name: "invalid-integration-class"}}), //nolint
		},
	}
}
