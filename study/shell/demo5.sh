
#设置虚拟ip
set_vir_ip(){
  cur_ip=$(hostname -I|awk -F" " '{print $1}')
  echo "cur_ip:"${cur_ip}
  nic=$(ip a|grep ${cur_ip}|awk -F " " '{print $8}')
  echo "nic:"${nic}
  #设置虚拟ip
  ifconfig ${nic}:1 ${VIR_IP}
}

#检查防火墙
open_firewall_port(){
  echo "检查防火墙是否开启"
  systemctl status firewalld
  if [ $? = 0 ];then
      echo 'firewall open '$1
      port_status=$(firewall-cmd --query-port=$1/tcp)
      if [ $port_status = 'no' ]; then
          firewall-cmd --zone=public --add-port=$1/tcp --permanent
          firewall-cmd --reload
      fi
  fi
}


echo $(date +"%Y%m%d%H%M%S":)"manager_report_script被调用...." > aa.log