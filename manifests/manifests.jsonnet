local httpbin = import './lib/httpbin.libsonnet';
local httpdetails = import './lib/httpdetails.libsonnet';
local httpecho = import './lib/httpecho.libsonnet';
local nginxIngressController = import './lib/nginx-ingress-controller.libsonnet';
local pomerium = import './lib/pomerium.libsonnet';
local openid = import './lib/reference-openid-provider.libsonnet';

{
  apiVersion: 'v1',
  kind: 'List',
  items: nginxIngressController.items + pomerium.items + httpbin.items + httpecho.items + openid.items + httpdetails.items,
}
