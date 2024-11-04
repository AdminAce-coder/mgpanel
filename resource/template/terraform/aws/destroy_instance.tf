
data "aws_instance" "to_delete" {
  instance_id = "i-088b668cc6ed86f45"
}

resource "aws_instance" "delete_instance" {
  instance_id = data.aws_instance.to_delete.id
}

output "deleted_instance_id" {
  value = aws_instance.delete_instance.id
}
