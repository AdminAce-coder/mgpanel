provider "aws" {
  region = "us-east-1"  # 替换为实际的区域
}

data "aws_vpc" "selected" {
  filter {
    name   = "vpc-id"  # 按照 VPC ID 查询
    values = ["vpc-0464b5baebfcaa821"]  # 替换为实际的 VPC ID
  }
}

data "aws_subnet" "selected" {
  filter {
    name   = "subnet-id"  # 按照子网 ID 查询
    values = ["subnet-0711e346568ae125b"]  # 替换为实际的子网 ID
  }
  vpc_id = data.aws_vpc.selected.id
}

data "aws_security_group" "selected" {
  filter {
    name   = "group-id"  # 按照安全组 ID 查询
    values = ["sg-0d51cce9c33fa3221"]  # 替换为实际的安全组 ID
  }
}

# 使用指定的 AMI ID 启动实例
resource "aws_instance" "web_server" {
  ami           = "ami-0866a3c8686eaeeba"  # 直接使用 AMI ID
  instance_type = "t2.micro"  # 根据需求调整实例类型

  key_name = "zxb"  # 替换为实际的密钥名称

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
