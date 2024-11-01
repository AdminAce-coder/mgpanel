provider "aws" {
  region = "us-east-1"  # 使用 Go 模板渲染区域
}

data "aws_vpc" "selected" {
  filter {
    name   = "vpc-id"
    values = ["vpc-0464b5baebfcaa821"]  # 使用 Go 模板渲染 VPC ID
  }
}

data "aws_subnet" "selected" {
  filter {
    name   = "subnet-id"
    values = ["subnet-0711e346568ae125b"]  # 使用 Go 模板渲染子网 ID
  }
  vpc_id = data.aws_vpc.selected.id
}

data "aws_security_group" "selected" {
  filter {
    name   = "group-id"
    values = ["sg-0d51cce9c33fa3221"]  # 使用 Go 模板渲染安全组 ID
  }
}

resource "aws_instance" "web_server" {
  ami           = "ami-0866a3c8686eaeeba"  # 使用 Go 模板渲染 AMI ID
  instance_type = "t2.micro"  # 使用 Go 模板渲染实例类型

  key_name = "zxb"  # 使用 Go 模板渲染密钥名称

  subnet_id = data.aws_subnet.selected.id

  vpc_security_group_ids = [data.aws_security_group.selected.id]

  root_block_device {
    volume_size = 20
    volume_type = "gp3"
  }

  tags = {
    Name = "MyEC2Instance"
  }
}

output "instance_id" {
  value = aws_instance.web_server.id
}
