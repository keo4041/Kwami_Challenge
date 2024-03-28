# AWS CloudFormation Web Server

This project demonstrates how to create and deploy a secure web server using AWS CloudFormation. The web server serves a simple "Hello World" page and is configured to redirect HTTP requests to HTTPS using a self-signed certificate.

## Prerequisites

- AWS Account
- AWS CLI installed and configured

## Deployment

1. Clone this repository:

   ```bash
   git clone https://github.com/keo4041/kwami_challenge.git
   cd kwami_challenge
   ```

2. Deploy the CloudFormation stack:

   ```bash
   aws cloudformation create-stack --stack-name web-server-stack --template-body file://cloud_formation_template.yaml
   ```

   This command will create an EC2 instance with the necessary security group and configurations.

3. Once the stack creation is complete, you can find the URL of the web server in the outputs section of the CloudFormation console.

## Testing

A test script if part of the config YAML and will run after the creation of the instance and the certificate.
If it fails, the logs can be find here /var/log/cloud-init-output.log. You can access them via AWS EC2 Instance Connect
To manually validate the server configuration, you can run the provided test script:

```bash
python test_server.py
```

This script checks if the server redirects HTTP requests to HTTPS with a status code 301 and the redirect url is HTTPS

## Cleanup

```bash
aws cloudformation delete-stack --stack-name web-server-stack
```

## Template Details

The CloudFormation template (`cloud_formation_template.yaml`) contains the following resources:

- **WebServerSecurityGroup**: A security group that allows HTTP and HTTPS traffic and SSH access on port 22, which is necessary for using EC2 Instance Connect.
- **WebServerInstance**: An EC2 instance that runs Apache HTTP Server with SSL support and serves the "Hello World" page as well as a python scrip to test the redirect from HTTP to HTTPS.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
