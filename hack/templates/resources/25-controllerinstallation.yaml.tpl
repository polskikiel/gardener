<%
  import os, yaml

  values={}
  if context.get("values", "") != "":
    values=yaml.load(open(context.get("values", "")), Loader=yaml.Loader)

  def value(path, default):
    keys=str.split(path, ".")
    root=values
    for key in keys:
      if isinstance(root, dict):
        if key in root:
          root=root[key]
        else:
          return default
      else:
        return default
    return root
%># ControllerInstallation object requests a specific controller to get deployed to a seed cluster.
# See https://github.com/gardener/gardener/blob/master/docs/proposals/01-extensibility.md.
---
apiVersion: core.gardener.cloud/v1beta1
kind: ControllerInstallation
metadata:
  name: ${value("metadata.name", "os-coreos")}<% annotations = value("metadata.annotations", {}); labels = value("metadata.labels", {}) %>
  % if annotations != {}:
  annotations: ${yaml.dump(annotations, width=10000, default_flow_style=None)}
  % endif
  % if labels != {}:
  labels: ${yaml.dump(labels, width=10000, default_flow_style=None)}
  % endif
spec:
  registrationRef:
    apiVersion: core.gardener.cloud/v1beta1
    kind: ControllerRegistration
    name: ${value("metadata.name", "os-coreos")}
  seedRef:
    apiVersion: core.gardener.cloud/v1beta1
    kind: Seed
    name: my-seed
