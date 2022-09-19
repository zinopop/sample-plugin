package base

//+kubebuilder:rbac:groups=``,resources=events,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=coordination.k8s.io,resources=leases,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=``,resources=configmaps,verbs=get;list;watch;create;update;patch;delete
// ----- IntegrationClass
//+kubebuilder:rbac:groups=core.katanomi.dev,resources=integrationclasses,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=core.katanomi.dev,resources=integrationclasses/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=core.katanomi.dev,resources=integrationclasses/finalizers,verbs=update
// ----- Webhook
//+kubebuilder:rbac:groups=core.katanomi.dev,resources=webhooks,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=core.katanomi.dev,resources=webhooks/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=core.katanomi.dev,resources=webhooks/finalizers,verbs=update
//+kubebuilder:rbac:groups=``,resources=secrets,verbs=get;list;watch
//+kubebuilder:rbac:groups=eventing.knative.dev,resources=brokers,verbs=get;list;watch
//+kubebuilder:rbac:groups=core.katanomi.dev,resources=integrationclasses,verbs=get;list;watch
// ----- Trigger
//+kubebuilder:rbac:groups=core.katanomi.dev,resources=triggers,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=core.katanomi.dev,resources=triggers/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=core.katanomi.dev,resources=triggers/finalizers,verbs=update
//+kubebuilder:rbac:groups=eventing.knative.dev,resources=brokers,verbs=get;list;watch
//+kubebuilder:rbac:groups=core.katanomi.dev,resources=integrationclasses,verbs=get;list;watch
// ----- GitTrigger
//+kubebuilder:rbac:groups=core.katanomi.dev,resources=gittriggers,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=core.katanomi.dev,resources=gittriggers/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=``,resources=secrets,verbs=get;list;watch
//+kubebuilder:rbac:groups=eventing.knative.dev,resources=brokers,verbs=get;list;watch
//+kubebuilder:rbac:groups=core.katanomi.dev,resources=integrationclasses,verbs=get;list;watch
//+kubebuilder:rbac:groups=core.katanomi.dev,resources=webhooks,verbs=get;list;watch;create;update;patch;delete;deletecollection
//+kubebuilder:rbac:groups=core.katanomi.dev,resources=triggers,verbs=get;list;watch;create;update;patch;delete;deletecollection
//+kubebuilder:rbac:groups=messaging.knative.dev,resources=subscriptions,verbs=get;list;watch
// ----- Events
//+kubebuilder:rbac:groups=``,resources=events,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=coordination.k8s.io,resources=leases,verbs=get;list;watch;create;update;patch;delete
