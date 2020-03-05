export type ScalingPolicy = object;

// ScalingPolicyList is a list of ScalingPolicy
export class ScalingPolicyList {
  // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
  public apiVersion: string;

  // List of scalingpolicies. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md
  public items: ScalingPolicy[];

  // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
  public kind: string;

  // ListMeta describes metadata that synthetic resources must have, including lists and various status objects. A resource may have only one of {ObjectMeta, ListMeta}.
  public metadata?: ScalingPolicyList.Metadata;

  constructor(desc: ScalingPolicyList) {
    this.apiVersion = ScalingPolicyList.apiVersion;
    this.items = desc.items;
    this.kind = ScalingPolicyList.kind;
    this.metadata = desc.metadata;
  }
}

export function isScalingPolicyList(o: any): o is ScalingPolicyList {
  return o && o.apiVersion === ScalingPolicyList.apiVersion && o.kind === ScalingPolicyList.kind;
}

export namespace ScalingPolicyList {
  export const apiVersion = "scalingpolicy.kope.io/v1alpha1";
  export const group = "scalingpolicy.kope.io";
  export const version = "v1alpha1";
  export const kind = "ScalingPolicyList";

  // ScalingPolicyList is a list of ScalingPolicy
  export interface Interface {
    // List of scalingpolicies. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md
    items: ScalingPolicy[];

    // ListMeta describes metadata that synthetic resources must have, including lists and various status objects. A resource may have only one of {ObjectMeta, ListMeta}.
    metadata?: ScalingPolicyList.Metadata;
  }
  // ListMeta describes metadata that synthetic resources must have, including lists and various status objects. A resource may have only one of {ObjectMeta, ListMeta}.
  export class Metadata {
    // continue may be set if the user set a limit on the number of items returned, and indicates that the server has more data available. The value is opaque and may be used to issue another request to the endpoint that served this list to retrieve the next set of available objects. Continuing a consistent list may not be possible if the server configuration has changed or more than a few minutes have passed. The resourceVersion field returned when using this continue value will be identical to the value in the first response, unless you have received this token from an error message.
    public continue?: string;

    // String that identifies the server's internal version of this object that can be used by clients to determine when objects have changed. Value must be treated as opaque by clients and passed unmodified back to the server. Populated by the system. Read-only. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#concurrency-control-and-consistency
    public resourceVersion?: string;

    // selfLink is a URL representing this object. Populated by the system. Read-only.
    public selfLink?: string;
  }
}