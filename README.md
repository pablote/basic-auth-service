# basic-auth-service

Basic Auth as an External Authorization service for Envoy Proxy

## Configuration

All configuration is done using env vars. Check `.env` file for a sample configuration.

* PORT (optional, defaults to `10000`): Listen port.
* BASIC_AUTH_SERVICE_HTPASSWD: List of users in Apache Htpasswd format, remember to escape all `$` characters. If defined, this has preference over username/password. See `.env` for an example.
* BASIC_AUTH_SERVICE_USERNAME: Basic authentication username.
* BASIC_AUTH_SERVICE_PASSWORD: Basic authentication password.
* BASIC_AUTH_SERVICE_HOST_ALLOWLIST (optional, defaults to `*`): Comma separated list of hosts, on glob format, to which authentication will be applied. If the incoming host matches any it will be processed, if not the call will be allowed.
* BASIC_AUTH_SERVICE_PATH_ALLOWLIST (optional, defaults to `*`): Comma separated list of paths, on glob format, to which authentication will be applied. If the incoming path matches any it will be processed, if not the call will be allowed.

Sample Envoy config to use this service:

* On `http_filters`:

```yaml
  - name: envoy.filters.http.ext_authz
    typed_config:
      "@type": type.googleapis.com/envoy.extensions.filters.http.ext_authz.v3.ExtAuthz
      failure_mode_allow: false
      http_service:
        server_uri:
          uri: basic-auth-service:10000
          cluster: ext-authz
          timeout: 0.25s
        authorization_response:
          allowed_upstream_headers:
            patterns:
              - exact: "x-auth-username"
      include_peer_certificate: true
      with_request_body:
        max_request_bytes: 1024
        allow_partial_message: true
        pack_as_bytes: true
      transport_api_version: V3
```

* On `clusters`:

```yaml
- name: ext-authz
  connect_timeout: 0.25s
  type: logical_dns
  lb_policy: round_robin
  load_assignment:
    cluster_name: ext-authz
    endpoints:
    - lb_endpoints:
      - endpoint:
          address:
            socket_address:
              address: basic-auth-service
              port_value: 10000
```

## Docker

This service is available as a docker image at [pablote/basic-auth-service](https://hub.docker.com/r/pablote/basic-auth-service)
