requirements:
- ecr repo
- arm64 (or change cpuArchitecture for taskDef)

container build:

```
cd source
docker build -t 403154239519.dkr.ecr.eu-west-1.amazonaws.com/eventbridge-fargate-poc:latest .
docker push 403154239519.dkr.ecr.eu-west-1.amazonaws.com/eventbridge-fargate-poc:latest
```

template deployment:
```
sam build
sam deploy
```

downsides:
- slow Fargate task provision
