pwd_lower=$(echo $(pwd) | tr '[:upper:]' '[:lower:]' | tr " " "\ ")
python3 make_ammo.py > ammo.txt
docker run --rm -v ${pwd_lower}:/var/loadtest --name yandex-tank -d -it --net host --entrypoint "/bin/bash" direvius/yandex-tank
docker exec yandex-tank yandex-tank -c const_load.yaml
docker exec yandex-tank yandex-tank -c line_load.yaml
docker stop yandex-tank
