resource "aws_lambda_function" "app" {
  s3_bucket         = aws_s3_bucket.lambda_bucket.id
  s3_key            = aws_s3_object.lambda_app.key
  function_name     = var.name
  role              = aws_iam_role.lambda_exec.arn
  handler           = var.name
  source_code_hash = data.archive_file.lambda_app.output_base64sha256
  runtime           = "go1.x"
  memory_size       = 1024
  timeout           = 30
}

resource "aws_lambda_permission" "api_gw" {
  statement_id  = "AllowExecutionFromAPIGateway"
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.app.function_name
  principal     = "apigateway.amazonaws.com"

  source_arn = "${aws_apigatewayv2_api.lambda.execution_arn}/*/*"
}
