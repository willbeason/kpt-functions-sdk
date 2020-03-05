import { KubernetesObject } from 'kpt-functions';
import * as apisMetaV1 from './io.k8s.apimachinery.pkg.apis.meta.v1';

// AuditSink represents a cluster level audit sink
export class AuditSink implements KubernetesObject {
  // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
  public apiVersion: string;

  // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
  public kind: string;

  public metadata: apisMetaV1.ObjectMeta;

  // Spec defines the audit configuration spec
  public spec?: AuditSinkSpec;

  constructor(desc: AuditSink.Interface) {
    this.apiVersion = AuditSink.apiVersion;
    this.kind = AuditSink.kind;
    this.metadata = desc.metadata;
    this.spec = desc.spec;
  }
}

export function isAuditSink(o: any): o is AuditSink {
  return o && o.apiVersion === AuditSink.apiVersion && o.kind === AuditSink.kind;
}

export namespace AuditSink {
  export const apiVersion = "auditregistration.k8s.io/v1alpha1";
  export const group = "auditregistration.k8s.io";
  export const version = "v1alpha1";
  export const kind = "AuditSink";

  // named constructs a AuditSink with metadata.name set to name.
  export function named(name: string): AuditSink {
    return new AuditSink({metadata: {name}});
  }
  // AuditSink represents a cluster level audit sink
  export interface Interface {
    metadata: apisMetaV1.ObjectMeta;

    // Spec defines the audit configuration spec
    spec?: AuditSinkSpec;
  }
}

// AuditSinkList is a list of AuditSink items.
export class AuditSinkList {
  // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
  public apiVersion: string;

  // List of audit configurations.
  public items: AuditSink[];

  // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
  public kind: string;

  public metadata?: apisMetaV1.ListMeta;

  constructor(desc: AuditSinkList) {
    this.apiVersion = AuditSinkList.apiVersion;
    this.items = desc.items.map((i) => new AuditSink(i));
    this.kind = AuditSinkList.kind;
    this.metadata = desc.metadata;
  }
}

export function isAuditSinkList(o: any): o is AuditSinkList {
  return o && o.apiVersion === AuditSinkList.apiVersion && o.kind === AuditSinkList.kind;
}

export namespace AuditSinkList {
  export const apiVersion = "auditregistration.k8s.io/v1alpha1";
  export const group = "auditregistration.k8s.io";
  export const version = "v1alpha1";
  export const kind = "AuditSinkList";

  // AuditSinkList is a list of AuditSink items.
  export interface Interface {
    // List of audit configurations.
    items: AuditSink[];

    metadata?: apisMetaV1.ListMeta;
  }
}

// AuditSinkSpec holds the spec for the audit sink
export class AuditSinkSpec {
  // Policy defines the policy for selecting which events should be sent to the webhook required
  public policy: Policy;

  // Webhook to send events required
  public webhook: Webhook;

  constructor(desc: AuditSinkSpec) {
    this.policy = desc.policy;
    this.webhook = desc.webhook;
  }
}

// Policy defines the configuration of how audit events are logged
export class Policy {
  // The Level that all requests are recorded at. available options: None, Metadata, Request, RequestResponse required
  public level: string;

  // Stages is a list of stages for which events are created.
  public stages?: string[];

  constructor(desc: Policy) {
    this.level = desc.level;
    this.stages = desc.stages;
  }
}

// ServiceReference holds a reference to Service.legacy.k8s.io
export class ServiceReference {
  // `name` is the name of the service. Required
  public name: string;

  // `namespace` is the namespace of the service. Required
  public namespace: string;

  // `path` is an optional URL path which will be sent in any request to this service.
  public path?: string;

  constructor(desc: ServiceReference) {
    this.name = desc.name;
    this.namespace = desc.namespace;
    this.path = desc.path;
  }
}

// Webhook holds the configuration of the webhook
export class Webhook {
  // ClientConfig holds the connection parameters for the webhook required
  public clientConfig: WebhookClientConfig;

  // Throttle holds the options for throttling the webhook
  public throttle?: WebhookThrottleConfig;

  constructor(desc: Webhook) {
    this.clientConfig = desc.clientConfig;
    this.throttle = desc.throttle;
  }
}

// WebhookClientConfig contains the information to make a connection with the webhook
export class WebhookClientConfig {
  // `caBundle` is a PEM encoded CA bundle which will be used to validate the webhook's server certificate. If unspecified, system trust roots on the apiserver are used.
  public caBundle?: string;

  // `service` is a reference to the service for this webhook. Either `service` or `url` must be specified.
  // 
  // If the webhook is running within the cluster, then you should use `service`.
  // 
  // Port 443 will be used if it is open, otherwise it is an error.
  public service?: ServiceReference;

  // `url` gives the location of the webhook, in standard URL form (`scheme://host:port/path`). Exactly one of `url` or `service` must be specified.
  // 
  // The `host` should not refer to a service running in the cluster; use the `service` field instead. The host might be resolved via external DNS in some apiservers (e.g., `kube-apiserver` cannot resolve in-cluster DNS as that would be a layering violation). `host` may also be an IP address.
  // 
  // Please note that using `localhost` or `127.0.0.1` as a `host` is risky unless you take great care to run this webhook on all hosts which run an apiserver which might need to make calls to this webhook. Such installs are likely to be non-portable, i.e., not easy to turn up in a new cluster.
  // 
  // The scheme must be "https"; the URL must begin with "https://".
  // 
  // A path is optional, and if present may be any string permissible in a URL. You may use the path to pass an arbitrary string to the webhook, for example, a cluster identifier.
  // 
  // Attempting to use a user or basic auth e.g. "user:password@" is not allowed. Fragments ("#...") and query parameters ("?...") are not allowed, either.
  public url?: string;
}

// WebhookThrottleConfig holds the configuration for throttling events
export class WebhookThrottleConfig {
  // ThrottleBurst is the maximum number of events sent at the same moment default 15 QPS
  public burst?: number;

  // ThrottleQPS maximum number of batches per second default 10 QPS
  public qps?: number;
}