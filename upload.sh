docker compose build
aws ecr get-login-password --region ap-northeast-1 --profile=techlinker --region ap-northeast-1  | docker login --username AWS --password-stdin 357238239555.dkr.ecr.ap-northeast-1.amazonaws.com
docker tag line-fortune-telling-image:latest 357238239555.dkr.ecr.ap-northeast-1.amazonaws.com/line-fortune-telling-image:latest
docker push 357238239555.dkr.ecr.ap-northeast-1.amazonaws.com/line-fortune-telling-image:latest
aws lambda update-function-code --function-name line-fortune-telling --image-uri 357238239555.dkr.ecr.ap-northeast-1.amazonaws.com/line-fortune-telling-image:latest --profile=techlinker
