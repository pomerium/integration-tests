{
  apiVersion: 'v1',
  kind: 'List',
  items: [
    {
      apiVersion: 'v1',
      kind: 'Service',
      metadata: {
        namespace: 'default',
        name: 'httpecho',
        labels: { app: 'httpecho' },
      },
      spec: {
        selector: { app: 'httpecho' },
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
        name: 'httpecho',
      },
      spec: {
        replicas: 1,
        selector: { matchLabels: { app: 'httpecho' } },
        template: {
          metadata: {
            labels: { app: 'httpecho' },
          },
          spec: {
            containers: [{
              name: 'httpecho',
              image: 'hashicorp/http-echo',
              imagePullPolicy: 'IfNotPresent',
              args: [
                '-listen=:80',
                '-text=' + |||
                  <!DOCTYPE html>
                  <head>
                    <link rel="shortcut icon" href="data:image/x-icon;," type="image/x-icon">
                  </head>
                  <body>

                  </body>
                |||,
              ],
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
