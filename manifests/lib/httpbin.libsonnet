{
  apiVersion: 'v1',
  kind: 'List',
  items: [
    {
      apiVersion: 'v1',
      kind: 'Service',
      metadata: {
        namespace: 'default',
        name: 'httpbin',
        labels: { app: 'httpbin' },
      },
      spec: {
        selector: { app: 'httpbin' },
        ports: [{
          name: 'http',
          port: 80,
          targetPort: 'http',
        }],
      },
    },
    {
      apiVersion: 'apps/v1',
      kind: 'Deployment',
      metadata: {
        namespace: 'default',
        name: 'httpbin',
      },
      spec: {
        replicas: 1,
        selector: { matchLabels: { app: 'httpbin' } },
        template: {
          metadata: {
            labels: { app: 'httpbin' },
          },
          spec: {
            containers: [{
              name: 'httpbin',
              image: 'kennethreitz/httpbin',
              imagePullPolicy: 'IfNotPresent',
              ports: [{
                name: 'http',
                containerPort: 80,
              }],
            }],
          },
        },
      },
    },
  ],
}
