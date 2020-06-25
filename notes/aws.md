# AWS notes

## ssh terms
- ssh- secure shell
- ssh -i ```-i``` is your identity, provide your ssh key here. Key could be store in .ssh folder
- scp is secure copy. This command allows you to upload files to a remote server securly. You will have to provide your identity as well in the form of a ssh key.
  ```
  scp -i ~/.ssh/[key] [filename] ubuntu@DNS.address:
## linux commands

- remove file  
  ```
  rm [file name]
  ```

- remove directory and all files inside of it.  
  ```
  rm -r [filename]
  ```

- run as Administrator
  ```
  sudo
  ```

## Accessing your remote ubuntu server
  ```
  ssh -i ~/.ssh/[key] ubuntu@DNS.address
  ```

## Change permissions on binary file after upload.
  ```
  sudo chmod [777 | 700] [filename]
  ```  

## creating a service file inside a linux server

- Change directory into /system
  ```
  cd /etc/systemd/system
  ```
- Create service file using nano editor  
  ```
  sudo nano [filename].service
  ```

- Paste in this code into the nano editor  

  ```
  [Unit]
  Description=Go Server

  [Service]
  ExecStart=/home/<username>/<path-to-exe>/<exe>
  WorkingDirectory=/home/<username>/<exe-working-dir>
  User=root
  Group=root
  Restart=always

  [Install]
  WantedBy=multi-user.target
  ```

enable the service worker with these commands

  - Add the service to systemd.  
    ```
    sudo systemctl enable <filename>.service
    ```
  - Activate the service.
    ```
    sudo systemctl start <filename>.service
    ```
  - Check if systemd started it.
    ```
    sudo systemctl status <filename>.service
    ```
  - Stop systemd if so desired.
    ```
    sudo systemctl stop <filename>.service
    ```

