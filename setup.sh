ssh-keygen -t rsa -b 4096 -C "your_email@example.com"
cat ~/.ssh/id_rsa.pub > /home/ubuntu/key.txt
git config --global user.name "YOUR GITHUB USERNAME"


snap install docker

cat key.txt
