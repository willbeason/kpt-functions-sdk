import { KubernetesObject } from 'kpt-functions';
import * as apiCoreV1 from './io.k8s.api.core.v1';
import * as apisMetaV1 from './io.k8s.apimachinery.pkg.apis.meta.v1';

// PodPreset is a policy resource that defines additional runtime requirements for a Pod.
export class PodPreset implements KubernetesObject {
  // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
  public apiVersion: string;

  // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
  public kind: string;

  public metadata: apisMetaV1.ObjectMeta;

  public spec?: PodPresetSpec;

  constructor(desc: PodPreset.Interface) {
    this.apiVersion = PodPreset.apiVersion;
    this.kind = PodPreset.kind;
    this.metadata = desc.metadata;
    this.spec = desc.spec;
  }
}

export function isPodPreset(o: any): o is PodPreset {
  return o && o.apiVersion === PodPreset.apiVersion && o.kind === PodPreset.kind;
}

export namespace PodPreset {
  export const apiVersion = "settings.k8s.io/v1alpha1";
  export const group = "settings.k8s.io";
  export const version = "v1alpha1";
  export const kind = "PodPreset";

  // named constructs a PodPreset with metadata.name set to name.
  export function named(name: string): PodPreset {
    return new PodPreset({metadata: {name}});
  }
  // PodPreset is a policy resource that defines additional runtime requirements for a Pod.
  export interface Interface {
    metadata: apisMetaV1.ObjectMeta;

    spec?: PodPresetSpec;
  }
}

// PodPresetList is a list of PodPreset objects.
export class PodPresetList {
  // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
  public apiVersion: string;

  // Items is a list of schema objects.
  public items: PodPreset[];

  // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
  public kind: string;

  // Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
  public metadata?: apisMetaV1.ListMeta;

  constructor(desc: PodPresetList) {
    this.apiVersion = PodPresetList.apiVersion;
    this.items = desc.items.map((i) => new PodPreset(i));
    this.kind = PodPresetList.kind;
    this.metadata = desc.metadata;
  }
}

export function isPodPresetList(o: any): o is PodPresetList {
  return o && o.apiVersion === PodPresetList.apiVersion && o.kind === PodPresetList.kind;
}

export namespace PodPresetList {
  export const apiVersion = "settings.k8s.io/v1alpha1";
  export const group = "settings.k8s.io";
  export const version = "v1alpha1";
  export const kind = "PodPresetList";

  // PodPresetList is a list of PodPreset objects.
  export interface Interface {
    // Items is a list of schema objects.
    items: PodPreset[];

    // Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
    metadata?: apisMetaV1.ListMeta;
  }
}

// PodPresetSpec is a description of a pod preset.
export class PodPresetSpec {
  // Env defines the collection of EnvVar to inject into containers.
  public env?: apiCoreV1.EnvVar[];

  // EnvFrom defines the collection of EnvFromSource to inject into containers.
  public envFrom?: apiCoreV1.EnvFromSource[];

  // Selector is a label query over a set of resources, in this case pods. Required.
  public selector?: apisMetaV1.LabelSelector;

  // VolumeMounts defines the collection of VolumeMount to inject into containers.
  public volumeMounts?: apiCoreV1.VolumeMount[];

  // Volumes defines the collection of Volume to inject into the pod.
  public volumes?: apiCoreV1.Volume[];
}