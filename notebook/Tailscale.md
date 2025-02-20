# Tailscale

I run Tailscale using the Kubernetes operator. I followed the instructions here: https://tailscale.com/kb/1236/kubernetes-operator

I needed to add the following to the Tailscale pods:

``` yaml
requests:
  limits:
    squat.ai/tun: "1"
```

To do this, I need to add a `ProxyClass` resource. This is in the `system/tailscale-proxy-class.yaml` file.

