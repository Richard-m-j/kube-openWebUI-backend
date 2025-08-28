ssh-keygen -t rsa -b 4096 -C "your_email@example.com"
cat ~/.ssh/id_rsa.pub > /home/ubuntu/key.txt
git config --global user.name "YOUR GITHUB USERNAME"

git clone git@github.com:Richard-m-j/kube-openWebUI-backend.git
git clone git@github.com:Richard-m-j/kube-open-webui-frontend.git
git clone git@github.com:Richard-m-j/open-webui.git
git clone git@github.com:Richard-m-j/ollama.git

snap install docker