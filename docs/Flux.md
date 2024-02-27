# Flux

Continuous delivery of applications to the Homecluster is managed using
GitHub and [Flux](https://fluxcd.io/).

## Configuring Flux for the Homecluster

Create a PAT (Personal Access Token) in GitHub with all `repo` permissions.

Set the `GITHUB_TOKEN` environment variable to the PAT.

```shell
export GITHUB_TOKEN=<gh-token>
```

Bootstrap Flux for the Homecluster.

```shell
flux bootstrap github \
  --token-auth \
  --owner=t-eckert \
  --repository=homecluster \
  --branch=main \
  --personal
```

This will configure Flux to watch the state of the `main` branch of the
repository and deploy Kubernetes configurations to the Homecluster.
