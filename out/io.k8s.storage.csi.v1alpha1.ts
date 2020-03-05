import { KubernetesObject } from 'kpt-functions';
import * as apisMetaV1 from './io.k8s.apimachinery.pkg.apis.meta.v1';

export class CSIDriver implements KubernetesObject {
  // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
  public apiVersion: string;

  // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
  public kind: string;

  // Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata
  public metadata: apisMetaV1.ObjectMeta;

  // Specification of the CSI Driver.
  public spec?: object;

  constructor(desc: CSIDriver.Interface) {
    this.apiVersion = CSIDriver.apiVersion;
    this.kind = CSIDriver.kind;
    this.metadata = desc.metadata;
    this.spec = desc.spec;
  }
}

export function isCSIDriver(o: any): o is CSIDriver {
  return o && o.apiVersion === CSIDriver.apiVersion && o.kind === CSIDriver.kind;
}

export namespace CSIDriver {
  export const apiVersion = "csi.storage.k8s.io/v1alpha1";
  export const group = "csi.storage.k8s.io";
  export const version = "v1alpha1";
  export const kind = "CSIDriver";

  // named constructs a CSIDriver with metadata.name set to name.
  export function named(name: string): CSIDriver {
    return new CSIDriver({metadata: {name}});
  }
  export interface Interface {
    // Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata
    metadata: apisMetaV1.ObjectMeta;

    // Specification of the CSI Driver.
    spec?: object;
  }
}

// CSIDriverList is a list of CSIDriver
export class CSIDriverList {
  // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
  public apiVersion: string;

  // List of csidrivers. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md
  public items: CSIDriver[];

  // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
  public kind: string;

  // ListMeta describes metadata that synthetic resources must have, including lists and various status objects. A resource may have only one of {ObjectMeta, ListMeta}.
  public metadata?: CSIDriverList.Metadata;

  constructor(desc: CSIDriverList) {
    this.apiVersion = CSIDriverList.apiVersion;
    this.items = desc.items.map((i) => new CSIDriver(i));
    this.kind = CSIDriverList.kind;
    this.metadata = desc.metadata;
  }
}

export function isCSIDriverList(o: any): o is CSIDriverList {
  return o && o.apiVersion === CSIDriverList.apiVersion && o.kind === CSIDriverList.kind;
}

export namespace CSIDriverList {
  export const apiVersion = "csi.storage.k8s.io/v1alpha1";
  export const group = "csi.storage.k8s.io";
  export const version = "v1alpha1";
  export const kind = "CSIDriverList";

  // CSIDriverList is a list of CSIDriver
  export interface Interface {
    // List of csidrivers. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md
    items: CSIDriver[];

    // ListMeta describes metadata that synthetic resources must have, including lists and various status objects. A resource may have only one of {ObjectMeta, ListMeta}.
    metadata?: CSIDriverList.Metadata;
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

export class CSINodeInfo implements KubernetesObject {
  // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
  public apiVersion: string;

  // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
  public kind: string;

  // Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata
  public metadata: apisMetaV1.ObjectMeta;

  // Specification of CSINodeInfo
  public spec?: object;

  // Status of CSINodeInfo
  public status?: object;

  constructor(desc: CSINodeInfo.Interface) {
    this.apiVersion = CSINodeInfo.apiVersion;
    this.kind = CSINodeInfo.kind;
    this.metadata = desc.metadata;
    this.spec = desc.spec;
    this.status = desc.status;
  }
}

export function isCSINodeInfo(o: any): o is CSINodeInfo {
  return o && o.apiVersion === CSINodeInfo.apiVersion && o.kind === CSINodeInfo.kind;
}

export namespace CSINodeInfo {
  export const apiVersion = "csi.storage.k8s.io/v1alpha1";
  export const group = "csi.storage.k8s.io";
  export const version = "v1alpha1";
  export const kind = "CSINodeInfo";

  // named constructs a CSINodeInfo with metadata.name set to name.
  export function named(name: string): CSINodeInfo {
    return new CSINodeInfo({metadata: {name}});
  }
  export interface Interface {
    // Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata
    metadata: apisMetaV1.ObjectMeta;

    // Specification of CSINodeInfo
    spec?: object;

    // Status of CSINodeInfo
    status?: object;
  }
}

// CSINodeInfoList is a list of CSINodeInfo
export class CSINodeInfoList {
  // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
  public apiVersion: string;

  // List of csinodeinfos. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md
  public items: CSINodeInfo[];

  // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
  public kind: string;

  // ListMeta describes metadata that synthetic resources must have, including lists and various status objects. A resource may have only one of {ObjectMeta, ListMeta}.
  public metadata?: CSINodeInfoList.Metadata;

  constructor(desc: CSINodeInfoList) {
    this.apiVersion = CSINodeInfoList.apiVersion;
    this.items = desc.items.map((i) => new CSINodeInfo(i));
    this.kind = CSINodeInfoList.kind;
    this.metadata = desc.metadata;
  }
}

export function isCSINodeInfoList(o: any): o is CSINodeInfoList {
  return o && o.apiVersion === CSINodeInfoList.apiVersion && o.kind === CSINodeInfoList.kind;
}

export namespace CSINodeInfoList {
  export const apiVersion = "csi.storage.k8s.io/v1alpha1";
  export const group = "csi.storage.k8s.io";
  export const version = "v1alpha1";
  export const kind = "CSINodeInfoList";

  // CSINodeInfoList is a list of CSINodeInfo
  export interface Interface {
    // List of csinodeinfos. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md
    items: CSINodeInfo[];

    // ListMeta describes metadata that synthetic resources must have, including lists and various status objects. A resource may have only one of {ObjectMeta, ListMeta}.
    metadata?: CSINodeInfoList.Metadata;
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