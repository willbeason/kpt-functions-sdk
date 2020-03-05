import { KubernetesObject } from 'kpt-functions';
import * as apisMetaV1 from './io.k8s.apimachinery.pkg.apis.meta.v1';

export class BackendConfig implements KubernetesObject {
  // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
  public apiVersion: string;

  // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
  public kind: string;

  // Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata
  public metadata: apisMetaV1.ObjectMeta;

  // BackendConfigSpec is the spec for a BackendConfig resource
  public spec?: BackendConfig.Spec;

  public status?: object;

  constructor(desc: BackendConfig.Interface) {
    this.apiVersion = BackendConfig.apiVersion;
    this.kind = BackendConfig.kind;
    this.metadata = desc.metadata;
    this.spec = desc.spec;
    this.status = desc.status;
  }
}

export function isBackendConfig(o: any): o is BackendConfig {
  return o && o.apiVersion === BackendConfig.apiVersion && o.kind === BackendConfig.kind;
}

export namespace BackendConfig {
  export const apiVersion = "cloud.google.com/v1beta1";
  export const group = "cloud.google.com";
  export const version = "v1beta1";
  export const kind = "BackendConfig";

  // named constructs a BackendConfig with metadata.name set to name.
  export function named(name: string): BackendConfig {
    return new BackendConfig({metadata: {name}});
  }
  export interface Interface {
    // Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata
    metadata: apisMetaV1.ObjectMeta;

    // BackendConfigSpec is the spec for a BackendConfig resource
    spec?: BackendConfig.Spec;

    status?: object;
  }
  // BackendConfigSpec is the spec for a BackendConfig resource
  export class Spec {
    // CDNConfig contains configuration for CDN-enabled backends.
    public cdn?: BackendConfig.Spec.Cdn;

    // ConnectionDrainingConfig contains configuration for connection draining. For now the draining timeout. May manage more settings in the future.
    public connectionDraining?: BackendConfig.Spec.ConnectionDraining;

    // IAPConfig contains configuration for IAP-enabled backends.
    public iap?: BackendConfig.Spec.Iap;

    public securityPolicy?: object;

    // SessionAffinityConfig contains configuration for stickyness parameters.
    public sessionAffinity?: BackendConfig.Spec.SessionAffinity;

    public timeoutSec?: number;
  }

  export namespace Spec {
    // CDNConfig contains configuration for CDN-enabled backends.
    export class Cdn {
      // CacheKeyPolicy contains configuration for how requests to a CDN-enabled backend are cached.
      public cachePolicy?: BackendConfig.Spec.Cdn.CachePolicy;

      public enabled: boolean;

      constructor(desc: BackendConfig.Spec.Cdn) {
        this.cachePolicy = desc.cachePolicy;
        this.enabled = desc.enabled;
      }
    }

    export namespace Cdn {
      // CacheKeyPolicy contains configuration for how requests to a CDN-enabled backend are cached.
      export class CachePolicy {
        // If true, requests to different hosts will be cached separately.
        public includeHost?: boolean;

        // If true, http and https requests will be cached separately.
        public includeProtocol?: boolean;

        // If true, query string parameters are included in the cache key according to QueryStringBlacklist and QueryStringWhitelist. If neither is set, the entire query string is included and if false the entire query string is excluded.
        public includeQueryString?: boolean;

        // Names of query strint parameters to exclude from cache keys. All other parameters are included. Either specify QueryStringBlacklist or QueryStringWhitelist, but not both.
        public queryStringBlacklist?: string[];

        // Names of query string parameters to include in cache keys. All other parameters are excluded. Either specify QueryStringBlacklist or QueryStringWhitelist, but not both.
        public queryStringWhitelist?: string[];
      }
    }
    // ConnectionDrainingConfig contains configuration for connection draining. For now the draining timeout. May manage more settings in the future.
    export class ConnectionDraining {
      // Draining timeout in seconds.
      public drainingTimeoutSec?: number;
    }
    // IAPConfig contains configuration for IAP-enabled backends.
    export class Iap {
      public enabled: boolean;

      // OAuthClientCredentials contains credentials for a single IAP-enabled backend.
      public oauthclientCredentials: BackendConfig.Spec.Iap.OauthclientCredentials;

      constructor(desc: BackendConfig.Spec.Iap) {
        this.enabled = desc.enabled;
        this.oauthclientCredentials = desc.oauthclientCredentials;
      }
    }

    export namespace Iap {
      // OAuthClientCredentials contains credentials for a single IAP-enabled backend.
      export class OauthclientCredentials {
        // Direct reference to OAuth client id.
        public clientID?: string;

        // Direct reference to OAuth client secret.
        public clientSecret?: string;

        // The name of a k8s secret which stores the OAuth client id & secret.
        public secretName: string;

        constructor(desc: BackendConfig.Spec.Iap.OauthclientCredentials) {
          this.clientID = desc.clientID;
          this.clientSecret = desc.clientSecret;
          this.secretName = desc.secretName;
        }
      }
    }
    // SessionAffinityConfig contains configuration for stickyness parameters.
    export class SessionAffinity {
      public affinityCookieTtlSec?: number;

      public affinityType?: string;
    }
  }
}

// BackendConfigList is a list of BackendConfig
export class BackendConfigList {
  // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
  public apiVersion: string;

  // List of backendconfigs. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md
  public items: BackendConfig[];

  // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
  public kind: string;

  // ListMeta describes metadata that synthetic resources must have, including lists and various status objects. A resource may have only one of {ObjectMeta, ListMeta}.
  public metadata?: BackendConfigList.Metadata;

  constructor(desc: BackendConfigList) {
    this.apiVersion = BackendConfigList.apiVersion;
    this.items = desc.items.map((i) => new BackendConfig(i));
    this.kind = BackendConfigList.kind;
    this.metadata = desc.metadata;
  }
}

export function isBackendConfigList(o: any): o is BackendConfigList {
  return o && o.apiVersion === BackendConfigList.apiVersion && o.kind === BackendConfigList.kind;
}

export namespace BackendConfigList {
  export const apiVersion = "cloud.google.com/v1beta1";
  export const group = "cloud.google.com";
  export const version = "v1beta1";
  export const kind = "BackendConfigList";

  // BackendConfigList is a list of BackendConfig
  export interface Interface {
    // List of backendconfigs. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md
    items: BackendConfig[];

    // ListMeta describes metadata that synthetic resources must have, including lists and various status objects. A resource may have only one of {ObjectMeta, ListMeta}.
    metadata?: BackendConfigList.Metadata;
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