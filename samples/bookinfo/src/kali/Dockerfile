FROM kalilinux/kali-linux-docker

# ENV DEBIAN_FRONTEND noninteractive

RUN apt-get -y update && apt-get -y dist-upgrade && apt-get -y autoremove \
               && apt-get clean && \
    apt-get -y install arp-scan bash-completion bettercap binutils dnsenum \
               dnsmap dnsrecon dnswalk dotdotpwn fierce gcc hping3 \
               knot-dnsutils masscan netcat nmap openssh-server tcpdump \
               skipfish sslcaudit sslyze tlssled w3m

RUN mkdir /var/run/sshd
RUN echo 'root:changeme' | chpasswd
RUN sed -i 's/PermitRootLogin prohibit-password/PermitRootLogin yes/;s/#PermitRootLogin yes/PermitRootLogin yes/' /etc/ssh/sshd_config

# SSH login fix. Otherwise user is kicked off after login
RUN sed 's@session\s*required\s*pam_loginuid.so@session optional pam_loginuid.so@g' -i /etc/pam.d/sshd

ENV NOTVISIBLE "in users profile"
RUN echo "export VISIBLE=now" >> /etc/profile

EXPOSE 22
CMD ["/usr/sbin/sshd", "-D"]
