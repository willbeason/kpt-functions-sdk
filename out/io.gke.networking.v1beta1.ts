import { KubernetesObject } from 'kpt-functions';
import * as apisMetaV1 from './io.k8s.apimachinery.pkg.apis.meta.v1';

export class ManagedCertificate implements KubernetesObject {
  // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
  public apiVersion: string;

  // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
  public kind: string;

  // Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata
  public metadata: apisMetaV1.ObjectMeta;

  public spec?: object;

  public status?: object;

  constructor(desc: ManagedCertificate.Interface) {
    this.apiVersion = ManagedCertificate.apiVersion;
    this.kind = ManagedCertificate.kind;
    this.metadata = desc.metadata;
    this.spec = desc.spec;
    this.status = desc.status;
  }
}

export function isManagedCertificate(o: any): o is ManagedCertificate {
  return o && o.apiVersion === ManagedCertificate.apiVersion && o.kind === ManagedCertificate.kind;
}

export namespace ManagedCertificate {
  export const apiVersion = "networking.gke.io/v1beta1";
  export const group = "networking.gke.io";
  export const version = "v1beta1";
  export const kind = "ManagedCertificate";

  // named constructs a ManagedCertificate with metadata.name set to name.
  export function named(name: string): ManagedCertificate {
    return new ManagedCertificate({metadata: {name}});
  }
  export interface Interface {
    // Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata
    metadata: apisMetaV1.ObjectMeta;

    spec?: object;

    status?: object;
  }
}

// ManagedCertificateList is a list of ManagedCertificate
export class ManagedCertificateList {
  // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
  public apiVersion: string;

  // List of managedcertificates. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md
  public items: ManagedCertificate[];

  // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
  public kind: string;

  // ListMeta describes metadata that synthetic resources must have, including lists and various status objects. A resource may have only one of {ObjectMeta, ListMeta}.
  public metadata?: ManagedCertificateList.Metadata;

  constructor(desc: ManagedCertificateList) {
    this.apiVersion = ManagedCertificateList.apiVersion;
    this.items = desc.items.map((i) => new ManagedCertificate(i));
    this.kind = ManagedCertificateList.kind;
    this.metadata = desc.metadata;
  }
}

export function isManagedCertificateList(o: any): o is ManagedCertificateList {
  return o && o.apiVersion === ManagedCertificateList.apiVersion && o.kind === ManagedCertificateList.kind;
}

export namespace ManagedCertificateList {
  export const apiVersion = "networking.gke.io/v1beta1";
  export const group = "networking.gke.io";
  export const version = "v1beta1";
  export const kind = "ManagedCertificateList";

  // ManagedCertificateList is a list of ManagedCertificate
  export interface Interface {
    // List of managedcertificates. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md
    items: ManagedCertificate[];

    // ListMeta describes metadata that synthetic resources must have, including lists and various status objects. A resource may have only one of {ObjectMeta, ListMeta}.
    metadata?: ManagedCertificateList.Metadata;
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