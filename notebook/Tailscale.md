https://tailscale.com/kb/1236/kubernetes-operator

I need to add the following to the Tailscale pods:

``` yaml
requests:
  limits:
    squat.ai/tun: "1"
```

To do this, I need to add a `ProxyClass` resource.
