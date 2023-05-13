pwd_lower=$(echo $(pwd) | tr '[:upper:]' '[:lower:]' | tr " " "\ ")
# echo "POST||/api/v1/units?layer=newlayer&token=10063865700249539947||||$(cat body.json)" | ./make_ammo.py > ammo.txt
docker run --rm -v ${pwd_lower}:/var/loadtest --name yandex-tank -d -it --net host --entrypoint "/bin/bash" direvius/yandex-tank
docker exec yandex-tank yandex-tank -c const_load.yaml ammo.txt
docker exec yandex-tank yandex-tank -c line_load.yaml ammo.txt
docker stop yandex-tank
